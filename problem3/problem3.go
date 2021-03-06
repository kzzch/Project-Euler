// The prime factors of 13195 are 5, 7, 13, and 29
// What is the larges prime factor of the number 600 851 475 143
// 775146 = upper bound
package main

import "fmt"
import "math"

func main() {
    var n, bound int64
    // n = 13195 // Test case 
    n = 600851475143
    
    // Calculate the upper bound of our search as
    // the square root of n
    bound = int64(math.Sqrt(float64(n)))
    
    // Create an array to hold our generated primes
    
    var is_prime [775146]bool
    // Create a slice to hold the prime factors
    prime_factors := make([]int64, 10)
     
    // Generate primes using the Sieve of Eratosthenes to 
    // remove the non-primes from our array
    fmt.Printf("Generating primes...\n")
   
    // Initialize the array 
    for z := int64(2); z < bound; z++ {
        is_prime[z] = true
    }

    // Sieve out the non-primes
    for step_size := int64(2); step_size < bound; step_size++ {
	for x := step_size + step_size; x < bound; x += step_size {
         	is_prime[x] = false
        }
    }
    fmt.Printf("Prime generation complete.\n")

    // Factorize using trial division
    fmt.Printf("Factorizing n...\n")
    for p := int64(2); p < bound; p++ {
        if is_prime[p] {
	    if int64(n) % p == 0 {
		prime_factors = append(prime_factors, p)
	    }
        }
    }

    fmt.Printf("Factorization complete...\n")
    fmt.Printf("Prime factors of %d are:\n", n)
    for pf := 2; pf < len(prime_factors); pf++ {
        if prime_factors[pf] != 0 {
            fmt.Printf("%d\n", prime_factors[pf])
        }
    }
}
