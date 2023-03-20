package parser

import (
	"fmt"

	"github.com/sllt/tao/tools/taoctl/pkg/ddl/console"
	"github.com/sllt/tao/tools/taoctl/pkg/ddl/gen"
)

type visitor struct {
	gen.BaseMySqlParserVisitor
	prefix string
	debug  bool
	logger console.Console
}

func (v *visitor) trace(msg ...any) {
	if v.debug {
		v.logger.Debug("Visit Trace: " + fmt.Sprint(msg...))
	}
}

func (v *visitor) panicWithExpr(expr Token, msg string) {
	if len(v.prefix) == 0 {
		err := fmt.Errorf("%v:%v %s", expr.GetLine(), expr.GetColumn(), msg)
		if v.debug {
			v.logger.Error(err)
		}

		panic(err)
		return
	}

	err := fmt.Errorf("%v line %v:%v %s", v.prefix, expr.GetLine(), expr.GetColumn(), msg)
	if v.debug {
		v.logger.Error(err)
	}

	panic(err)
}
