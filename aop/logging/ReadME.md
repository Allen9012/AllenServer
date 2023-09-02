该logging来自于
https://github.com/ServiceWeaver/weaver/tree/main/runtime/logging

Logging
Service Weaver provides a logging API, weaver.Logger. By using Service Weaver's logging API, you can cat, tail, search, and filter logs from every one of your Service Weaver applications (past or present). Service Weaver also integrates the logs into the environment where your application is deployed. If you deploy a Service Weaver application to Google Cloud, for example, logs are automatically exported to Google Cloud Logging.

Use the Logger method of a component implementation to get a logger scoped to the component. For example:
```go
type Adder interface {
    Add(context.Context, int, int) (int, error)
}

type adder struct {
    weaver.Implements[Adder]
}

func (a *adder) Add(ctx context.Context, x, y int) (int, error) {
    // adder embeds weaver.Implements[Adder] which provides the Logger method.
    logger := a.Logger(ctx)
    logger.Debug("A debug log.")
    logger.Info("An info log.")
    logger.Error("An error log.", fmt.Errorf("an error"))
    return x + y, nil
}

```

Logs look like this:

```text
D1103 08:55:15.650138 main.Adder 73ddcd04 adder.go:12] A debug log.
I1103 08:55:15.650149 main.Adder 73ddcd04 adder.go:13] An info log.
E1103 08:55:15.650158 main.Adder 73ddcd04 adder.go:14] An error log. err="an error"
```
The first character of a log line indicates whether the log is a [D]ebug, [I]nfo, or [E]rror log entry. Then comes the date in MMDD format, followed by the time. Then comes the component name followed by a logical node id. If two components are co-located in the same OS process, they are given the same node id. Then comes the file and line where the log was produced, followed finally by the contents of the log.

Service Weaver also allows you to attach key-value attributes to log entries. These attributes can be useful when searching and filtering logs.

```go
logger.Info("A log with attributes.", "foo", "bar")  // adds foo="bar"
```
If you find yourself adding the same set of key-value attributes repeatedly, you can pre-create a logger that will add those attributes to all log entries:
```go
fooLogger = logger.With("foo", "bar")
fooLogger.Info("A log with attributes.")  // adds foo="bar"
```
Note: You can also add normal print statements to your code. These prints will be captured and logged by Service Weaver, but they won't be associated with a particular component, they won't have file:line information, and they won't have any attributes, so we recommend you use a weaver.Logger whenever possible.
```text
S1027 14:40:55.210541 stdout d772dcad] This was printed by fmt.Println
```