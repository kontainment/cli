module github.com/kontainment/cli

go 1.19

require github.com/spf13/cobra v1.6.1

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.22+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/kontainment/engine v0.0.0-20230115225714-7fd5017602bf // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/tools v0.1.12 // indirect
)

// for rapid local development changes and testing
replace github.com/kontainment/engine => /Users/bpalmer/github/kontainment/engine