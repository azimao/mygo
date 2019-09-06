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
	cloud.google.com/go v0.44.3 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190906004059-62cf760a6c9e // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pty v1.1.8 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/rogpeppe/go-internal v1.3.1 // indirect
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	golang.org/x/exp v0.0.0-20190829153037-c13cbed26979 // indirect
	golang.org/x/image v0.0.0-20190829233526-b3c06291d021 // indirect
	golang.org/x/mobile v0.0.0-20190830201351-c6da95954960 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20190904154756-749cb33beabd // indirect
	golang.org/x/tools v0.0.0-20190906115428-bc9f4f258ada // indirect
	google.golang.org/api v0.9.0 // indirect
	google.golang.org/appengine v1.6.2 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.23.0 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	honnef.co/go/tools v0.0.1-2019.2.2 // indirect
)

go 1.13
