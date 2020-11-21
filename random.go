package main

import (
	"fmt"
	"math/rand"
)

// Import the math/rand package

func main() {

	for i := 1; i <= 5; i++ {
		// Generate and print a random integer from 0 to 9
		fmt.Println(rand.Intn(10))
	}
}

// using time and rand

package main

import "fmt"
import "math/rand"
// Import the time package
import "time"

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().Unix())

    for i := 1; i <= 5; i++ {
        fmt.Println(rand.Intn(10))
    }
}
