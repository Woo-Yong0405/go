package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/Woo-Yong0405/nomadcoin/explorer"
	"github.com/Woo-Yong0405/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:		Set the port of the server\n")
	fmt.Printf("-mode=rest:		Choose between html and rest\n")
	runtime.Goexit()
}

func Start() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	fmt.Println(*port, *mode)

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
