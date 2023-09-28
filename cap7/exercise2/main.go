package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Printf("%#U\n#%U\n#%U\n%v", x, x, x, x)
		fmt.Printf("%#U\n#%U\n#%U\n%v", x, x, x, x)
		fmt.Printf("%#U\n#%U\n#%U\n%v", x, x, x, x)
	}

}
