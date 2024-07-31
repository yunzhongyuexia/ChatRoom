package tools

import "testing"

func TestGetUID(t *testing.T) {
	id := GetUID()
	t.Log("UID:", id)
}
