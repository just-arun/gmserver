package password

import "testing"

func ComparePasswordTest(t *testing.T) {
	if CompareHash([]byte("shit"), []byte("shit")) == nil {
		t.Error("The match should not pass")
	}
}
