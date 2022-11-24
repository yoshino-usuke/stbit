package bit

import (
	"testing"
)

func Test_bit_padding(t *testing.T) {
	t.Run("padding", func(t *testing.T) {
		{
			b, err := newBitsBinary("0")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "00000", b.value.padding(5).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("111")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "00111", b.value.padding(5).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("1")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "1", b.value.padding(1).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}

		{
			b, err := newBitsBinary("000")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "000", b.value.padding(1).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "00", b.value.padding(2).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "000", b.value.padding(3).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("1")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "001", b.value.padding(3).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "0", b.value.padding(1).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
		{
			b, err := newBitsBinary("")
			if err != nil {
				t.Fatal(err)
			}
			if want, got := "00", b.value.padding(2).ToString(); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
		}
	})
}

func Test_bit_trimming(t *testing.T) {
	for _, test := range []struct {
		name, value, want string
		offset, n         int
		back              bool
	}{
		{
			name:   "case 1",
			value:  "0",
			want:   "0",
			offset: 0,
			n:      1,
		},
		{
			name:   "case 2",
			value:  "010",
			want:   "1",
			offset: 1,
			n:      1,
		},
		{
			name:   "case 3",
			value:  "010",
			want:   "0",
			offset: -1,
			n:      1,
		},
		{
			name:   "case 4",
			value:  "010000",
			want:   "1",
			offset: -5,
			n:      1,
		},
		{
			name:   "case 5",
			value:  "010000",
			want:   "1",
			offset: 1,
			n:      1,
		},
		{
			name:   "case 6",
			value:  "01",
			want:   "0",
			offset: -2,
			n:      1,
		},
		{
			name:   "case 7",
			value:  "1",
			want:   "1",
			offset: 0,
			n:      1,
		},
		{
			name:   "case 8",
			value:  "00",
			want:   "0",
			offset: -1,
			n:      1,
		},
		{
			name:   "case 9",
			value:  "0",
			want:   "0",
			offset: -1,
			n:      1,
		},
		{
			name:   "case 10",
			value:  "010000",
			want:   "01",
			offset: 0,
			n:      2,
		},
		{
			name:   "case 11",
			value:  "000001",
			want:   "01",
			offset: -2,
			n:      2,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b, err := newBitsBinary(test.value)
			if err != nil {
				t.Fatal(err)
			}
			if got := b.value.trimming(test.offset, test.n).ToString(); got != test.want {
				t.Errorf("want %s,got %s", test.want, got)
			}
		})
	}
}

