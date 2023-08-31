package testdata

import _ "embed"

var (
	//go:embed unformat.api
	unformatApi string
	//go:embed kotlin.api
	kotlinApi string
	//go:embed user.sql
	userSql string

	list = Files{
		{
			IsDir: true,
			Path:  "version",
			Cmd:   "taoctl --version",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local",
			Cmd:   "taoctl api --o sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local/assign",
			Cmd:   "taoctl api --o=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote",
			Cmd:   "taoctl api --o sample.api --remote https://github.com/sllt/tao-template --branch main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/shorthand",
			Cmd:   "taoctl api -o sample.api -remote https://github.com/sllt/tao-template -branch main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/assign",
			Cmd:   "taoctl api --o=sample.api --remote https://github.com/sllt/tao-template --branch=main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api -remote https://github.com/sllt/tao-template -branch=main",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true",
			Cmd:   "taoctl api --o sample.api && taoctl api dart --api sample.api --dir . --hostname 127.0.0.1 --legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api dart -api sample.api -dir . -hostname 127.0.0.1 -legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api dart --api=sample.api --dir=. --hostname=127.0.0.1 --legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api && taoctl api dart -api=sample.api -dir=. -hostname=127.0.0.1 -legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false",
			Cmd:   "taoctl api --o sample.api && taoctl api dart --api sample.api --dir . --hostname 127.0.0.1 --legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api dart -api sample.api -dir . -hostname 127.0.0.1 -legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api dart --api=sample.api --dir=. --hostname=127.0.0.1 --legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api && taoctl api dart -api=sample.api -dir=. -hostname=127.0.0.1 -legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/doc",
			Cmd:   "taoctl api --o sample.api && taoctl api doc --dir . --o .",
		},
		{
			IsDir: true,
			Path:  "api/doc/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api doc -dir . -o .",
		},
		{
			IsDir: true,
			Path:  "api/doc/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api doc --dir=. --o=.",
		},
		{
			IsDir: true,
			Path:  "api/doc/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api && taoctl api doc -dir=. -o=.",
		},
		{
			Path:    "api/format/unformat.api",
			Content: unformatApi,
			Cmd:     "taoctl api format --dir . --iu",
		},
		{
			Path:    "api/format/shorthand/unformat.api",
			Content: unformatApi,
			Cmd:     "taoctl api format -dir . -iu",
		},
		{
			Path:    "api/format/assign/unformat.api",
			Content: unformatApi,
			Cmd:     "taoctl api format --dir=. --iu",
		},
		{
			Path:    "api/format/assign/shorthand/unformat.api",
			Content: unformatApi,
			Cmd:     "taoctl api format -dir=. -iu",
		},
		{
			IsDir: true,
			Path:  "api/go/style/default",
			Cmd:   "taoctl api --o sample.api && taoctl api go --api sample.api --dir .",
		},
		{
			IsDir: true,
			Path:  "api/go/style/default/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api go -api sample.api -dir .",
		},
		{
			IsDir: true,
			Path:  "api/go/style/assign/default",
			Cmd:   "taoctl api --o=sample.api && taoctl api go --api=sample.api --dir=.",
		},
		{
			IsDir: true,
			Path:  "api/go/style/assign/default/shorthand",
			Cmd:   "taoctl api -o=sample.api && taoctl api go -api=sample.api -dir=.",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero",
			Cmd:   "taoctl api --o sample.api && taoctl api go --api sample.api --dir . --style goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api go -api sample.api -dir . -style goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api go --api=sample.api --dir=. --style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/assign/shorthand",
			Cmd:   "taoctl api -o=sample.api && taoctl api go -api=sample.api -dir=. -style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/java",
			Cmd:   "taoctl api --o sample.api && taoctl api java --api sample.api --dir .",
		},
		{
			IsDir: true,
			Path:  "api/java/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api java -api sample.api -dir .",
		},
		{
			IsDir: true,
			Path:  "api/java/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api java --api=sample.api --dir=.",
		},
		{
			IsDir: true,
			Path:  "api/java/shorthand/assign",
			Cmd:   "taoctl api -o=sample.api && taoctl api java -api=sample.api -dir=.",
		},
		{
			IsDir: true,
			Path:  "api/new/style/default",
			Cmd:   "taoctl api new greet",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero",
			Cmd:   "taoctl api new greet --style goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/assign",
			Cmd:   "taoctl api new greet --style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/shorthand",
			Cmd:   "taoctl api new greet -style goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/shorthand/assign",
			Cmd:   "taoctl api new greet -style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/ts",
			Cmd:   "taoctl api --o sample.api && taoctl api ts --api sample.api --dir . --unwrap --webapi .",
		},
		{
			IsDir: true,
			Path:  "api/ts/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api ts -api sample.api -dir . -unwrap -webapi .",
		},
		{
			IsDir: true,
			Path:  "api/ts/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api ts --api=sample.api --dir=. --unwrap --webapi=.",
		},
		{
			IsDir: true,
			Path:  "api/ts/shorthand/assign",
			Cmd:   "taoctl api -o=sample.api && taoctl api ts -api=sample.api -dir=. -unwrap -webapi=.",
		},
		{
			IsDir: true,
			Path:  "api/validate",
			Cmd:   "taoctl api --o sample.api && taoctl api validate --api sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/shorthand",
			Cmd:   "taoctl api -o sample.api && taoctl api validate -api sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/assign",
			Cmd:   "taoctl api --o=sample.api && taoctl api validate --api=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/shorthand/assign",
			Cmd:   "taoctl api -o=sample.api && taoctl api validate -api=sample.api",
		},
		{
			IsDir: true,
			Path:  "env/show",
			Cmd:   "taoctl env > env.txt",
		},
		{
			IsDir: true,
			Path:  "env/check",
			Cmd:   "taoctl env check -f -v",
		},
		{
			IsDir: true,
			Path:  "env/install",
			Cmd:   "taoctl env install -v",
		},
		{
			IsDir: true,
			Path:  "kube",
			Cmd:   "taoctl kube deploy --image alpine --name foo --namespace foo --o foo.yaml --port 8888",
		},
		{
			IsDir: true,
			Path:  "kube/shorthand",
			Cmd:   "taoctl kube deploy -image alpine -name foo -namespace foo -o foo.yaml -port 8888",
		},
		{
			IsDir: true,
			Path:  "kube/assign",
			Cmd:   "taoctl kube deploy --image=alpine --name=foo --namespace=foo --o=foo.yaml --port=8888",
		},
		{
			IsDir: true,
			Path:  "kube/shorthand/assign",
			Cmd:   "taoctl kube deploy -image=alpine -name=foo -namespace=foo -o=foo.yaml -port=8888",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache",
			Cmd:   "taoctl model mongo --dir . --type user --style goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/shorthand",
			Cmd:   "taoctl model mongo -dir . -type user -style goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/assign",
			Cmd:   "taoctl model mongo --dir=. --type=user --style=goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/shorthand/assign",
			Cmd:   "taoctl model mongo -dir=. -type=user -style=goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache",
			Cmd:   "taoctl model mongo --dir . --type user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/shorthand",
			Cmd:   "taoctl model mongo -dir . -type user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/assign",
			Cmd:   "taoctl model mongo --dir=. --type=user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/shorthand/assign",
			Cmd:   "taoctl model mongo -dir=. -type=user",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/user.sql",
			Cmd:     "taoctl model mysql ddl --database user --dir cache --src user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/user.sql",
			Cmd:     "taoctl model mysql ddl -database user -dir cache -src user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/assign/user.sql",
			Cmd:     "taoctl model mysql ddl --database=user --dir=cache --src=user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/assign/user.sql",
			Cmd:     "taoctl model mysql ddl -database=user -dir=cache -src=user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/user.sql",
			Cmd:     "taoctl model mysql ddl --database user --dir nocache --src user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/user.sql",
			Cmd:     "taoctl model mysql ddl -database user -dir nocache -src user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/assign/user.sql",
			Cmd:     "taoctl model mysql ddl --database=user --dir=nocache --src=user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/assign/user.sql",
			Cmd:     "taoctl model mysql ddl -database=user -dir=nocache -src=user.sql",
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource",
			Cmd:   `taoctl model mysql datasource --url $DSN --dir cache --table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand",
			Cmd:   `taoctl model mysql datasource -url $DSN -dir cache -table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2",
			Cmd:   `taoctl model mysql datasource -url $DSN -dir cache -t "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/assign",
			Cmd:   `taoctl model mysql datasource --url=$DSN --dir=cache --table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand/assign",
			Cmd:   `taoctl model mysql datasource -url=$DSN -dir=cache -table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2/assign",
			Cmd:   `taoctl model mysql datasource -url=$DSN -dir=cache -t="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource",
			Cmd:   `taoctl model mysql datasource --url $DSN --dir nocache --table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand",
			Cmd:   `taoctl model mysql datasource -url $DSN -dir nocache -table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2",
			Cmd:   `taoctl model mysql datasource -url $DSN -dir nocache -t "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/assign",
			Cmd:   `taoctl model mysql datasource --url=$DSN --dir=nocache --table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand/assign",
			Cmd:   `taoctl model mysql datasource -url=$DSN -dir=nocache -table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2/assign",
			Cmd:   `taoctl model mysql datasource -url=$DSN -dir=nocache -t="*" -c`,
		},
		{
			IsDir: true,
			Path:  "rpc/new",
			Cmd:   "taoctl rpc new greet",
		},
		{
			IsDir: true,
			Path:  "rpc/template",
			Cmd:   "taoctl rpc template --o greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/shorthand",
			Cmd:   "taoctl rpc template -o greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/assign",
			Cmd:   "taoctl rpc template --o=greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/shorthand/assign",
			Cmd:   "taoctl rpc template -o=greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/protoc",
			Cmd:   "taoctl rpc template --o greet.proto && taoctl rpc protoc greet.proto --go_out . --go-grpc_out . --zrpc_out .",
		},
		{
			IsDir: true,
			Path:  "rpc/protoc/assign",
			Cmd:   "taoctl rpc template --o=greet.proto && taoctl rpc protoc greet.proto --go_out=. --go-grpc_out=. --zrpc_out=.",
		},
	}
)
