package best_strings

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// BS изменяемая строка, все операции вызываемые от объекта меняют строку внутри,
// чтобы сделать операцию без изменения внутренней строки нужно вызвать метод от bs.str.
// Пример: strings.TrimSpace(bs.str)
type BS struct {
	str  string
	size int
}

func newBS(str string) BS {
	return BS{str, utf8.RuneCountInString(str)}
}

// Set устанавливает новую строку объекту
func (bs *BS) Set(newString string) {
	bs.str = newString
	bs.size = utf8.RuneCountInString(bs.str)
}

// ToUpperFirst переводит первую букву строки в верхний регистр.
// Например "123 \n word" перейдёт в "123 \n Word"
func (bs *BS) ToUpperFirst() {
	for i, _ := range bs.str {
		if (bs.str[i] >= 'A' && bs.str[i] <= 'Z') || (bs.str[i] >= 'a' && bs.str[i] <= 'z') {
			if bs.str[i] >= 'a' && bs.str[i] <= 'z' {
				var mini = string(bs.str[i])
				bs.str = bs.str[0:i] + strings.ToUpper(mini) + bs.str[i+1:]
			}
			return
		}
	}
}

// ToLowerFirst переводит первую букву строки в нижний регистр.
// Например "123 \n Word" перейдёт в "123 \n word"
func (bs *BS) ToLowerFirst() {
	for i, _ := range bs.str {
		if (bs.str[i] >= 'A' && bs.str[i] <= 'Z') || (bs.str[i] >= 'a' && bs.str[i] <= 'z') {
			if bs.str[i] >= 'A' && bs.str[i] <= 'Z' {
				var mini = string(bs.str[i])
				bs.str = bs.str[0:i] + strings.ToLower(mini) + bs.str[i+1:]
			}
			return
		}
	}
}

func (bs *BS) IsDigit(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}

	return unicode.IsDigit(rune(bs.str[index])), nil
}

func (bs *BS) IsNumber(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsNumber(rune(bs.str[index])), nil
}

func (bs *BS) IsSymbol(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsSymbol(rune(bs.str[index])), nil
}

func (bs *BS) IsSpace(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsSpace(rune(bs.str[index])), nil
}

func (bs *BS) IsLower(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsLower(rune(bs.str[index])), nil
}

func (bs *BS) IsUpper(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsUpper(rune(bs.str[index])), nil
}

func (bs *BS) IsLetter(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsLetter(rune(bs.str[index])), nil
}

func (bs *BS) IsTitle(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsTitle(rune(bs.str[index])), nil
}

func (bs *BS) IsMark(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsMark(rune(bs.str[index])), nil
}

func (bs *BS) IsPrint(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsPrint(rune(bs.str[index])), nil
}

func (bs *BS) IsPunct(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsPunct(rune(bs.str[index])), nil
}

func (bs *BS) IsControl(index int) (bool, error) {
	if index < 0 || index >= bs.size {
		return false, errors.New("Index out of range, index: " +
			strconv.Itoa(index) + ", string size: " + strconv.Itoa(bs.size) + "\n")
	}
	return unicode.IsControl(rune(bs.str[index])), nil
}

func (bs *BS) Replace(old, new string, n int) {
	bs.str = strings.Replace(bs.str, old, new, n)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) ReplaceAll(old, new string) {
	bs.str = strings.ReplaceAll(bs.str, old, new)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) Title() {
	bs.str = strings.Title(bs.str)
}

func (bs *BS) ToUpper() {
	bs.str = strings.ToUpper(bs.str)
}

func (bs *BS) ToLower() {
	bs.str = strings.ToLower(bs.str)
}

func (bs *BS) TrimSpace() {
	bs.str = strings.TrimSpace(bs.str)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) Trim(cutset string) {
	bs.str = strings.Trim(bs.str, cutset)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) TrimLeft(cutset string) {
	bs.str = strings.TrimLeft(bs.str, cutset)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) TrimRight(cutset string) {
	bs.str = strings.TrimRight(bs.str, cutset)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) TrimPrefix(cutset string) {
	bs.str = strings.TrimPrefix(bs.str, cutset)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) TrimSuffix(cutset string) {
	bs.str = strings.TrimSuffix(bs.str, cutset)
	bs.size = utf8.RuneCountInString(bs.str)
}

func (bs *BS) Map(mapping func(rune) rune) {
	bs.str = strings.Map(mapping, bs.str)
}
