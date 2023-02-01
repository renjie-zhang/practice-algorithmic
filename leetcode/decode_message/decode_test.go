package decode_message

import (
	"fmt"
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	var key = "the quick brown fox jumps over the lazy dog"
	var message = "vkbs bs t suepuv"
	fmt.Print(DecodeMessage(key, message))
}
