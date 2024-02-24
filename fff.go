package main

import "fmt"


    func sieveOfEratosthenes(n int) []int {
        prime := make([]bool, n+1)
        for i := 2; i <= n; i++ {
            prime[i] = true
        }
    
        for p := 2; p*p <= n; p++ {
            if prime[p] == true {
                for i := p * p; i <= n; i += p {
                    prime[i] = false
                }
            }
        }
    
        var primes []int
        for p := 2; p <= n; p++ {
            if prime[p] == true {
                primes = append(primes, p)
            }
        }
        return primes
    }
    
    func main() {
        n := 100 // Change it to the range you want
        primes := sieveOfEratosthenes(n)
        fmt.Println("Prime numbers up to", n, "are:")
        for _, prime := range primes {
            fmt.Print(prime, " \n")
        }
    }
    

