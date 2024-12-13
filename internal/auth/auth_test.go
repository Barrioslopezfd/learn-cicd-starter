package auth

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got, err := GetAPIKey(http.Header{
		"Authorization": []string{"ApiKey ASDFGHJKL"},
	})
	if err != nil {
		log.Fatal(err)
	}
	want := "ASDFGHJKL"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("I wanted this: %v, but i got this: %v", want, got)
	}
}
