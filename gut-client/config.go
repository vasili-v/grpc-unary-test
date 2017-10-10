package main

import "flag"

var (
	server string
	total  int
	limit  int
	size   int
)

func init() {
	flag.StringVar(&server, "s", ":5555", "address:port of server")
	flag.IntVar(&total, "n", 5, "number of requests to send")
	flag.IntVar(&limit, "l", 100, "limit for messages to send simultaneously")
	flag.IntVar(&size, "size", 60, "payload size")

	flag.Parse()
}
