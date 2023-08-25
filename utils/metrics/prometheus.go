package metrics

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"greatestworks/aop/logging"
	"greatestworks/aop/protos"
)

// escaper is used to format the labels according to [1]. Prometheus labels can
// be any sequence of UTF-8 characters, but the backslash (\), double-quote ("),
// and line feed (\n) characters have to be escaped as \\, \", and \n, respectively.
//
// [1] https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-format-details
var escaper = strings.NewReplacer("\\", `\\`, "\n", `\n`, "\"", `\"`)

// TranslateMetricsToPrometheusTextFormat translates Service Weaver
// metrics (keyed by weavelet id) to a text format that can be
// scraped by Prometheus [1].
//
// [1] https://prometheus.io/
func TranslateMetricsToPrometheusTextFormat(w *bytes.Buffer, ms []*MetricSnapshot, lisAddr, path string) {
	writeHelper(w, lisAddr, path)

	// Sort by name, breaking ties by id.
	sort.SliceStable(ms, func(i, j int) bool {
		if ms[i].Name != ms[j].Name {
			return ms[i].Name < ms[j].Name
		}
		return ms[i].Id < ms[j].Id
	})

	// Display the user metrics first, followed by the Service Weaver
	//  Otherwise, the user's metrics can get buried within
	// the ServiceWeaver
	weaverMetrics := map[string][]*MetricSnapshot{}
	userMetrics := map[string][]*MetricSnapshot{}
	for _, m := range ms {
		if strings.HasPrefix(m.Name, "serviceweaver_") {
			weaverMetrics[m.Name] = append(weaverMetrics[m.Name], m)
		} else {
			userMetrics[m.Name] = append(userMetrics[m.Name], m)
		}
	}
	sortedUserMetrics := maps.Keys(userMetrics)
	sortedWeaverMetrics := maps.Keys(weaverMetrics)
	sort.Strings(sortedUserMetrics)
	sort.Strings(sortedWeaverMetrics)

	// Show the metrics grouped by metric name.
	for _, m := range sortedUserMetrics {
		translateMetrics(w, userMetrics[m])
	}
	if len(weaverMetrics) > 0 {
		fmt.Fprintf(w, "# ┌─────────────────────────────────────┐\n")
		fmt.Fprintf(w, "# │ SERVICEWEAVER AUTOGENERATED METRICS │\n")
		fmt.Fprintf(w, "# └─────────────────────────────────────┘\n\n")
	}
	for _, m := range sortedWeaverMetrics {
		translateMetrics(w, weaverMetrics[m])
	}
}

// writeHelper generates a config.yaml file that can be used by prometheus to
// scrape the exported
func writeHelper(w *bytes.Buffer, lisAddr, path string) {
	const help = `# Metrics in Prometheus text format [1].
#
# To visualize and query the metrics, make sure Prometheus is installed on
# your local machine and then add the following stanza to your Prometheus yaml
# config file:
#
# scrape_configs:
# - job_name: 'prometheus-serviceweaver-scraper'
#   scrape_interval: 5s
#   metrics_path: %s
#   static_configs:
#     - targets: ['%s']
#
# [1]: https://prometheus.io

`
	fmt.Fprintf(w, help, path, lisAddr)
}

// translateMetrics translates a slice of metrics from the Service Weaver format
// to the Prometheus text format. For more details regarding the metric text
// format for Prometheus, see [1].
//
// [1] https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-format-details
func translateMetrics(w *bytes.Buffer, metrics []*MetricSnapshot) string {
	metric := metrics[0]

	// Write the metric HELP. Note that all metrics have the same metric name,
	// so we should display the help and the type only once.
	if len(metric.Help) > 0 {
		w.WriteString("# HELP " + metric.Name + " " + metric.Help + "\n")
	}

	// Write the metric TYPE.
	w.WriteString("# TYPE " + metric.Name)

	isHistogram := false
	switch metric.Type {
	case protos.MetricType_COUNTER:
		w.WriteString(" counter\n")
	case protos.MetricType_GAUGE:
		w.WriteString(" gauge\n")
	case protos.MetricType_HISTOGRAM:
		w.WriteString(" histogram\n")
		isHistogram = true
	}

	for idx, metric := range metrics {
		// Trim labels.
		labels := maps.Clone(metric.Labels)
		delete(labels, "serviceweaver_app")
		delete(labels, "serviceweaver_version")
		if node, ok := labels["serviceweaver_node"]; ok {
			labels["serviceweaver_node"] = logging.Shorten(node)
		}

		// Write the metric definitions.
		//
		// For counter and gauge metrics the definition looks like:
		// metric_name [
		//  "{" label_name "=" `"` label_value `"` { "," label_name "=" `"` label_value `"` } [ "," ] "}"
		// ] value [ timestamp ]
		//
		// For histograms:
		//  Each bucket count of a histogram named x is given as a separate sample
		//  line with the name x_bucket and a label {le="y"} (where y is the upper bound of the bucket).
		//
		//  The bucket with label {le="+Inf"} must exist. Its value must be identical to the value of x_count.
		//
		//  The buckets must appear in increasing numerical order of their label values (for the le).
		//
		//  The sample sum for a summary or histogram named x is given as a separate sample named x_sum.
		//
		//  The sample count for a summary or histogram named x is given as a separate sample named x_count.
		if isHistogram {
			hasInf := false

			var count uint64
			for idx, bound := range metric.Bounds {
				count += metric.Counts[idx]
				writeEntry(w, metric.Name, float64(count), "_bucket", labels, "le", bound)
				if math.IsInf(bound, +1) {
					hasInf = true
				}
			}

			// Account for the +Inf bucket.
			count += metric.Counts[len(metric.Bounds)]
			if !hasInf {
				writeEntry(w, metric.Name, float64(count), "_bucket", labels, "le", math.Inf(+1))
			}
			writeEntry(w, metric.Name, metric.Value, "_sum", labels, "", 0)
			writeEntry(w, metric.Name, float64(count), "_count", labels, "", 0)
		} else { // counter or gauge
			writeEntry(w, metric.Name, metric.Value, "", labels, "", 0)
		}
		if isHistogram && idx != len(metrics)-1 {
			w.WriteByte('\n')
		}
	}
	w.WriteByte('\n')
	return w.String()
}

// writeEntry generates a metric definition entry.
func writeEntry(w *bytes.Buffer, metricName string, value float64, suffix string,
	labels map[string]string, extraLabelName string, extraLabelValue float64) {
	w.WriteString(metricName)
	if len(suffix) > 0 {
		w.WriteString(suffix)
	}
	writeLabels(w, labels, extraLabelName, extraLabelValue)
	w.WriteString(" " + strconv.FormatFloat(value, 'f', -1, 64) + "\n")
}

// writeEntry generates the metric labels.
func writeLabels(w *bytes.Buffer, labels map[string]string,
	extraLabelName string, extraLabelValue float64) {
	if len(labels) == 0 && extraLabelName == "" {
		return
	}

	sortedLabels := maps.Keys(labels)
	sort.Strings(sortedLabels)

	separator := "{"
	for _, l := range sortedLabels {
		w.WriteString(separator + l + `="`)
		escaper.WriteString(w, labels[l]) //nolint:errcheck // bytes.Buffer.Write does not error
		w.WriteByte('"')
		separator = ","
	}
	if len(extraLabelName) > 0 {
		// Set for a histogram metric only.
		w.WriteString(separator + extraLabelName + `="`)
		w.WriteString(strconv.FormatFloat(extraLabelValue, 'f', -1, 64) + "\"")
	}
	w.WriteString("}")
}
