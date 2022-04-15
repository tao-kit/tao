package pg

import (
	// imports the driver, don't remove this comment, golint requires.
	_ "github.com/lib/pq"
	"manlu.org/tao/core/stores/pg"
)

// New creates a new postgresql store.
// Deprecated: use pg.New instead.
var New = pg.New
