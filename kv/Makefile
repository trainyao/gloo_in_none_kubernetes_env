

.PHONY : main vendor

main : vendor
	docker build . -f Dockerfile -t trainyao/gloo_in_none_kubernetes_env_kv

vendor :
	go mod vendor

