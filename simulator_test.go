package main

import (
	"bytes"
	"simulator/entity"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	list := []string{"foo", "bar"}

	isInSlice := stringInSlice("foo", list)

	if !isInSlice {
		t.Error("Expected foo to be in [\"foo\",\"bar\"]")
	}
}

func TestTableStringInSlice(t *testing.T) {
	tests := []struct {
		word     string
		list     []string
		expected bool
	}{
		{"foo", []string{"foo", "bar"}, true},
		{"bar", []string{"foo", "bar"}, true},
		{"baz", []string{"foo", "bar"}, false},
	}

	for _, test := range tests {
		if output := stringInSlice(test.word, test.list); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received {}", test.word, test.expected, output)
		}
	}
}

func TestDestinationToJson(t *testing.T) {
	order := entity.Order{
		Uuid:        "some-long-id",
		Destination: "1",
	}

	json := destinationToJson(order, "-40.123212312", "-40.123212312")

	expected := []byte("{\"order\":\"some-long-id\",\"lat\":\"-40.123212312\",\"lng\":\"-40.123212312\"}")

	if bytes.Compare(json, expected) != 0 {
		t.Error("Expected {}, returnde {}", expected, json)
	}
}