func Test_bit_on_off(t *testing.T) {
	t.Run("正常系 on", func(t *testing.T) {
		value := "0"
		b, err := newBitsBinary(value)
		if err != nil {
			t.Fatal(err)
		}
		t.Run("1", func(t *testing.T) {
			if err := b.On(1); err != nil {
				t.Fatal(err)
			}
			if want, got := "1", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(1) {
				t.Errorf("want true")
			}
		})
		t.Run("2", func(t *testing.T) {
			if err := b.On(2); err != nil {
				t.Fatal(err)
			}
			if want, got := "11", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(2) {
				t.Errorf("want true")
			}
		})
		t.Run("3", func(t *testing.T) {
			if err := b.On(3); err != nil {
				t.Fatal(err)
			}
			if want, got := "111", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(3) {
				t.Errorf("want true")
			}
		})
		t.Run("5", func(t *testing.T) {
			if err := b.On(5); err != nil {
				t.Fatal(err)
			}
			if want, got := "10111", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(5) {
				t.Errorf("want true")
			}
		})
		t.Run("6", func(t *testing.T) {
			if err := b.On(6); err != nil {
				t.Fatal(err)
			}
			if want, got := "110111", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(6) {
				t.Errorf("want true")
			}
		})
		t.Run("8", func(t *testing.T) {
			if err := b.On(8); err != nil {
				t.Fatal(err)
			}
			if want, got := "10110111", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(8) {
				t.Errorf("want true")
			}
		})
		t.Run("11", func(t *testing.T) {
			if err := b.On(11); err != nil {
				t.Fatal(err)
			}
			if want, got := "10010110111", b.GetValue().(string); want != got {
				t.Errorf("want %s,got %s", want, got)
			}
			if !b.IsOn(11) {
				t.Errorf("want true")
			}
		})
	})
	t.Run("正常系2 on", func(t *testing.T) {
		for _, test := range []struct {
			name, value, want string
			arg               interface{}
		}{
			{
				name:  "すでにonの時は変化無し",
				value: "00000001",
				arg:   1,
				want:  "1",
			},
			{
				name:  "すでにonの時は変化無し1",
				value: "01",
				arg:   1,
				want:  "1",
			},
			{
				name:  "すでにonの時は変化無し2",
				value: "1000000001",
				arg:   10,
				want:  "1000000001",
			},
			{
				name:  "すでにonの時は変化無し3",
				value: "11",
				arg:   1,
				want:  "11",
			},
			{
				name:  "すでにonの時は変化無し4",
				value: "10",
				arg:   2,
				want:  "10",
			},
			{
				name:  "1をあげる",
				value: "0",
				arg:   1,
				want:  "1",
			},
			{
				name:  "1をあげる2",
				value: "10",
				arg:   1,
				want:  "11",
			},
			{
				name:  "1をあげる3",
				value: "00",
				arg:   1,
				want:  "1",
			},
			{
				name:  "2をあげる",
				value: "00000000",
				arg:   2,
				want:  "10",
			},
			{
				name:  "3をあげる",
				value: "00000000",
				arg:   3,
				want:  "100",
			},
			{
				name:  "3をあげる2",
				value: "00000001",
				arg:   3,
				want:  "101",
			},
			{
				name:  "50をあげる",
				value: "00000111",
				arg:   50,
				want:  "10000000000000000000000000000000000000000000000111",
			},
			{
				name:  "100をあげる",
				value: "10000000000000000000000000000000000000000000000111",
				arg:   100,
				want:  "1000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000111",
			},
			{
				name:  "100をあげる2",
				value: "",
				arg:   100,
				want:  "1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			},
			{
				name:  "2000をあげる2",
				value: "",
				arg:   2000,
				want:  "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				b, err := newBitsBinary(test.value)
				if err != nil {
					t.Fatal(err)
				}
				if err := b.On(test.arg); err != nil {
					t.Fatal(err)
				}
				if got := b.GetValue().(string); got != test.want {
					t.Errorf("want %s,got %v", test.want, got)
				}
				if !b.IsOn(test.arg) {
					t.Errorf("want true")
				}
			})
		}

	})

	t.Run("正常系 off", func(t *testing.T) {
		for _, test := range []struct {
			name, value, want string
			arg               int
		}{
			{
				name:  "すでにoffの時は変化無し",
				value: "0",
				arg:   1,
				want:  "0",
			},
			{
				name:  "1を落とす",
				value: "11111111",
				arg:   1,
				want:  "11111110",
			},
			{
				name:  "1を落とす 2",
				value: "",
				arg:   1,
				want:  "0",
			},
			{
				name:  "2を落とす",
				value: "01111111",
				arg:   2,
				want:  "1111101",
			},
			{
				name:  "2を落とす 2",
				value: "1",
				arg:   2,
				want:  "1",
			},
			{
				name:  "2を落とす 3",
				value: "",
				arg:   2,
				want:  "0",
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				b, err := newBitsBinary(test.value)
				if err != nil {
					t.Fatal(err)
				}
				if err := b.Off(test.arg); err != nil {
					t.Fatal(err)
				}
				if got := b.GetValue().(string); got != test.want {
					t.Errorf("want %s,got %s", test.want, got)
				}
			})
		}
	})
}

func TestBit_isOn(t *testing.T) {
	for _, test := range []struct {
		name, value string
		want        bool
		arg         interface{}
	}{
		{
			name:  "case 1",
			value: "",
			arg:   1,
			want:  false,
		},
		{
			name:  "case 2",
			value: "0",
			arg:   1,
			want:  false,
		},
		{
			name:  "case 3",
			value: "01",
			arg:   1,
			want:  true,
		},
		{
			name:  "case 4",
			value: "10",
			arg:   1,
			want:  false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b, err := newBitsBinary(test.value)
			if err != nil {
				t.Fatal(err)
			}
			if got := b.IsOn(test.arg); got != test.want {
				t.Errorf("want %v got %v", test.want, got)
			}
		})
	}
}

type HogeBitType int

func Test_ParseToPointer(t *testing.T) {
	b, err := newBitsBinary("")
	if err != nil {
		t.Fatal(err)
	}
	var n HogeBitType = 10
	if err := b.On(n); err != nil {
		t.Fatal(err)
	}

	if err := b.On("hoge"); err == nil {
		t.Errorf("want err")
	}
}

