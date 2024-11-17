package main

import "server-pulsa/delivery"

func main() {
	delivery.NewServer().Run()
}
