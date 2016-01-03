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
	if !ok || cloned_dash == &hat || err != nil {
		t.Errorf("Create method should return another instance.")
	}
}
