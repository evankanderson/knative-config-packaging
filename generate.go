package knative_config_packaging

import _ "knative.dev/operator/cmd/fetcher"

//go:generate go run knative.dev/operator/cmd/fetcher --config fetcher-config.yaml --out base
