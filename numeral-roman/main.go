package main

import (
	"strings"
)

type RomanNumeral struct {
	val    int
	symbol string
}

var romanNumerals RomanNumerals = []RomanNumeral{
	{val: 1000, symbol: "M"},
	{val: 900, symbol: "CM"},
	{val: 500, symbol: "D"},
	{val: 400, symbol: "CD"},
	{val: 100, symbol: "C"},
	{val: 90, symbol: "XC"},
	{val: 50, symbol: "L"},
	{val: 40, symbol: "XL"},
	{val: 10, symbol: "X"},
	{val: 9, symbol: "IX"},
	{val: 5, symbol: "V"},
	{val: 4, symbol: "IV"},
	{val: 1, symbol: "I"},
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	roman := string(symbols)
	for _, rn := range r {
		if rn.symbol == roman {
			return rn.val
		}
	}
	return 0
}

func ConvertToRoman(num int) string {
	var res strings.Builder

	for _, rs := range romanNumerals {
		for num >= rs.val {
			res.WriteString(rs.symbol)
			num -= rs.val
		}
	}

	return res.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) {
			if val := romanNumerals.ValueOf(symbol, roman[i+1]); val != 0 {
				total += val
				i++
				continue
			}
		}
		total += romanNumerals.ValueOf(symbol)
	}

	return total
}
