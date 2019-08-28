module mygo

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.43.0
	golang.org/x/build => github.com/golang/build v0.0.0-20190801144930-db014ec7fdd2
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190731235908-ec7cb31e5a56
	golang.org/x/image => github.com/golang/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190719004257-d2bd2a29d028
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190620143337-7c3f2128ad9b
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

	golang.org/x/sys => github.com/golang/sys v0.0.0-20190804053845-51ab0e2deafa
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190805165405-2756c524cc1c
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.7.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1

	google.golang.org/genproto => github.com/juelite/google.golang.org-genproto v0.0.0-20181121072036-37b0f0ff607c
	google.golang.org/grpc => github.com/grpc/grpc-go v1.22.1
)

require (
	github.com/gin-gonic/gin v1.4.0
	gopkg.in/go-playground/validator.v8 v8.18.2
)

go 1.13
