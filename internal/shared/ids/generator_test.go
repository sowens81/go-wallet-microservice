package ids

import "testing"

func TestGenerateShortID(t *testing.T) {
	id := GenerateShortID()

	if len(id) != 12 {
		t.Errorf("expected length 12, got %d", len(id))
	}
}
