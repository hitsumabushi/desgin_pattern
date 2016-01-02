package outer_singleton

import (
	"github.com/hitsumabushi/desgin_pattern_golang/05_singleton/singleton"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance_1 := singleton.GetInstance()
	instance_2 := singleton.GetInstance()
	if instance_1 != instance_2 {
		t.Errorf("Get singleton should return always same instance, but this method doesn't.")
	}
}
