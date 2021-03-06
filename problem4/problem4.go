package main

import "fmt"

func main() {
    var total, lp int
    var ip bool
    for d := 999; d > 1; d-- {
        for i := d; i > 1; i-- {
	    total = d * i
            ip = is_palindrome(total)
		// fmt.Printf("%d, %v\n", total, ip)
	    if ip == true {
		if lp < total {
			lp = total 
		}
	    }
        }
    }
    fmt.Printf("Done, largest palindrome is %d.\n", lp)
}

func is_palindrome(t int) bool {
// Get the irdividual digits of total
        var MAX_ARRAY [6]int
        var size int
        number := t
        for i:= 0; number > 0; i++ {
            MAX_ARRAY[i] = number % 10
            size++
            number /= 10
        }

	switch size {
	case 6: 
        if MAX_ARRAY[0] == MAX_ARRAY[5] && MAX_ARRAY[1] == MAX_ARRAY[4] && MAX_ARRAY[2] == MAX_ARRAY[3] {
                    return true
        }
        
	case 5:
	if MAX_ARRAY[0] == MAX_ARRAY[4] && MAX_ARRAY[1] == MAX_ARRAY[3] {
		return true
        }

        case 4:
	if MAX_ARRAY[0] == MAX_ARRAY[3] && MAX_ARRAY[1] == MAX_ARRAY[2] {
		return true
	}

	case 3:
	if MAX_ARRAY[0] == MAX_ARRAY[2] {
		return true
	}

        case 2:
        if MAX_ARRAY[0] == MAX_ARRAY[1] {
		return true
	}

        case 1: return true
	
	}
	return false



        
        /*	
	for z := 0; z < len(MAX_ARRAY); z++ {
	    fmt.Printf("%d", MAX_ARRAY[z])
            // fmt.Printf("%d %d", MAX_ARRAY[0], MAX_ARRAY[len(MAX_ARRAY) - 
        }
	fmt.Printf("\n")
	*/

}
