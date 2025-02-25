package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var n int
	fmt.Fscan(in, &n)

	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &input[i])
	}

	in.ReadString('\n')
	outputStr, _ := in.ReadString('\n')
	outputStr = strings.TrimRight(outputStr, "\n")
	if strings.HasPrefix(outputStr, " ") ||
		strings.HasSuffix(outputStr, " ") ||
		strings.Contains(outputStr, "  ") {
		fmt.Fprintln(out, "no")
		return
	}

	output := strings.Fields(outputStr)
	if len(output) != n {
		fmt.Fprintln(out, "no")
		return
	}

	outputInts := make([]int, n)
	for i, s := range output {
		if hasLeadingZero(s) {
			fmt.Fprintln(out, "no")
			return
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintln(out, "no")
			return
		}
		outputInts[i] = num
	}

	if !slices.IsSorted(outputInts) {
		fmt.Fprintln(out, "no")
		return
	}

	slices.Sort(input)
	for i := 0; i < n; i++ {
		if input[i] != outputInts[i] {
			fmt.Fprintln(out, "no")
			return
		}
	}

	fmt.Fprintln(out, "yes")
}

func hasLeadingZero(s string) bool {
	return len(s) > 1 && (s[0] == '0' || (s[0] == '-' && s[1] == '0'))
}
