package bit

import (
	"fmt"
	"strconv"
	"strings"
)

type BitsBinary struct {
	value BitsBinaryValue
}

type BitsBinaryValue string

func (s BitsBinaryValue) ToString() string {
	return string(s)
}
func (s BitsBinaryValue) ToInt() (int, error) {
	return strconv.Atoi(s.ToString())
}
func (s BitsBinaryValue) Len() int {
	return len(s)
}

func newBitsBinary(value string) (*BitsBinary, error) {
	return &BitsBinary{
		value: BitsBinaryValue(value)}, nil
}

func (s *BitsBinary) GetValue() interface{} {
	return s.value.ToString()
}

func (s *BitsBinary) On(id interface{}) error {
	return s.update(id, true)
}

func (s *BitsBinary) Off(id interface{}) error {
	return s.update(id, false)
}

func (s *BitsBinary) update(id interface{}, on bool) error {
	pointer, err := parseToPointer(id)
	if err != nil {
		return err
	}
	value := s.value.padding(pointer)
	n, err := value.trimming(-pointer, 1).ToInt()
	if err != nil {
		return err
	}
	if on {
		n |= 1
	} else {
		n &= 0
	}
	ss := strconv.Itoa(n)
	head := value.Len() - pointer
	s.value = BitsBinaryValue(fmt.Sprintf("%s%s%s",
		value.trimming(0, head),
		ss,
		value.trimming(head+len(ss), value.Len()-(head+len(ss))))).suppression(1)
	return nil
}

// 足りない桁数を0埋めします
func (s BitsBinaryValue) padding(pointer int) BitsBinaryValue {
	value := s
	if got := value.Len(); got <= pointer {
		value = BitsBinaryValue(strings.Repeat("0", pointer-got)) + value
	}
	return value
}

// 左から余分な0を省きます
func (s BitsBinaryValue) suppression(offset int) BitsBinaryValue {
	dst := s
	for i, k := 0, BitsBinaryValue(strings.Repeat("0", offset)); i < s.Len(); i += offset {
		if dst.Len() <= offset {
			break
		}
		if s[i:i+offset] != k {
			break
		}
		dst = s[i+offset:]
	}
	return dst
}

func (s BitsBinaryValue) trimming(offset, n int) BitsBinaryValue {
	if n < 1 {
		return ""
	}
	abs := offset
	if abs < 0 {
		abs *= -1
	}
	if s.Len() == abs {
		offset = 0
	}
	//後ろから
	if offset < 0 {
		offset = s.Len() + offset
	}
	if s.Len() < offset+n {
		return ""
	}
	return s[offset : offset+n]
}

func (s *BitsBinary) IsOn(id interface{}) bool {
	pointer, err := parseToPointer(id)
	if err != nil {
		return false
	}
	value := s.value.padding(pointer).trimming(-pointer, 1)
	n, err := value.ToInt()
	if err != nil {
		return false
	}
	return n == 1
}

func (s *BitsBinary) OnesCount() int {
	return s.value.onesCount()
}

func (s *BitsBinaryValue) onesCount() int {
	return strings.Count(s.ToString(), "1")
}
