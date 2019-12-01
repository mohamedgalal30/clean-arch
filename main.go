package main

import (
	"fmt"
	"os"
	"ticket/interfaces/api"
)

func main() {
	switch len(os.Args) > 1 {
	case true:
		switch os.Args[1] {
		case "api":
			fmt.Println(">>>>>>>>>>ttttttt[Ticket API Running]")
			api.Main()
			break
		case "ops":
			removeOPSFromOsArgs()
			fmt.Println(">>>>>>>>>>>>>>[OPS Running]")
			//ops.Main()
			break
		}
		break
	case false:
		fmt.Println(">>>>>>>>>>>>>>>rrrrr[Ticket API Running]")
		api.Main()

	}
}

func removeOPSFromOsArgs() {
	os.Args = append(os.Args[:1], os.Args[2:]...)
}
