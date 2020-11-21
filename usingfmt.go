package main

import "fmt"

func main() {
    weather := "sunny"
    
    // Print "Today, the weather is ____" using Printf method from the fmt package
    
fmt.Printf("Today, the weather is %s", weather)
}

// for integer

package main

import "fmt"

func main() {
    month := 9
    day := 5
    
    // Print "The date today is ____/____" using fmt.Printf
    fmt.Printf("The date today is  %d/%d", month, day)
    
}

// Unlike fmt.Println,fmt.Printf won't start on a new line after the output string. As you can see below, 
// even though a second fmt.Printf was used, 
// the second output string continues right after the end of the first output with no line-break.


// By putting \n in a string, you can add a line-break. The \n character itself won't appear in the text.
//  As shown in the example below, when \n is put at the end of the string,
//   the next print statement will start on a new line.