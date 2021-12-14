package base62

import (
	"math"
	"strings"
)

const (
	BASE62 = "0123456789abcdefghijklmnoporstuvwxyzABCDEFGHIJKLMNOPORSTUVWXYZ"
	LENGTH = 62
)

var (
	ma = map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 39, "E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61,
	}
) // encode int to base62 string.
func EncodeBase62(v int) string {
	if v == 0 {
		return "0"
	}
	var result []byte
	for v > 0 {
		quotient := v / LENGTH
		residue := v % LENGTH
		result = append(result, BASE62[residue])
		v = quotient
	}
	return string(result)
}

// decode string to int
func DecodeBase62(v string) int {
	strings.TrimSpace(v)
	num := 0
	for k, vi := range v {
		num += ma[string(vi)] * int(math.Pow(LENGTH, float64(k)))
	}
	return num
}
