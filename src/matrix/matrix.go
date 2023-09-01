package matrix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
}

func New() *Matrix {
	return &Matrix{
		data: make([][]int, 0),
	}
}

func (m *Matrix) Data() [][]int {
	return m.data
}

// O(n)
func (m *Matrix) ReadFile(path string) error {
	var (
		err  error
		file *os.File
		line []string
	)

	if file, err = os.Open(path); err != nil {
		return err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// O(n)
	for fileScanner.Scan() {
		line = strings.Split(fileScanner.Text(), ",")
		m.data = append(m.data, parseLine(line))
	}

	return nil
}

// O(m*n)
func (m *Matrix) ReplaceRepeated() {

	counter := make(map[int]int)

	// O(m*n)
	for i := range m.data {
		for j := range m.data[i] {
			counter[m.data[i][j]] += 1

			if counter[m.data[i][j]] > 1 {
				m.data[i][j] = 0
			}
		}
	}

}

// O(m*n)
func (m *Matrix) Print() {

	// O(m*n)
	for _, row := range m.data {
		fmt.Print("| ")

		for _, col := range row {
			fmt.Printf("%7d ", col)
		}

		fmt.Println("|")
	}

	fmt.Println()
}

func parseLine(line []string) []int {
	var (
		err        error
		convNumber int
		convLine   []int = make([]int, len(line))
	)

	for idx, number := range line {
		convNumber, err = strconv.Atoi(number)

		if err != nil {
			panic(err)
		}

		convLine[idx] = convNumber
	}

	return convLine
}
