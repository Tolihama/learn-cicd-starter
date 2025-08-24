package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	result, _ := GetAPIKey(make(map[string][]string))
	wantedResult := ""
	if !reflect.DeepEqual(wantedResult, result) {
		t.Fatalf("expected: %v, got: %v", wantedResult, result)
	}
}
