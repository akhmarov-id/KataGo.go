package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите выражение: ")
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.Replace(text, " ", "", -1)
	err := err(text)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	if strings.Contains(text, "+") {
		a := strings.Split(text, "+")
		plus(a)
	}
	if strings.Contains(text, "-") {
		a1 := strings.Split(text, "-")
		minus(a1)
	}
	if strings.Contains(text, "*") {
		a2 := strings.Split(text, "*")
		multi(a2)
	}
	if strings.Contains(text, "/") {
		a3 := strings.Split(text, "/")
		divide(a3)
	}
}

func err(a string) error {
	b := []rune(a)
	var err = fmt.Errorf("ошибка")
	for i := 0; i < len(b); i++ {
		switch {
		case strings.Count(string(b), "+")+strings.Count(string(b), "-")+strings.Count(string(b), "*")+strings.Count(string(b), "/") > 1 ||
			strings.Count(string(b), "+")+strings.Count(string(b), "-")+strings.Count(string(b), "*")+strings.Count(string(b), "/") == 0:
			{
				return err
			}
		case b[0] == '+' || b[0] == '-' || b[0] == '/' || b[0] == '*' || b[len(b)-1] == '+' || b[len(b)-1] == '-' || b[len(b)-1] == '/' || b[len(b)-1] == '*':
			{
				return err
			}
		case unicode.IsDigit(b[i]):
			{
				break
			}
		case b[i] == '+' || b[i] == '-' || b[i] == '/' || b[i] == '*':
			{
				break
			}
		case b[i] == 'I' || b[i] == 'V' || b[i] == 'X' || b[i] == 'L' || b[i] == 'C' || b[i] == 'D':
			{
				break
			}
		default:
			return err
		}
	}
	return nil
}

func plus(b []string) {
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	if (b1 == 0 || b2 == 0) && (rim(b[0]) == 0 || rim(b[1]) == 0) || (b1 > 10 || b1 < 0 || b2 > 10 || b2 < 0) {
		fmt.Println("ошибка: -калькулятор умеет работать только с арабскими или римскими цифрами одновременно\n\t-калькулятор может принимать на вход числа от 1 до 10 включительно")
		return
	}
	if b1 > 0 && b2 > 0 && b1 < 11 && b2 < 11 {
		fmt.Println(b1 + b2)
	} else if b1 == b2 {
		b1 = rim(b[0])
		b2 = rim(b[1])
		a := b1 + b2
		p := arab(a)
		fmt.Println(p)
	}
}

func minus(b []string) {
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	if (b1 == 0 || b2 == 0) && (rim(b[0]) == 0 || rim(b[1]) == 0) || (b1 > 10 || b1 < 0 || b2 > 10 || b2 < 0) {
		fmt.Println("ошибка: -калькулятор умеет работать только с арабскими или римскими цифрами одновременно\n\t-калькулятор может принимать на вход числа от 1 до 10 включительно")
		return
	}
	if b1 > 0 && b2 > 0 && b1 < 11 && b2 < 11 {
		fmt.Println(b1 - b2)
	} else if b1 == b2 {
		b1 = rim(b[0])
		b2 = rim(b[1])
		a := b1 - b2
		p := arab(a)
		if a < 1 {
			fmt.Println("ошибка: результатом работы калькулятора с римскими числами могут быть только положительные числа")
			return
		}
		fmt.Println(p)
	}
}

func multi(b []string) {
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	if (b1 == 0 || b2 == 0) && (rim(b[0]) == 0 || rim(b[1]) == 0) || (b1 > 10 || b1 < 0 || b2 > 10 || b2 < 0) {
		fmt.Println("ошибка: -калькулятор умеет работать только с арабскими или римскими цифрами одновременно\n\t-калькулятор может принимать на вход числа от 1 до 10 включительно")
		return
	}
	if b1 > 0 && b2 > 0 && b1 < 11 && b2 < 11 {
		fmt.Println(b1 * b2)
	} else if b1 == b2 {
		b1 = rim(b[0])
		b2 = rim(b[1])
		a := b1 * b2
		p := arab(a)
		fmt.Println(p)
	}
}

func divide(b []string) {
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	if (b1 == 0 || b2 == 0) && (rim(b[0]) == 0 || rim(b[1]) == 0) || (b1 > 10 || b1 < 0 || b2 > 10 || b2 < 0) {
		fmt.Println("ошибка: -калькулятор умеет работать только с арабскими или римскими цифрами одновременно\n\t-калькулятор может принимать на вход числа от 1 до 10 включительно")
		return
	}
	if b1 > 0 && b2 > 0 && b1 < 11 && b2 < 11 {
		fmt.Println(b1 / b2)
	} else if b1 == b2 {
		b1 = rim(b[0])
		b2 = rim(b[1])
		a := b1 / b2
		p := arab(a)
		if a < 1 {
			fmt.Println("ошибка: результатом работы калькулятора с римскими числами могут быть только положительные числа")
			return
		}
		fmt.Println(p)
	}
}

func rim(s string) int {
	m := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	result := 0
	if _, aMap := m[s]; aMap {
		result = m[s]
	}
	return result
}

func arab(s int) string {
	var result string
	if s == 100 {
		result = result + "C"
		return result
	}
	for s >= 90 {
		result = result + "XC"
		s -= 90
	}
	for s >= 50 {
		result = result + "L"
		s -= 50
	}
	for s >= 40 {
		result = result + "XL"
		s -= 40
	}
	for s >= 10 {
		result = result + "X"
		s -= 10
	}
	for s >= 9 {
		result = result + "IX"
		s -= 9
	}
	for s >= 5 {
		result = result + "V"
		s -= 5
	}
	for s >= 4 {
		result = result + "IV"
		s -= 4
	}
	for s >= 1 {
		result = result + "I"
		s -= 1
	}
	return result
}
