package querybuilder

import (
	"fmt"
	"strings"
)

// CompoundQuery structure
type CompoundQuery struct {
	Operator    string
	Expressions []ExpressionIface
	Queries     []*CompoundQuery
}

// ExpressionIface interface
type ExpressionIface interface {
	ToString() string
}

func parseArgs(args []interface{}) ([]ExpressionIface, []*CompoundQuery) {
	expressions := make([]ExpressionIface, 0)
	queries := make([]*CompoundQuery, 0)

	for _, arg := range args {
		switch arg.(type) {
		case *CompoundQuery:
			queries = append(queries, arg.(*CompoundQuery))
			//fmt.Println("its a query")
		case ExpressionIface:
			expressions = append(expressions, arg.(ExpressionIface))
			//fmt.Println("its an expression")
		default:
			//fmt.Println("x>>", x)
		}
	}

	return expressions, queries
}

// And compound query builder
func And(args ...interface{}) *CompoundQuery {
	expressions, queries := parseArgs(args)

	return &CompoundQuery{
		Operator:    "and",
		Expressions: expressions,
		Queries:     queries,
	}
}

// Or compound query builder
func Or(args ...interface{}) *CompoundQuery {
	expressions, queries := parseArgs(args)

	return &CompoundQuery{
		Operator:    "or",
		Expressions: expressions,
		Queries:     queries,
	}
}

// Not compound query builder
func Not(args ...interface{}) *CompoundQuery {
	expressions, queries := parseArgs(args)

	return &CompoundQuery{
		Operator:    "not",
		Expressions: expressions,
		Queries:     queries,
	}
}

// Append chains onto existing compound query
func (q *CompoundQuery) Append(args ...interface{}) *CompoundQuery {
	expressions, queries := parseArgs(args)

	q.Expressions = append(q.Expressions, expressions...)
	q.Queries = append(q.Queries, queries...)

	return q
}

// ToString parses compound query to string
func (q *CompoundQuery) ToString() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("(%v ", q.Operator))

	// write expressions to string
	for _, exp := range q.Expressions {
		str.WriteString(exp.ToString())
	}
	// write queries to string
	for _, qry := range q.Queries {
		str.WriteString(qry.ToString())
	}

	str.WriteString(")")

	return str.String()
}
