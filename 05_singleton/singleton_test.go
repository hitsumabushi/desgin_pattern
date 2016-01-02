package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	instance_1 := GetInstance()
	instance_2 := GetInstance()
	if instance_1 != instance_2 {
		t.Errorf("Get singleton should return always same instance, but this method doesn't.")
	}
}
