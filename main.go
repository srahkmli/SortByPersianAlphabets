package SortByPersianAlphabets

import "sort"

var persianAlphabet = []rune{
	'ا', 'ب', 'پ', 'ت', 'ث', 'ج', 'چ', 'ح', 'خ', 'د', 'ذ', 'ر', 'ز', 'ژ',
	'س', 'ش', 'ص', 'ض', 'ط', 'ظ', 'ع', 'غ', 'ف', 'ق', 'ک', 'گ', 'ل', 'م',
	'ن', 'و', 'ه', 'ی',
}

// getIndex returns the index of a rune in the Persian alphabet.
// If the rune is not found, it returns -1.
func getIndex(r rune) int {
	for i, v := range persianAlphabet {
		if v == r {
			return i
		}
	}
	return -1
}

// Less returns true if rune a should be sorted before rune b.
func Less(a, b rune) bool {
	indexA, indexB := getIndex(a), getIndex(b)
	return indexA < indexB
}

// SortPersianStrings sorts a slice of Persian strings using the Persian alphabet order.
func SortPersianStrings(strings []string) {
	sort.SliceStable(strings, func(i, j int) bool {
		strA, strB := []rune(strings[i]), []rune(strings[j])

		minLen := len(strA)
		if len(strB) < minLen {
			minLen = len(strB)
		}

		for k := 0; k < minLen; k++ {
			if strA[k] != strB[k] {
				return Less(strA[k], strB[k])
			}
		}

		return len(strA) < len(strB)
	})
}
