package main

// Unmarshal JSON
// Implement decodePerson with json.Unmarshal into a pointer. Return decoding errors for
// malformed JSON or a nonnumeric age; omitted fields keep zero values.

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func decodePerson(data []byte) (Person, error) {
	_ = json.Unmarshal
	return Person{}, nil // TODO: decode into a pointer and return the error.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		data    string
		want    Person
		invalid bool
	}{
		{"{\"name\":\"Ada\",\"age\":30}", Person{Name: "Ada", Age: 30}, false},
		{"{}", Person{}, false},
		{"{\"name\":\"Grace\",\"extra\":true}", Person{Name: "Grace"}, false},
		{"{", Person{}, true},
		{"{\"age\":\"30\"}", Person{}, true},
	}
	for _, tc := range cases {
		got, err := decodePerson([]byte(tc.data))
		if tc.invalid {
			if err == nil {
				t.Errorf("decodePerson(%q) should return an error", tc.data)
			}
			continue
		}
		if err != nil || got != tc.want {
			t.Errorf("decodePerson(%q)=(%+v,%v); want (%+v,nil)", tc.data, got, err, tc.want)
		}
	}
}