func Test_BitHex16(t *testing.T) {
	for _, test := range []struct {
		name, value, want string
		wantErr           bool
	}{
		{
			name:    "復元テスト 1",
			value:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "復元テスト 2",
			value:   "10",
			want:    "10",
			wantErr: false,
		},
		{
			name:    "復元テスト 3",
			value:   "1011",
			want:    "1011",
			wantErr: false,
		},
		{
			name:    "復元テスト 4",
			value:   "hhh",
			wantErr: true,
		},
		{
			name:    "復元テスト 5",
			value:   "abc",
			want:    "ABC",
			wantErr: false,
		},
		{
			name:    "復元テスト 6",
			value:   "000",
			want:    "0",
			wantErr: false,
		},
		{
			name:    "復元テスト 7",
			value:   "0000",
			want:    "0",
			wantErr: false,
		},
		{
			name:    "復元テスト 8",
			value:   "111110000",
			want:    "111110000",
			wantErr: false,
		},
		{
			name:    "復元テスト 9",
			value:   "FFFFFFFFFFFFF",
			want:    "FFFFFFFFFFFFF",
			wantErr: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b, err := NewBitsHex(test.value)
			if !test.wantErr && err != nil {
				t.Error(err)
			} else if test.wantErr && err == nil {
				t.Errorf("want err")
			} else if !test.wantErr {
				if got := b.GetValue().(string); got != test.want {
					t.Errorf("want %v got %v", test.want, got)
				}
			}

		})
	}

	for _, test := range []struct {
		name, value string
		want        BitsBinaryValue
		wantErr     bool
	}{
		{
			name:    "toBinary 1",
			value:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "toBinary 2",
			value:   "000000000010",
			want:    "00010000",
			wantErr: false,
		},
		{
			name:    "toBinary 3",
			value:   "1011",
			want:    "0001000000010001",
			wantErr: false,
		},
		{
			name:    "toBinary 4",
			value:   "hhh",
			wantErr: true,
		},
		{
			name:    "toBinary 5",
			value:   "abc",
			want:    "101010111100",
			wantErr: false,
		},
		{
			name:    "toBinary 6",
			value:   "000",
			want:    "0000",
			wantErr: false,
		},
		{
			name:    "toBinary 7",
			value:   "D3FF",
			want:    "1101001111111111",
			wantErr: false,
		},
		{
			name:    "toBinary 8",
			value:   "FFFFFFFFFFFFF",
			want:    "1111111111111111111111111111111111111111111111111111",
			wantErr: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b, err := NewBitsHex(test.value)
			if !test.wantErr && err != nil {
				t.Error(err)
			} else if test.wantErr && err == nil {
				t.Errorf("want err")
			} else if !test.wantErr {
				if got := b.value; got != test.want {
					t.Errorf("want %v got %v", test.want, got)
				}
			}

		})
	}

	for _, test := range []struct {
		name        string
		value, want BitsBinaryValue
		wantErr     bool
	}{
		{
			name:    "toHex 1",
			value:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "toHex 2",
			value:   "000000000010",
			want:    "2",
			wantErr: false,
		},
		{
			name:    "toHex 3",
			value:   "1101001111111111",
			want:    "D3FF",
			wantErr: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := BitsHex{BitsBinary{test.value}}
			if got := b.toHex(); got != test.want {
				t.Errorf("want %v got %v", test.want, got)
			}
		})
	}

	for _, test := range []struct {
		name, value         string
		onValue, offValue   int
		wantHex2, wantHex16 BitsBinaryValue
		wantErr             bool
	}{
		{
			name:      "on 1",
			value:     "",
			onValue:   1,
			wantHex2:  "1",
			wantHex16: "1",
			wantErr:   false,
		},
		{
			name:      "on 2",
			value:     "",
			onValue:   3,
			wantHex2:  "100",
			wantHex16: "4",
			wantErr:   false,
		},
		{
			name:      "on 3",
			value:     "e", //001110
			onValue:   6,
			wantHex2:  "101110",
			wantHex16: "2E",
			wantErr:   false,
		},
		{
			name:      "on 4",
			value:     "2e", //101110
			onValue:   10,
			wantHex2:  "1000101110",
			wantHex16: "22E",
			wantErr:   false,
		},
		{
			name:      "off 1",
			value:     "",
			offValue:  1,
			wantHex2:  "0",
			wantHex16: "0",
			wantErr:   false,
		},
		{
			name:      "off 2",
			value:     "",
			offValue:  3,
			wantHex2:  "0",
			wantHex16: "0",
			wantErr:   false,
		},
		{
			name:      "off 3",
			value:     "e", //001110
			offValue:  2,
			wantHex2:  "1100",
			wantHex16: "C",
			wantErr:   false,
		},
		{
			name:      "off 4",
			value:     "2e", //101110
			offValue:  6,
			wantHex2:  "1110",
			wantHex16: "E",
			wantErr:   false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			b, err := NewBitsHex(test.value)
			if !test.wantErr && err != nil {
				t.Error(err)
			} else if test.wantErr && err == nil {
				t.Errorf("want err")
			} else if err == nil {

				if test.onValue > 0 {
					if err := b.On(test.onValue); err != nil {
						t.Error(err)
					}
				} else if test.offValue > 0 {
					if err := b.Off(test.offValue); err != nil {
						t.Error(err)
					}
				}
				if got := b.value; got != test.wantHex2 {
					t.Errorf("want %v got %v", test.wantHex2, got)
				}
				if got := b.toHex(); got != test.wantHex16 {
					t.Errorf("want %v got %v", test.wantHex2, got)
				}
			}
		})
	}

}
