package main

import (
	"fmt"
	"testing"
)

func TestMessageWithEnvelope_Operations(t *testing.T) {
	for _, test := range []struct {
		operation    string
		expectCreate bool
		expectUpdate bool
		expectDelete bool
	}{
		{"None", false, false, false},
		{"CREATE", true, false, false},
		{"create", false, false, false},
		{"Create", false, false, false},
		{"UPDATE", false, true, false},
		{"uPdATe", false, false, false},
		{"upDATE", false, false, false},
		{"DELETE", false, false, true},
		{"delete", false, false, false},
		{"DEELETE", false, false, false},
	} {
		t.Run("When Operation is "+test.operation, func(t *testing.T) {
			mwe := MessageWithEnvelope{
				Operation: test.operation,
			}

			t.Run(fmt.Sprintf("Create() should return %v", test.expectCreate), func(t *testing.T) {
				rcvd := mwe.Create()
				if test.expectCreate != rcvd {
					t.Errorf("Expected %v, received %v", test.expectCreate, rcvd)
				}
			})

			t.Run(fmt.Sprintf("Update() should return %v", test.expectUpdate), func(t *testing.T) {
				rcvd := mwe.Update()
				if test.expectUpdate != rcvd {
					t.Errorf("Expected %v, received %v", test.expectUpdate, rcvd)
				}
			})

			t.Run(fmt.Sprintf("Delete() should return %v", test.expectDelete), func(t *testing.T) {
				rcvd := mwe.Delete()
				if test.expectDelete != rcvd {
					t.Errorf("Expected %v, received %v", test.expectDelete, rcvd)
				}
			})
		})
	}
}

func TestParseMessage(t *testing.T) {
	msg := `{
    "operation": "CREATE",
    "provenance": "wordpress",
    "version": "1.0.0",
    "message": {
        "slug": "a-post",
        "title": "This is my first post<3",
        "author": "a.n. other",
        "date": "2019-04-20T17:13:06Z",
        "body": "Hello, world!"
    }
}
`

	_, err := ParseMessage([]byte(msg))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
}
