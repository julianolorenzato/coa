package common_test

import (
	"github.com/julianolorenzato/fibit/common"
	"testing"
)

func TestExtractBitsU32(t *testing.T) {
	var word uint32 = 0b1110_1100_1011_1010_0111_0110_0111_0110
	var expected uint32 = 0b0_1110_1100

	_, got := common.ExtractBitsU32(word, 7, 9)
	if got != expected {
		t.Errorf("Failed, got: %b, expected: %b", got, expected)
	}
}
