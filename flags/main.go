package main

import (
	"flag"
	"fmt"
)

func main() {

	// Define flags
	var (
		flag1 = flag.String("flag1", "default value", "description")
		flag2 = flag.Int("flag2", 1234, "description")
		flag3 = flag.Bool("flag3", false, "description")
	)

	// Parse flags
	flag.Parse()

	// Print flags
	fmt.Println(*flag1)
	fmt.Println(*flag2)
	fmt.Println(*flag3)
}
