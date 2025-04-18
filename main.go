package main

import (
    "fmt"
    "os"
    "ci_cd_test_go/quadratic"
    "strconv"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Использование: ./quadratic a b c")
        os.Exit(1)
    }

    a, _ := strconv.ParseFloat(os.Args[1], 64)
    b, _ := strconv.ParseFloat(os.Args[2], 64)
    c, _ := strconv.ParseFloat(os.Args[3], 64)

    x1, x2, hasRoots := quadratic.Solve(a, b, c)

    if !hasRoots {
        fmt.Println("Действительных корней нет")
        return
    }

    if x1 == x2 {
        fmt.Printf("Один корень: x = %.2f\n", x1)
    } else {
        fmt.Printf("Два корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
    }
}