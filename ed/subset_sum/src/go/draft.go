package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subsetSum(arr []int, n, target, index, currSum int) bool {
	if currSum == target {
		return true
	}
	if currSum > target || index == n {
		return false
	}
	if subsetSum(arr, n, target, index+1, currSum+arr[index]) {
		return true
	}
	if subsetSum(arr, n, target, index+1, currSum) {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	k, _ := strconv.Atoi(firstLine[1])

	if !scanner.Scan() {
		return
	}
	secondLine := strings.Fields(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(secondLine[i])
	}

	if subsetSum(arr, n, k, 0, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}