package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"manlu.org/tao/core/stores/cache"
	"manlu.org/tao/core/stores/sqlc"
	"manlu.org/tao/core/stores/sqlx"
	"manlu.org/tao/core/stringx"
	"manlu.org/tao/tools/taoctl/model/sql/builderx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"manlu.org/tao/core/stores/sqlc"
	"manlu.org/tao/core/stores/sqlx"
	"manlu.org/tao/core/stringx"
	"manlu.org/tao/tools/taoctl/model/sql/builderx"
)
`
)
