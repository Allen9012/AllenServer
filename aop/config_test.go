package aop_test

import (
	"fmt"
	"github.com/Allen9012/AllenServer/aop"
	"github.com/Allen9012/AllenServer/aop/codegen"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBinaryPath(t *testing.T) {
	type testCase struct {
		name   string
		dir    string
		binary string
		expect string
	}

	cwd, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf(`filepath.Abs("."): %v`, err)
	}
	for _, c := range []testCase{
		{"Relative/Relative", ".", "./foo", filepath.Join(cwd, "foo")},
		{"Relative/Abs", ".", "/tmp/foo", "/tmp/foo"},
		{"Abs/Relative", "/bin", "./foo", "/bin/foo"},
		{"Abs/Abs", "/bin", "/tmp/foo", "/tmp/foo"},
	} {
		t.Run(c.name, func(t *testing.T) {
			spec := fmt.Sprintf("[serviceweaver]\nbinary = '%s'\n", c.binary)
			cfgFile := filepath.Join(c.dir, "weaver.toml")
			cfg, err := aop.ParseConfig(cfgFile, spec, codegen.ComponentConfigValidator)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if got, want := cfg.Binary, c.expect; got != want {
				t.Fatalf("binary: got %q, want %q", got, want)
			}
		})
	}
}

func TestParseConfigSection(t *testing.T) {
	type section struct {
		Foo string
		Bar string
		Baz int
	}
	type testCase struct {
		name         string
		initialValue section
		config       string
		expect       section
	}
	for _, c := range []testCase{
		{"missing", section{}, ``, section{}},
		{"empty", section{}, "[section]\n", section{}},
		{
			"full",
			section{},
			`section = { Foo = "foo", Bar = "bar", Baz = 100 }`,
			section{"foo", "bar", 100},
		},
		{
			"partial",
			section{Baz: 200},
			`section = {Foo = "foo", Bar = "bar" }`,
			section{"foo", "bar", 200},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			config, err := aop.ParseConfig("", c.config, codegen.ComponentConfigValidator)
			if err != nil {
				t.Fatal(err)
			}
			got := c.initialValue
			err = aop.ParseConfigSection("section", "", config.Sections, &got)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(c.expect, got); diff != "" {
				t.Fatalf("ParseConfigSection: (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfigErrors(t *testing.T) {
	type testCase struct {
		name          string
		cfg           string
		expectedError string
	}
	for _, c := range []testCase{
		{
			name: "same-process-inter-group-conflict",
			cfg: `
[serviceweaver]
colocate = [["a", "main.go"], ["a", "c"]]
`,
			expectedError: "placed multiple times",
		},
		{
			name: "same-process-intra-group-conflict",
			cfg: `
[serviceweaver]
colocate = [["a", "main.go", "a"]]
`,
			expectedError: "placed multiple times",
		},
		{
			name: "conflicting sections",
			cfg: `
[serviceweaver]
name = "foo"

["greatestworks"]
binary = "/tmp/foo"
`,
			expectedError: "conflicting",
		},
		{
			name: "unknown key",
			cfg: `
[serviceweaver]
badkey = "foo"
`,
			expectedError: "unknown",
		},
		{
			name: "bad rollout",
			cfg: `
[serviceweaver]
rollout = "hello"
`,
			expectedError: "invalid duration",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			_, err := aop.ParseConfig("weaver.toml", c.cfg, codegen.ComponentConfigValidator)
			if err == nil {
				t.Fatalf("unexpected success when expecting %q in\n%s", c.expectedError, c.cfg)
			}
			if !strings.Contains(err.Error(), c.expectedError) {
				t.Fatalf("error %v does not contain %q in\n%s", err, c.expectedError, c.cfg)
			}
		})
	}
}
