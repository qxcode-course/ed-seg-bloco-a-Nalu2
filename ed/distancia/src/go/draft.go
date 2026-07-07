
package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Valid(seq []rune, idx int, l int, d rune) bool {
    start:= idx - l
    if start < 0 {
        start = 0
    }
    for i := start; i < idx; i++ {
        if seq[i] == d {
            return false
        }
    }
    end := idx + l
    if end >= len(seq) {
        end = len(seq) - 1
    }
    for i := idx + 1; i <= end; i++ {
		if seq[i] == d {
			return false
		}
	}
	return true
}

func solve(seq []rune, idx int, l int) bool {
    if idx == len(seq) {
        return true
    }
    if seq[idx] != '.' {
        return solve(seq, idx+1, l)
    }
    for d := 0; d <= l; d++ {
        charr := rune('0' + d); 
        if Valid(seq, idx, l, charr) {
            seq[idx] = charr
            if solve(seq, idx + 1, l) {
                return true
            }
            seq[idx] = '.'
        }
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    line := scanner.Text()
    seq := []rune(line)
    if !scanner.Scan() {
        return
    }
    l, _ := strconv.Atoi(scanner.Text())
    if solve(seq, 0, l) {
        fmt.Println(string(seq))
    }
}