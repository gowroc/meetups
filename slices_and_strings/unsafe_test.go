package slices

import (
	"fmt"
	"testing"
)

func TestStringIsImmutable(t *testing.T) {
	s := "dog"
	b := []byte(s)
	b[0] = 'h'
	if s != "dog" {
		t.Fatal("Expected 'dog'")
	}
}

func TestStringIsNotImmutable(t *testing.T) {
	s := "dog"
	b := UnsafeStrToByteArray(s)
	b = append(b, 'a')
	b[0] = 'e'
	fmt.Println(s, string(b))
}
