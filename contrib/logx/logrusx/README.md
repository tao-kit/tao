### Quick Start

Prerequesites:

- Install `tao`:

```console
go get -u github.com/sllt/tao
```

Download the module:

```console
go get -u github.com/sllt/tao/contrib/logx/logrusx
```

For example:

```go
package main

import (
	"context"
	"time"

	"github.com/sllt/tao/core/logx"
	"github.com/sllt/tao/contrib/logx/logrusx"
)

func main() {
	writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
		logger.SetFormatter(&logrus.JSONFormatter{})
	})
	logx.SetWriter(writer)

	logx.Infow("infow foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Errorw("errorw foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Sloww("sloww foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Error("error")
	logx.Infov(map[string]interface{}{
		"url":     "localhost:8080/hello",
		"attempt": 3,
		"backoff": time.Second,
		"value":   "foo",
	})
	logx.WithDuration(1100*time.Microsecond).Infow("infow withduration",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.WithContext(context.Background()).WithDuration(1100*time.Microsecond).Errorw(
		"errorw withcontext withduration",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.WithDuration(1100*time.Microsecond).WithContext(context.Background()).Errorw(
		"errorw withduration withcontext",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
}
```
