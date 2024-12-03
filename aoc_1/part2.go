package aoc1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func P2() {
    
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


    m := make(map[int]int)
    // counter
    for i:=0; i < len(second); i++ {
        val := second[i]
        m[val] += 1
    }

    // loop over keys in map
    for key := range m {
        m[key] *= key
    }

    // get similarity score
    score := 0
    for i := range len(first) {
        score +=  m[first[i]]
    }

    fmt.Println("score=%d",score)

    duration := time.Since(start)
    fmt.Println("execution time=%f",duration)
}

