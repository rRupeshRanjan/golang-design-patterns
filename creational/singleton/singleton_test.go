package singleton

import (
	"testing"
)

func TestGetBookInstance(t *testing.T) {
	book1 := GetBookInstance()
	book1.SetName("test book")
	book1.SetId(1)

	book2 := GetBookInstance()
	book2.SetName("another test book")

	if book2.GetId() != 1 {
		t.Errorf("Expected id: %d, instead got %d", 1, 2)
	}

	if book1.GetName() != "another test book" {
		t.Errorf("Expected name: %s, instead got %s", "another test book", "test book")
	}
}
