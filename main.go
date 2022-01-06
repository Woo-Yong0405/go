package main

import (
	"github.com/Woo-Yong0405/nomadcoin/cli"
	"github.com/Woo-Yong0405/nomadcoin/db"
)

func main() {
	defer db.DB().Close()
	cli.Start()
}
