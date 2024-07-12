package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var commands = [3]string{"AVG", "SUM", "MED"}

func requestData() string {
	var line string
	fmt.Scan(&line)
	return line
}

func commandToIndex(command string) int {
	for index, c := range commands {
		if c == command {
			return index
		}
	}
	return -1
}

func checkCommand(command string) error {
	if commandToIndex(command) == -1 {
		return errors.New("Ошибка! Команда не найдена!")
	}
	return nil
}

func requestCommand() string {
	for {
		fmt.Printf("Ввберите действие [%s]:", strings.Join(commands[:], ","))
		command := strings.ToUpper(requestData())
		if err := checkCommand(command); err != nil {
			fmt.Println(err.Error())
		} else {
			return command
		}
	}
}

func convertLineToFloatArray(line string) ([]float64, error) {
	lines := strings.Split(line, ",")
	numbers := make([]float64, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		number, err := strconv.ParseFloat(lines[i], 64)
		if err != nil {
			log.Print(err.Error())
		} else {
			numbers = append(numbers, number)
		}
	}

	if len(numbers) == 0 {
		return []float64{}, errors.New("ОШИБКА! Неудалось распознать числа в последовательности!")
	}
	return numbers, nil
}

func requestLine() string {
	fmt.Print("Введите последовательность чисел через запятую: ")
	return requestData()
}

func sum(numbers []float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total
}

func avg(numbers []float64) float64 {
	return sum(numbers) / (float64)(len(numbers))
}

func med(numbers []float64) float64 {
	sort.Slice(numbers[:], func(i, j int) bool {
		return numbers[:][i] > numbers[:][j]
	})
	size := len(numbers)
	index := (int)(size / 2)
	if size%2 != 0 {
		return numbers[index]
	}
	priorIndex := index - 1
	return (numbers[priorIndex] + numbers[index]) / 2
}

func calc(command string, numbers []float64) float64 {
	switch commandToIndex(command) {
	case 0:
		return avg(numbers)
	case 1:
		return sum(numbers)
	case 2:
		return med(numbers)
	}
	return 0
}

func main() {
	command := requestCommand()
	line := requestLine()

	numbers, err := convertLineToFloatArray(line)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Printf("Результат команды [%s]: %f\n", command, calc(command, numbers))
	}
}
