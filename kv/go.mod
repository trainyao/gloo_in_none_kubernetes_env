module github.com/trainyao/gloo_in_none_kubernetes_env/kv

go 1.13

require (
	github.com/trainyao/gloo_in_none_kubernetes_env/kv/kv v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0
	google.golang.org/grpc v1.26.0
)

replace github.com/trainyao/gloo_in_none_kubernetes_env/kv/kv => ./kv
