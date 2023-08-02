package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertRomanToArabic(a string) (int, error) {
	switch strings.ToUpper(a) {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, fmt.Errorf("Строку %s невозможно преобразовать в число, либо оно не удовлетворяет условиям задачи", a)
	}
} // функция преобразования римской цифры от 1 до 10 в арабскую

func convertArabicToRoman(a int) string {
	decimalArabic := a / 10
	unitArabic := a - (10 * decimalArabic)
	var decimalRoman, unitRoman, romanNumber string
	switch unitArabic {
	case 1:
		unitRoman = "I"
	case 2:
		unitRoman = "II"
	case 3:
		unitRoman = "III"
	case 4:
		unitRoman = "IV"
	case 5:
		unitRoman = "V"
	case 6:
		unitRoman = "VI"
	case 7:
		unitRoman = "VII"
	case 8:
		unitRoman = "VIII"
	case 9:
		unitRoman = "IX"
	default:
		unitRoman = ""
	}
	switch decimalArabic {
	case 1:
		decimalRoman = "X"
	case 2:
		decimalRoman = "XX"
	case 3:
		decimalRoman = "XXX"
	case 4:
		decimalRoman = "XL"
	case 5:
		decimalRoman = "L"
	case 6:
		decimalRoman = "LX"
	case 7:
		decimalRoman = "LXX"
	case 8:
		decimalRoman = "LXXX"
	case 9:
		decimalRoman = "XC"
	case 10:
		decimalRoman = "C"
	default:
		decimalRoman = ""
	}
	if decimalArabic == 0 {
		romanNumber = unitRoman
	} else if decimalArabic == 10 && unitArabic == 0 {
		romanNumber = decimalRoman
	} else {
		romanNumber = decimalRoman + unitRoman
	}
	return romanNumber

} // функция преобразования арабского числа к римскому

func isRomanNumber(a string) bool {
	_, err := convertRomanToArabic(a)
	if err != nil {
		return false
	}
	return true
} // функция проверки является ли число римским

func convertStringToNumber(a string) (int, error) {
	number, err := strconv.Atoi(a)
	if err != nil {
		number, err = convertRomanToArabic(a)
		if err != nil {
			return 0, err
		}
	}
	if number > 10 || number < 1 {
		return 0, fmt.Errorf("Число должно быть от 1 до 10")
	}
	return number, nil

} // функция преобразования строки к числу

func finalOperationRoman(a, b int, c string) (string, error) {
	switch c {
	case "+":
		return convertArabicToRoman(a + b), nil
	case "-":
		if a-b < 1 {
			return "0", fmt.Errorf("В римской системе нет отрицательных чисел")
		}
		return convertArabicToRoman(a - b), nil
	case "*":
		return convertArabicToRoman(a * b), nil
	case "/":
		if a/b < 1 {
			return "0", fmt.Errorf("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
		}
		return convertArabicToRoman(a / b), nil
	default:
		return "0", fmt.Errorf("Калькулятор умеет выполнять только операции сложения, вычитания, умножения и деления с двумя числами")
	}

} // функция выполнения арифметического действия с арабскими числами

func finalOperationArabic(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("Калькулятор умеет выполнять только операции сложения, вычитания, умножения и деления с двумя числами")
	}

} // функция выполнения арифметического действия с римскими числами

func main() {
	fmt.Println("Что посчитать?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialString := scanner.Text() //водим строку
	splitString := strings.Split(initialString, " ")
	if len(splitString) != 3 {
		fmt.Println("Формат записи не соответствует заданию")
		return
	}
	firstNumber := splitString[0]
	operationSign := splitString[1]
	secondNumber := splitString[2]
	numberOne, err1 := convertStringToNumber(firstNumber)
	numberTwo, err2 := convertStringToNumber(secondNumber)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	isFirstNumberRoman := isRomanNumber(firstNumber)
	isSecondNumberRoman := isRomanNumber(secondNumber)
	if (isFirstNumberRoman == true && isSecondNumberRoman == false) || (isFirstNumberRoman == false && isSecondNumberRoman == true) {
		fmt.Println("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
		return
	}
	if isFirstNumberRoman == true && isSecondNumberRoman == true {
		result, err3 := finalOperationRoman(numberOne, numberTwo, operationSign)
		if err3 != nil {
			fmt.Println(err3)
		} else {
			fmt.Println(result)
		}
		return
	}
	result, err3 := finalOperationArabic(numberOne, numberTwo, operationSign)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(result)
	}

}
