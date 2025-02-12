module github.com/mheers/opa-live-playground

go 1.23.0

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/maddalax/htmgo/framework v1.0.3-0.20241116145200-825c4dd7ecca
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f
	raygun v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace raygun => github.com/paclabsnet/raygun v0.1.4
