package accounthist

import (
	"testing"

	"github.com/eoscanada/eos-go"
	"github.com/stretchr/testify/assert"
)

func TestActionKey(t *testing.T) {
	mamaUint, _ := eos.StringToName("mama")
	key1Bytes := encodeActionKey(mamaUint, 1, uint64(1))
	assert.Equal(t,
		[]byte{
			0x2,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x60, 0xa4, 0x91,
			0x01,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe,
		},
		key1Bytes,
	)

	b1Uint, _ := eos.StringToName("b1")
	key1Bytes = encodeActionKey(b1Uint, 2, uint64(256))
	assert.Equal(t,
		[]byte{
			0x2,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x38,
			0x02,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xff,
		},
		key1Bytes,
	)
}
