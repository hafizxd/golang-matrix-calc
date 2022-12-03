package main

import (
	"fmt"
	"errors"
)

type Matrix struct {
	raw    [][]float64
	Row    int
	Column int
}

type AddOrSubstractHandler func(float64, float64) float64

var ErrMatrixesNotCoincided error = errors.New("Ukuran dari dua matriks harus sama")

func scanTwoMatrix() (*Matrix, *Matrix) {
	fmt.Println("* Matrix 1")
	m1 := ScanMatrix()
	fmt.Println("*")
	fmt.Println("* Matrix 2")
	m2 := ScanMatrix()

	return m1, m2
}

func ScanMatrix() *Matrix {
	var row, col int

	fmt.Print("Input jumlah baris : ")
	fmt.Scanln(&col)
	fmt.Print("Input jumlah kolom : ")
	fmt.Scanln(&row)

	arr := make([][]float64, col)

	for i := 0; i < col; i++ {
		arr[i] = make([]float64, row)

		for j := 0; j < row; j++ {
			fmt.Print("Input nilai baris-", i+1, " kolom-", j+1, " : ")
			fmt.Scanln(&arr[i][j])
		}
	}

	m := GenerateMatrix(row, col)
	m.SetData(arr)

	return m
}

func (m *Matrix) String() string {
	arr := ""
	for i := 0; i < m.Column; i++ {
		arr += fmt.Sprintf("  %v\n", m.raw[i])
	}
	return fmt.Sprintf("[\n%v]", arr)
}

func GenerateMatrix(row int, col int) *Matrix {
	arr := make([][]float64, col)

	for i := range arr {
		arr[i] = make([]float64, row)
	}

	return &Matrix{
		raw: arr,
		Row: row,
		Column: col,
	}
}

func (m *Matrix) SetData(data [][]float64) {
	m.raw = data
	m.Row = len(data[0])
	m.Column = len(data)
}

func isSameSize(m1 *Matrix, m2 *Matrix) bool {
	return m1.Row == m2.Row && m1.Column == m2.Column
}

func isSameRowCol(m1 *Matrix, m2 *Matrix) bool {
	return m1.Row == m2.Column
}

func (m *Matrix) setRow(index int, data []float64) {
	if len(data) == m.Row {
		m.raw[index] = data
		return
	}
	
	for i := 0; i < m.Row; i++ {
		if i >= len(data) {
			m.raw[index][i] = 0
			continue
		}
		m.raw[index][i] = data[i]
	}
}

func (m *Matrix) getColumn(index int) (arr []float64) {
	arr = make([]float64, m.Column)
	for i, r := range m.raw {
		arr[i] = r[index]
	}
	return
}

func (m *Matrix) addOrSubstract(mat *Matrix, handler AddOrSubstractHandler) *Matrix {
	result := GenerateMatrix(m.Row, m.Column)

	for i := 0; i < m.Column; i++ {
		rowSum := make([]float64, m.Row)
		r1 := m.raw[i]
		r2 := mat.raw[i]

		for j := 0; j < m.Row; j++ {
			rowSum[j] = handler(r1[j], r2[j])
		}

		result.setRow(i, rowSum)
	}

	return result
}

func (m *Matrix) Plus(mat *Matrix) (*Matrix, error) {
	if !isSameSize(m, mat) {
		return nil, ErrMatrixesNotCoincided
	}
	
	return m.addOrSubstract(mat, func (a float64, b float64) float64 {
		return a + b
	}), nil;
}

func (m *Matrix) Minus(mat *Matrix) (*Matrix, error) {
	if !isSameSize(m, mat) {
		return nil, ErrMatrixesNotCoincided
	}
	
	return m.addOrSubstract(mat, func (a float64, b float64) float64 {
		return a - b
	}), nil;
}

func (m *Matrix) ScalarMultiply(k float64) *Matrix {
	result := GenerateMatrix(m.Row, m.Column)
	for i := 0; i < m.Column; i++ {
		rowSum := make([]float64, m.Row)

		for j, val := range m.raw[i] {
			rowSum[j] = val * k
		}
		result.setRow(i, rowSum)
	}

	return result
}

func (m *Matrix) Transpose() (matrix *Matrix) {
	matrix = GenerateMatrix(m.Column, m.Row)

	for i := 0; i < m.Row; i++ {
		matrix.setRow(i, m.getColumn(i))
	}

	return
}