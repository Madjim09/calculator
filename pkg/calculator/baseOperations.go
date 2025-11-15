package calculations

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// RunCalculatorSession запускает калькулятор
func RunCalculatorSession() bool {
	val1, val2, oper := readUserInput()
	res := calculation(val1, val2, oper)
	return outputResult(res, val1, val2, oper)
}

// readUserInput читает пользовательский ввод
func readUserInput() (float64, float64, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите левый операнд: ")
	val1 := readFloat(scanner)

	fmt.Print("Введите операцию: ")
	op := readValidOperand(scanner)

	fmt.Print("Введите правый операнд: ")
	val2 := readFloat(scanner)

	return val1, val2, op
}

// readFloat читает операнды
func readFloat(scanner *bufio.Scanner) float64 {
	for {
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		if text == "" {
			fmt.Print("Пустой ввод. Повторите: ")
			continue
		}

		val, err := parseValidFloat(text)
		if err != nil {
			fmt.Printf("Неверный операнд: %s\n", err)
			fmt.Print("Введите число еще раз (пример: 1, -1, 3.14): ")
			continue
		}

		return val
	}
}

// parseValidFloat парсит и валидирует операнды
func parseValidFloat(text string) (float64, error) {
	val, err := strconv.ParseFloat(text, 64)

	if err != nil {
		return 0, fmt.Errorf("введено не число")
	}
	if math.IsInf(val, 0) {
		return 0, fmt.Errorf("бесконечность не поддерживается")
	}
	if math.IsNaN(val) {
		return 0, fmt.Errorf("NaN не поддерживается")
	}

	return val, nil
}

// readValidOperand читает и валидирует оператор
func readValidOperand(scanner *bufio.Scanner) string {
	validOps := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	for {
		scanner.Scan()
		op := strings.TrimSpace(scanner.Text())
		if validOps[op] {
			return op
		}

		fmt.Println("Неверный оператор.")
		fmt.Print("Введите оператор еще раз (+, -, *, /): ")
	}
}

// calculation выполняет операции
func calculation(val1, val2 float64, operator string) float64 {
	for {
		switch operator {
		case "+":
			return val1 + val2
		case "-":
			return val1 - val2
		case "*":
			return val1 * val2
		case "/":
			for val2 == 0 {
				fmt.Println("Деление на ноль невозможно.")
				fmt.Print("Введите правый операнд еще раз: ")
				scanner := bufio.NewScanner(os.Stdin)
				val2 = readFloat(scanner)
			}
			return val1 / val2
		}
	}
}

// outputResult выводит результат и запрашивает продолжение
func outputResult(res, val1, val2 float64, op string) bool {
	fmt.Printf("%.3f %s %.3f = %.3f\n", val1, op, val2, res)

	fmt.Print("Хотите продолжить? [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch answer {
		case "y", "yes", "да", "д":
			fmt.Println()
			return true
		case "n", "no", "нет", "н":
			return false
		default:
			fmt.Println("Неверный ввод.")
			fmt.Print("Введите y (да) или n (нет): ")
		}
	}
}
