package aoc1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func P1() {
    
    file, _ := os.Open("./aoc_1/data2.txt")
    defer file.Close()

    r := bufio.NewReader(file)
    
    var first = make([]int, 0, 1000)
    var second = make([]int, 0, 1000)
    
    start := time.Now()
    
    for {
        line, _, err := r.ReadLine()
        if len(line) > 0 {
            str := strings.Trim(string(line), " ")
            values := strings.Split(str, "   ")
            i, _ := strconv.Atoi(values[0])
            j, _ := strconv.Atoi(values[1])
            first = append(first, i)
            second = append(second, j)
        }
        if err != nil {
            break
        }
    }

    sort.Ints(first)
    sort.Ints(second)

    sum := 0
    for i := 0; i < len(first); i++ {
        diff := absDiffUint(first[i], second[i])
        sum += diff
    }
    duration := time.Since(start)

    fmt.Println("execution time=%f",duration)
    fmt.Println("sum=%d", sum)
}

