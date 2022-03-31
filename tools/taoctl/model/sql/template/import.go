package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"manlu.org/tao/core/stores/builder"
	"manlu.org/tao/core/stores/cache"
	"manlu.org/tao/core/stores/sqlc"
	"manlu.org/tao/core/stores/sqlx"
	"manlu.org/tao/core/stringx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"manlu.org/tao/core/stores/builder"
	"manlu.org/tao/core/stores/sqlc"
	"manlu.org/tao/core/stores/sqlx"
	"manlu.org/tao/core/stringx"
)
`
)
