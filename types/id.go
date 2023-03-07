package types

import "encoding/binary"

type ID uint64

func (id ID) ToBytes() []byte {
	ret := make([]byte, 8)
	binary.LittleEndian.PutUint64(ret, uint64(id))
	return ret
}
