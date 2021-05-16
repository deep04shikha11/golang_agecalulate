// this program will calculate your age 
// first import packeges
package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
// start writing function 
func main(){
	// this will include read module
	scanner :=bufio.NewScanner(os.Stdin)
	// this will show prompt message to user 
	fmt.Println("Enter your birth year: ")
	// this will read actual line as string type by user 
	scanner.Scan()
	// this will convert text to integer to perfrom age calculation 
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// following will show answer after calculation 
	fmt.Println("Your age : ", 2021-input )
}