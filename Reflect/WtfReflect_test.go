package Reflect

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name          string
		input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				name string
			}{"Tim"},
			[]string{"Tim"},
		},
		{
			"Struct with two string field",
			struct {
				name     string
				nickName string
			}{"Tim", "Atom"},
			[]string{"Tim", "Atom"},
		},
		{
			"Struct with a int field",
			struct {
				name string
				id   int
			}{"Tim", 10},
			[]string{"Tim"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string
			Walk(test.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
