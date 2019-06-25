package querybuilder

import (
	"errors"
	"fmt"
	"strings"
)

// Matchall struct
type Matchall struct {
	ExpressionIface
}

// ToString matchall
func (m *Matchall) ToString() string {
	return "(matchall)"
}

// Boost struct
type Boost struct {
	Value int
}

// Term operator struct
type Term struct {
	ExpressionIface
	Value string
	Field string
	Boost *Boost
}

// ToString converts term struct to string
func (t *Term) ToString() string {
	var str strings.Builder
	str.WriteString("(term ")
	if t.Field != "" {
		str.WriteString(fmt.Sprintf("field='%v' ", t.Field))
	}
	if t.Boost != nil {
		str.WriteString(fmt.Sprintf("boost='%v' ", t.Boost.Value))
	}

	str.WriteString(fmt.Sprintf("'%v')", escapeString(t.Value)))
	return str.String()
}

// Phrase search operator
type Phrase struct {
	ExpressionIface
	Value string
	Field string
	Boost *Boost
}

// ToString converts phrase struct to string
func (t *Phrase) ToString() string {
	var str strings.Builder
	str.WriteString("(phrase ")
	if t.Field != "" {
		str.WriteString(fmt.Sprintf("field='%v' ", t.Field))
	}
	if t.Boost != nil {
		str.WriteString(fmt.Sprintf("boost='%v' ", t.Boost.Value))
	}

	str.WriteString(fmt.Sprintf("'%v')", escapeString(t.Value)))
	return str.String()
}

// Near search operator
type Near struct {
	ExpressionIface
	Value    string
	Field    string
	Boost    *Boost
	Distance int
}

// ToString converts phrase struct to string
func (t *Near) ToString() string {
	var str strings.Builder
	str.WriteString("(near ")
	if t.Field != "" {
		str.WriteString(fmt.Sprintf("field='%v' ", t.Field))
	}
	if t.Boost != nil {
		str.WriteString(fmt.Sprintf("boost='%v' ", t.Boost.Value))
	}

	str.WriteString(fmt.Sprintf("distance='%v' '%v')", t.Distance, escapeString(t.Value)))
	return str.String()
}

// Prefix search operator
type Prefix struct {
	ExpressionIface
	Value string
	Field string
	Boost *Boost
}

// ToString converts phrase struct to string
func (t *Prefix) ToString() string {
	var str strings.Builder
	str.WriteString("(prefix ")
	if t.Field != "" {
		str.WriteString(fmt.Sprintf("field='%v' ", t.Field))
	}
	if t.Boost != nil {
		str.WriteString(fmt.Sprintf("boost='%v' ", t.Boost.Value))
	}

	str.WriteString(fmt.Sprintf("'%v')", escapeString(t.Value)))
	return str.String()
}

// Range search operator
type Range struct {
	ExpressionIface
	Field string
	Min   *RangeParameter
	Max   *RangeParameter
}

// RangeParameter struct
type RangeParameter struct {
	Value string
}

// ToString converts range to string
func (r *Range) ToString() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("(range field=%v ", r.Field))

	if r.Min != nil && r.Max == nil {
		str.WriteString(fmt.Sprintf("[%v,}", r.Min.Value))
	} else if r.Min == nil && r.Max != nil {

	} else if r.Min != nil && r.Max != nil {
		str.WriteString(fmt.Sprintf("[%v,%v]", r.Min.Value, r.Max.Value))
	} else {
		panic(errors.New("range must include at least 1 bounding range parameter"))
	}

	str.WriteString(")")

	return str.String()
}
