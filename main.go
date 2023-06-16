package main

import (
	"fmt"
	"net/http"

	"go.uber.org/fx/fxevent"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("hello world")
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			zap.NewExample,
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(NewEchoHandler),
			AsRoute(NewHelloHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
