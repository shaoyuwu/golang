module goods-client

go 1.16

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20220930113650-c6815a8c17ad // indirect
	github.com/cloudflare/circl v1.2.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-micro/plugins/v4/registry/consul v1.1.0
	github.com/google/uuid v1.3.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/urfave/cli/v2 v2.17.1 // indirect
	github.com/xanzy/ssh-agent v0.3.2 // indirect
	go-micro.dev/v4 v4.9.0
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
	golang.org/x/sync v0.0.0-20220929204114-8fcdb60fdcc0 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	google.golang.org/protobuf v1.28.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
