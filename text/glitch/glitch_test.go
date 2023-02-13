package glitch

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	sb := &strings.Builder{}
	for i := 0; i < 2000; i++ {
		sb.WriteString("a")
		zalgo := EncodeString(sb.String(), WithMaxHeight(1), WithRandomization(0))
		if sb.Len()*7 != len(zalgo) {
			t.Fatalf("expected string to be %d length, but got %d", sb.Len()*7, len(zalgo))
		}
	}
}
