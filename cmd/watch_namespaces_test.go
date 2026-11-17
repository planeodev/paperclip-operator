package main

import (
	"reflect"
	"testing"
)

func TestParseWatchNamespaces(t *testing.T) {
	cases := map[string][]string{
		"":                                   nil,
		" , ":                                nil,
		"paperclip-planeo":                   {"paperclip-planeo"},
		"paperclip-planeo, paperclip-acme ,": {"paperclip-planeo", "paperclip-acme"},
	}
	for in, want := range cases {
		if got := parseWatchNamespaces(in); !reflect.DeepEqual(got, want) {
			t.Errorf("parseWatchNamespaces(%q) = %v, want %v", in, got, want)
		}
	}
}
