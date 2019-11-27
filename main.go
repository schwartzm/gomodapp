package main 

import (
	"fmt"
	sayings "github.com/schwartzm/gomodlib"
)

func main() {
	fmt.Printf("Happy saying: %q\n", sayings.Happy())
	fmt.Printf("Sad saying: %q\n", sayings.Sad())
	fmt.Printf("Joyous saying: %q\n", sayings.Joyous())
	fmt.Printf("Mad saying: %q\n", sayings.Mad())
	fmt.Printf("A Hello saying: %q\n", sayings.Hello())
	fmt.Println("Have a great day. Bye.")
}
