package bit

import (
	"errors"
	"regexp"
)

type BitsHex struct {
	BitsBinary
}

var reg = regexp.MustCompile(`[[:xdigit:]]+`)

func NewBitsHex(value string) (*BitsHex, error) {
	b := &BitsHex{BitsBinary{BitsBinaryValue(value)}}
	if value, err := b.toBinary(); err != nil {
		return nil, err
	} else {
		b.value = value
	}
	return b, nil
}

func (s *BitsHex) GetValue() interface{} {
	return s.toHex().ToString()
}

func (s *BitsHex) toHex() BitsBinaryValue {
	var dst BitsBinaryValue
	src := s.value
	if sub := src.Len() % 4; sub != 0 {
		src = src.padding(src.Len() + (4 - sub))
	}
	srcBitLen := src.Len()
	for 0 < srcBitLen {
		switch src[srcBitLen-4 : srcBitLen] {
		case "0000":
			dst = "0" + dst
		case "0001":
			dst = "1" + dst
		case "0010":
			dst = "2" + dst
		case "0011":
			dst = "3" + dst
		case "0100":
			dst = "4" + dst
		case "0101":
			dst = "5" + dst
		case "0110":
			dst = "6" + dst
		case "0111":
			dst = "7" + dst
		case "1000":
			dst = "8" + dst
		case "1001":
			dst = "9" + dst
		case "1010":
			dst = "A" + dst
		case "1011":
			dst = "B" + dst
		case "1100":
			dst = "C" + dst
		case "1101":
			dst = "D" + dst
		case "1110":
			dst = "E" + dst
		case "1111":
			dst = "F" + dst
		}
		srcBitLen -= 4
	}
	return dst.suppression(1)
}

func (s *BitsHex) toBinary() (BitsBinaryValue, error) {
	return HexToBinary(s.value)
}

func (s BitsBinaryValue) isHex() bool {
	return s == "" || reg.MatchString(s.ToString())
}

func HexToBinary(src BitsBinaryValue) (BitsBinaryValue, error) {
	var dst BitsBinaryValue
	if !src.isHex() {
		return "", errors.New("invalid value")
	}
	for i, src := 0, src; i < src.Len(); i++ {
		switch string(src[src.Len()-(1+i)]) {
		case "0":
			dst = "0000" + dst
		case "1":
			dst = "0001" + dst
		case "2":
			dst = "0010" + dst
		case "3":
			dst = "0011" + dst
		case "4":
			dst = "0100" + dst
		case "5":
			dst = "0101" + dst
		case "6":
			dst = "0110" + dst
		case "7":
			dst = "0111" + dst
		case "8":
			dst = "1000" + dst
		case "9":
			dst = "1001" + dst
		case "a", "A":
			dst = "1010" + dst
		case "b", "B":
			dst = "1011" + dst
		case "c", "C":
			dst = "1100" + dst
		case "d", "D":
			dst = "1101" + dst
		case "e", "E":
			dst = "1110" + dst
		case "f", "F":
			dst = "1111" + dst
		}
	}
	return dst.suppression(4), nil
}
