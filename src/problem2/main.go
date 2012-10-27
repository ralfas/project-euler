package main

import "fmt"

func main () {

    sum := 0

    i := 1
    j := 1

    for j < 4000000 {

        if j%2 == 0 {
            sum += j
        }

        i, j = j, j+i
    }

    fmt.Println(sum)
}