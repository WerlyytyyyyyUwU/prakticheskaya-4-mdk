package main

import "fmt"

func read(n int) [][]float64 {
	m := make([][]float64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&m[i][j])
		}
	}
	return m
}

func printM(m [][]float64) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%8.3f", m[i][j])
		}
		fmt.Println()
	}
}

func add(a, b [][]float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			r[i][j] = a[i][j] + b[i][j]
		}
	}
	return r
}

func scalar(a [][]float64, k float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			r[i][j] = a[i][j] * k
		}
	}
	return r
}

func mul(a, b [][]float64) [][]float64 {
	n := len(a)
	r := make([][]float64, n)
	for i := 0; i < n; i++ {
		r[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for t := 0; t < n; t++ {
				s += a[i][t] * b[t][j]
			}
			r[i][j] = s
		}
	}
	return r
}

/* Новые функции для разделения логики */

func chooseOp() int {
	var op int
	fmt.Println("Выберите операцию:")
	fmt.Println("1 = сложение")
	fmt.Println("2 = умножение матрицы на число")
	fmt.Println("3 = умножение матриц")
	fmt.Println("0 = выход")
	fmt.Scan(&op)
	return op
}

func chooseSize() int {
	var n int
	fmt.Println("Размер (2 или 3):")
	fmt.Scan(&n)
	return n
}

func handleAdd(n int) {
	fmt.Println("Матрица A:")
	A := read(n)
	fmt.Println("Матрица B:")
	B := read(n)
	fmt.Println("Результат A + B:")
	printM(add(A, B))
}

func handleScalar(n int) {
	fmt.Println("Матрица A:")
	A := read(n)
	fmt.Println("Число k:")
	var k float64
	fmt.Scan(&k)
	fmt.Println("Результат A * k:")
	printM(scalar(A, k))
}

func handleMul(n int) {
	fmt.Println("Матрица A:")
	A := read(n)
	fmt.Println("Матрица B:")
	B := read(n)
	fmt.Println("Результат A * B:")
	printM(mul(A, B))
}

func main() {
	op := chooseOp()
	if op == 0 {
		return
	}
	n := chooseSize()
	if n != 2 && n != 3 {
		fmt.Println("Неподдерживаемый размер")
		return
	}
	switch op {
	case 1:
		handleAdd(n)
	case 2:
		handleScalar(n)
	case 3:
		handleMul(n)
	default:
		fmt.Println("Неизвестная операция")
	}
}
