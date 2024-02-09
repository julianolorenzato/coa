package common

import (
	"errors"
)

func ExtractBitsU32(word uint32, skip, take uint) (error, uint32) {
	if skip > 32 || take > 32 || skip+take > 32 {
		return errors.New("skip or take or the sum of both greater than 32"), 0
	}

	var bitMask uint32 = ((1 << take) - 1) << skip

	return nil, (bitMask & word) >> skip
}
