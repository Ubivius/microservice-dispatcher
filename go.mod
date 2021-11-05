module github.com/Ubivius/microservice-dispatcher

go 1.16

require (
	agones.dev/agones v1.18.0
	github.com/Ubivius/pkg-telemetry v1.0.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.26.1
	go.opentelemetry.io/otel v1.1.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/controller-runtime v0.10.2
)
