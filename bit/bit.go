package bit

import (
	"errors"
	"fmt"
	"strconv"
)

type BitsType int

const (
	BitsTypeNone BitsType = iota
	BitsTypeBinaryString
	BitsTypeHexString
)

func NewBitsSupport(value interface{}, bt BitsType) (StbitInterface, error) {
	switch t := value.(type) {
	case string:
		switch bt {
		case BitsTypeBinaryString:
			return newBitsBinary(t)
		case BitsTypeHexString:
			return NewBitsHex(t)
		default:
			return nil, errors.New("not supported")
		}
	default:
		return nil, errors.New("not supported")
	}
}

func parseToPointer(id interface{}) (int, error) {
	var value int
	switch t := id.(type) {
	case int:
		value = t
	case int8:
		value = int(t)
	case int16:
		value = int(t)
	case int32:
		value = int(t)
	case int64:
		value = int(t)
	//case uint, uint8, uint16, uint32, uint64:
	//	return uint64(t), nil
	default:
		var err error
		if value, err = strconv.Atoi(fmt.Sprintf("%d", id)); err != nil || value == 0 {
			return 0, errors.New("args need int")
		}
	}
	if value <= 0 {
		return 0, errors.New("id need more than 0")
	}
	return value, nil
}
