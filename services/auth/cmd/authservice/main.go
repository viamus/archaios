package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/viamus/archaios/libs/serviceutils/envy"
	"github.com/viamus/archaios/libs/serviceutils/opentelemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {

	envy.LoadDotEnvIfDev()

	if err := opentelemetry.InitTracer(context.Background(), "user-service"); err != nil {
		log.Fatalf("Fatal Error configuring OTEL: %v", err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := opentelemetry.LoggerWithTrace(ctx)
		logger.Info("recebida requisição no endpoint /hello")

		fmt.Fprintln(w, "Hello from user-service!")
	}

	port := envy.Get("PORT")
	log.Printf("🚀 Server running at http://localhost%s/hello", port)

	wrappedHandler := otelhttp.NewHandler(http.HandlerFunc(handler), "/hello")

	err := http.ListenAndServe(port, wrappedHandler)
	if err != nil {
		log.Fatal(err)
	}
}
