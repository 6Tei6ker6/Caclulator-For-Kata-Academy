package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
	"strconv"
)

func checkslice(inp []string)error {
	var operator string
	var foundstr int
	operand := make([]string, 0)
	arabic := make([]int, 0)
	roman := make([]string, 0)

	switch{
	case len(inp) != 3:
		return errors.New("Неверно введены данные")
	case len(inp[1]) != 1:
		return errors.New("Неверно введен оператор")
	case inp[1] == "+" || inp[1] == "-" || inp[1] == "/" || inp[1] == "*":
		operator = inp[1]
		operand = append(operand, inp[0])
		operand = append(operand, inp[2])
	default: 
		return errors.New("Неверно введен оператор")
	}

	for _, idk := range operand{
		num, err := strconv.Atoi(idk)
		if err != nil{
				foundstr++
				roman = append(roman, idk)
		} else {
			arabic = append(arabic, num)
		}
	}

	switch foundstr {
		case 1:
			return errors.New("Использованы разные системы исчисления или неправильно введены операнды")
		case 0:
			numCheck := arabic[0] > 0 && arabic[0] < 11 && arabic[1] > 0 && arabic[1] < 11
			if numCheck == true {
				switch operator {
				case "+":
					fmt.Println(arabic[0] + arabic[1])
				case "-":
					fmt.Println(arabic[0] - arabic[1])
				case "/":
					fmt.Println(arabic[0] / arabic[1])
				case "*":
					fmt.Println(arabic[0] * arabic[1])
				}
			} else {
				return errors.New("Калькулятор работает только с целыми числами от 1 до 10")
			}
		case 2:
			op1 := romanToInt(roman[0])
			op2 := romanToInt(roman[1])
			numCheck := op1 > 0 && op1 < 11 && op2 > 0 && op2 < 11
			if numCheck == true {
			switch operator {
			case "+":
				fmt.Println(intToRoman(op1 + op2))
			case "-":
				if op1 <= op2{
					return errors.New("Ошибка, в римской системе исчисления нету нуля и отрицательных чисел")
				}
				fmt.Println(intToRoman(op1 - op2))
			case "/":
				fmt.Println(intToRoman(op1 / op2))
			case "*":
				fmt.Println(intToRoman(op1 * op2))
			}
			} else {
			return errors.New("Калькулятор работает только с целыми числами от 1 до 10")
		}
	}
	return nil
}

func calc()error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в калькулятор для Kata Academy")
	for {
		fmt.Println("Введите значение")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		slofstr := strings.Split(strings.ToUpper(strings.TrimSpace(text)), " ")
		err = checkslice(slofstr)
		if err != nil{
			return err
		}
	}
}

func main() {
	err := calc()
	if err != nil {
		fmt.Println(err)
	}
}


func romanToInt(s string)int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := m[s[i]]
		sign := 1
		if tmp < last {
			sign = -1
		}
		res += sign * tmp
		last = tmp
	}
	return res
}

func intToRoman(num int) string {
	symbol := [...]string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := ""
	for i := 0; num != 0; i++ {
		for num >= value[i] {
			num -= value[i]
			str += symbol[i]
		}
	}
	return str
}


