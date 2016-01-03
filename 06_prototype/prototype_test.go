package prototype

import "testing"

func TestPrototype(t *testing.T) {
	manager := Manager{map[string]product{}}
	dash := LinePen{"---"}
	hat := LinePen{"^^^"}
	manager.Register("dash", &dash)
	manager.Register("hat", &hat)
	clone, err := manager.Create("dash")
	cloned_dash, ok := clone.(*LinePen)
	// Test 1
	if !ok || cloned_dash == &hat || err != nil {
		t.Error("Create method should return another instance.")
	}
	// Test 2
	str := "test string"
	if dash.Use(str) != cloned_dash.Use(str) {
		t.Error("Create method should copy of original instance.")
	}
	// Test 3
	_, err = manager.Create("not registerd")
	if err == nil {
		t.Error("Create method should return error for unregistered names.")
	}

}
