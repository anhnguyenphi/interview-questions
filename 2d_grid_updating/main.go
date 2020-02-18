package _d_grid_updating

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var result = make([]string, 0)
var num = 0

func generate(s string, left int, right int) {
	if len(s) == num * 2 {
		result = append(result, s)
		return
	}

	if left < num {
		generate(s + "(", left + 1, right)
	}

	if right < left {
		generate(s + ")", left, right + 1)
	}
}


func generateParenthesis(n int) []string {
	num = n
	generate("", 0, 0)

	return result
}

func main() {
	fmt.Println(generateParenthesis(1))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
