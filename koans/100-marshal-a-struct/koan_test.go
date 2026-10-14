package main

// Marshal a struct
// Complete the JSON tags and encodeRecord. Use lowercase name and tags keys; omit an empty
// nickname. Nil tags encode as null, empty non-nil tags as []. Unexported fields are
// excluded.

import (
	"encoding/json"
	"testing"
)

type Record struct {
	Name     string   `json:"Name"` // TODO: use the name key.
	Tags     []string `json:"tags"`
	Nickname string   `json:"nickname"` // TODO: omit an empty nickname.
	hidden   string
}

func encodeRecord(record Record) ([]byte, error) {
	return json.Marshal(record.Name) // TODO: encode the struct.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		value Record
		want  string
	}{
		{Record{Name: "Ada", hidden: "private"}, "{\"name\":\"Ada\",\"tags\":null}"},
		{Record{Name: "Go", Tags: []string{}}, "{\"name\":\"Go\",\"tags\":[]}"},
		{Record{Name: "Grace", Tags: []string{"admin"}, Nickname: "G"}, "{\"name\":\"Grace\",\"tags\":[\"admin\"],\"nickname\":\"G\"}"},
	}
	for _, tc := range cases {
		data, err := encodeRecord(tc.value)
		if err != nil || string(data) != tc.want {
			t.Errorf("encodeRecord(%+v)=(%s,%v); want %s", tc.value, data, err, tc.want)
		}
	}
}
