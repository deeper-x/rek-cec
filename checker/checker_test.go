package checker

import "testing"

func TestCallService(t *testing.T) {
	ok, _ := CallService("http://yahoo.com", 200)
	if !ok {
		t.Errorf("service not called")
	}
}
