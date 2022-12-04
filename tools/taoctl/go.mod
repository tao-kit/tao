module github.com/sllt/tao/tools/taoctl

go 1.16

require (
	github.com/emicklei/proto v1.11.1
	github.com/fatih/structtag v1.2.0
	github.com/go-sql-driver/mysql v1.7.0
	github.com/iancoleman/strcase v0.2.0
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.0
	github.com/withfig/autocomplete-tools/integrations/cobra v0.0.0-20220705165518-2761d7f4b8bc
	github.com/zeromicro/antlr v0.0.1
	github.com/zeromicro/ddl-parser v1.0.4
	golang.org/x/text v0.3.7
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	github.com/sllt/tao v1.4.1
)

replace github.com/sllt/tao v1.4.1 => ../../
