package parser

import (
	"sort"
	"testing"

	"github.com/sllt/tao/tools/taoctl/pkg/ddl/gen"
	"github.com/stretchr/testify/assert"
)

func TestVisitor_VisitTableConstraint(t *testing.T) {
	p := NewParser(WithDebugMode(true))
	accept := func(p *gen.MySqlParser, visitor *visitor) any {
		ctx := p.TableConstraint()
		return visitor.visitTableConstraint(ctx)
	}

	t.Run("uniqueKeyTableConstraint", func(t *testing.T) {
		v, err := p.testMysqlSyntax("test.sql", accept, "UNIQUE INDEX `data_UNIQUE` (`data` ASC)")
		assert.Nil(t, err)
		tc, ok := v.(*TableConstraint)
		assert.True(t, ok)
		assertEqualStringSlice(t, []string{"data"}, tc.ColumnUniqueKey)

		v, err = p.testMysqlSyntax("test.sql", accept, "UNIQUE INDEX `data_UNIQUE` (`data` ASC) INVISIBLE VISIBLE")
		assert.Nil(t, err)
		tc, ok = v.(*TableConstraint)
		assert.True(t, ok)
		assertEqualStringSlice(t, []string{"data"}, tc.ColumnUniqueKey)

		v, err = p.testMysqlSyntax("test.sql", accept, "UNIQUE INDEX `data_UNIQUE` (`data` ASC) INVISIBLE VISIBLE")
		assert.Nil(t, err)
		tc, ok = v.(*TableConstraint)
		assert.True(t, ok)
		assertEqualStringSlice(t, []string{"data"}, tc.ColumnUniqueKey)

		v, err = p.testMysqlSyntax("test.sql", accept, "UNIQUE INDEX `data_UNIQUE` (`column1` ASC, `column2`) INVISIBLE VISIBLE")
		assert.Nil(t, err)
		tc, ok = v.(*TableConstraint)
		assert.True(t, ok)
		assertEqualStringSlice(t, []string{"column1", "column2"}, tc.ColumnUniqueKey)
	})

	t.Run("primaryKeyTableConstraint", func(t *testing.T) {
		v, err := p.testMysqlSyntax("test.sql", accept, "PRIMARY KEY (`description_id`)")
		assert.Nil(t, err)
		tc, ok := v.(*TableConstraint)
		assert.True(t, ok)
		assertEqualStringSlice(t, []string{"description_id"}, tc.ColumnPrimaryKey)
	})

}

func assertEqualStringSlice(t *testing.T, expected, actual []string) {
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}
