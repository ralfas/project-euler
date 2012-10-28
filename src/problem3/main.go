package main

import (
    "fmt"
    "math"
    "sort"
)

func factors (number uint64, factors_ch chan<- []uint64) {

    factors := []uint64 {}

    for i := uint64(1); i <= uint64(math.Sqrt(float64(number))); i++ {

        if number%i == 0 {
            factors = append(factors, i)
        }
    }

    factors_ch <- append(factors, number)
}

func overlap (arrA, arrB []uint64) (overlap []uint64) {

    for i := 0; i < len(arrA); i++ {

        index := sort.Search(len(arrB), func(j int) bool {
            return arrB[j] >= arrA[i]
        })

        if index < len(arrB) && arrB[index] == arrA[i] {
            overlap = append(overlap, arrA[i])
        }

    }
    
    return
}

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate (ch chan<- uint64) {
    for i := uint64(2); ; i++ {
        ch <- i // Send 'i' to channel 'ch'.
    }
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter (src <-chan uint64, dst chan<- uint64, prime uint64) {
    for i := range src { // Loop over values received from 'src'.
        if i%prime != 0 {
            dst <- i // Send 'i' to channel 'dst'.
        }
    }
}

// The prime sieve: Daisy-chain filter processes together.
func primes (number uint64, primes_ch chan<- []uint64) {

    ch := make(chan uint64) // Create a new channel.
    go generate(ch) // Start generate() as a subprocess.

    primes := []uint64 {}
    counter := 0

    for {
        prime := <-ch

        if prime > uint64(math.Sqrt(float64(number))) {
            primes_ch <- primes
        }

        primes = append(primes, prime)
        if counter == 50 {
            counter = 0
            fmt.Println("Latest prime: ", prime)
        } else {
            counter++
        }

        ch1 := make(chan uint64)
        go filter(ch, ch1, prime)
        ch = ch1
    }
}

func main () {

    var target uint64 = 600851475143
    
    primes_ch := make(chan []uint64) // channel where primes are sent
    factors_ch := make(chan []uint64) // channel where factors are sent

    go factors(target, factors_ch)
    go primes(target, primes_ch)

    fmt.Println(overlap(<-primes_ch, <-factors_ch))
}