package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func avg(g []int) float64 {
	if len(g) == 0 {
		return 0
	}
	sum := 0
	for _, v := range g {
		sum += v
	}
	return float64(sum) / float64(len(g))
}

func parseGrades(s string) []int {
	out := []int{}
	for _, f := range strings.Fields(s) {
		if v, err := strconv.Atoi(f); err == nil {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	in := bufio.NewReader(os.Stdin)
	students := map[string][]int{}

	for {
		fmt.Println("\n1=Добавить  2=Список  3=Фильтр по среднему  0=Выход")
		fmt.Print("> ")
		cmd, _ := in.ReadString('\n')
		switch strings.TrimSpace(cmd) {
		case "1":
			fmt.Print("ФИО: ")
			name, _ := in.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Оценки (через пробел): ")
			line, _ := in.ReadString('\n')
			students[name] = parseGrades(line)
		case "2":
			for name, g := range students {
				fmt.Printf("%s -> %v  средний=%.2f\n", name, g, avg(g))
			}
		case "3":
			fmt.Print("Порог (например 4): ")
			line, _ := in.ReadString('\n')
			t, _ := strconv.ParseFloat(strings.TrimSpace(line), 64)
			for name, g := range students {
				if avg(g) < t {
					fmt.Printf("%s -> %.2f\n", name, avg(g))
				}
			}
		case "0":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
 