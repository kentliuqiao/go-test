package main

import (
	"testing"
	"testing/quick"
)

func TestConvertToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1 to I", args: args{num: 1}, want: "I"},
		{name: "2 to II", args: args{num: 2}, want: "II"},
		{name: "3 to III", args: args{num: 3}, want: "III"},
		{name: "4 to IV", args: args{num: 4}, want: "IV"},
		{name: "5 to V", args: args{num: 5}, want: "V"},
		{name: "6 to VI", args: args{num: 6}, want: "VI"},
		{name: "9 to IX", args: args{9}, want: "IX"},
		{name: "10 to X", args: args{10}, want: "X"},
		{name: "11 to XI", args: args{11}, want: "XI"},
		{name: "18 to XVIII", args: args{18}, want: "XVIII"},
		{name: "20 to XX", args: args{20}, want: "XX"},
		{name: "39 to XXXIX", args: args{39}, want: "XXXIX"},
		{name: "40 to XL", args: args{40}, want: "XL"},
		{name: "47 to XLVII", args: args{47}, want: "XLVII"},
		{name: "49 to XLIX", args: args{49}, want: "XLIX"},
		{name: "50 to L", args: args{50}, want: "L"},
		{name: "90 to XC", args: args{num: 90}, want: "XC"},
		{name: "100 to C", args: args{100}, want: "C"},
		{name: "400 to CD", args: args{400}, want: "CD"},
		{name: "500 to D", args: args{500}, want: "D"},
		{name: "900 to CM", args: args{900}, want: "CM"},
		{name: "1000 to M", args: args{1000}, want: "M"},
		{name: "1984 to MCMLXXXIV", args: args{1984}, want: "MCMLXXXIV"},
		{name: "3999 to MMMCMXCIX", args: args{3999}, want: "MMMCMXCIX"},
		{name: "2014 to MMXIV", args: args{2014}, want: "MMXIV"},
		{name: "1006 to MVI", args: args{1006}, want: "MVI"},
		{name: "798 to DCCXCVIII", args: args{798}, want: "DCCXCVIII"},
	}
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := ConvertToRoman(tt.args.num); got != tt.want {
	// 			t.Errorf("ConvertToRoman() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	for _, tt := range tests[:] {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToArabic(tt.want); got != tt.args.num {
				t.Errorf("ConvertToArabic() = %v, want %v", got, tt.args.num)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	asseertion := func(arabic int) bool {
		roman := ConvertToRoman(arabic)
		number := ConvertToArabic(roman)
		return number == arabic
	}
	if err := quick.Check(asseertion, nil); err != nil {
		t.Error("check failed", err)
	}
}
