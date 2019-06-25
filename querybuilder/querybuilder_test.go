package querybuilder

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	now := time.Now()
	past := time.Now().Add(-30 * 24 * 60 * time.Minute)
	tests := []struct {
		q *CompoundQuery
		e string
	}{
		{
			q: And(
				&Term{
					Value: "helloworld",
				},
				&Term{
					Value: "bar",
					Field: "foo",
				},
				&Term{
					Field: "qux",
					Value: "baz",
					Boost: &Boost{Value: 1},
				},
			),
			e: "(and (term 'helloworld')(term field='foo' 'bar')(term field='qux' boost='1' 'baz'))",
		},
		{
			q: And(
				&Phrase{
					Value: "hello world",
				},
				&Phrase{
					Value: "hello world",
					Field: "foo",
				},
				&Phrase{
					Value: "hello world",
					Field: "bar",
					Boost: &Boost{Value: 2},
				},
			),
			e: "(and (phrase 'hello world')(phrase field='foo' 'hello world')(phrase field='bar' boost='2' 'hello world'))",
		},
		{
			q: And(
				&Near{
					Value: "hello world",
				},
				&Near{
					Value: "hello world",
					Field: "foo",
				},
				&Near{
					Value:    "hello world",
					Field:    "bar",
					Boost:    &Boost{Value: 2},
					Distance: 3,
				},
			),
			e: "(and (near distance='0' 'hello world')(near field='foo' distance='0' 'hello world')(near field='bar' boost='2' distance='3' 'hello world'))",
		},
		{
			q: And(
				&Prefix{
					Value: "hello",
				},
				&Prefix{
					Value: "hello",
					Field: "foo",
				},
				&Prefix{
					Value: "hello",
					Field: "bar",
					Boost: &Boost{Value: 2},
				},
			),
			e: "(and (prefix 'hello')(prefix field='foo' 'hello')(prefix field='bar' boost='2' 'hello'))",
		},
		{
			q: And(
				Or(
					&Term{
						Field: "foo",
						Value: "yes",
					},
					&Term{
						Field: "foo",
						Value: "maybe",
					},
				),
				Not(
					&Term{
						Field: "foo",
						Value: "no",
					},
				),
			),
			e: "(and (or (term field='foo' 'yes')(term field='foo' 'maybe'))(not (term field='foo' 'no')))",
		},
		{
			q: And(
				&Range{
					Field: "year",
					Min:   &RangeParameter{Number(2000)},
				},
				&Range{
					Field: "created_at",
					Min:   &RangeParameter{Time(past)},
					Max:   &RangeParameter{Time(now)},
				},
				&Range{
					Field: "price",
					Min:   &RangeParameter{Number(1000000)},
				},
			),
			e: fmt.Sprintf("(and (range field=year [2000,})(range field=created_at [%v,%v])(range field=price [1000000,}))", past.Format("2006-01-02T15:04:05Z"), now.Format("2006-01-02T15:04:05Z")),
		},
		{
			q: And(
				&Term{
					Value: `30' a\b`,
				},
			),
			e: `(and (term '30\' a\\b'))`,
		},
	}

	for _, tt := range tests {
		//fmt.Println(tt.q.ToString())
		assert.Equal(t, tt.e, tt.q.ToString())
	}

}
