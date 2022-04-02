package template

// Error defines an error template
const Error = `package {{.pkg}}

import "manlu.org/tao/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
