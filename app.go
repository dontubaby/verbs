package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

const (
	adress = "127.0.0.1"
	port   = ":1408"
)

var verbs = map[int]string{
	1:  "Don't communicate by sharing memory, share memory by communicating.",
	2:  "Concurrency is not parallelism.",
	3:  "Channels orchestrate; mutexes serialize.",
	4:  "The bigger the interface, the weaker the abstraction.",
	5:  "Make the zero value useful.",
	6:  "interface{} says nothing.",
	7:  "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	8:  "A little copying is better than a little dependency.",
	9:  "Syscall must always be guarded with build tags.",
	10: "Cgo must always be guarded with build tags.",
	11: "Cgo is not Go.",
	12: "With the unsafe package there are no guarantees.",
	13: "Clear is better than clever.",
	14: "Reflection is never clear.",
	15: "Errors are values.",
	16: "Don't just check errors, handle them gracefully.",
	17: "Design the architecture, name the components, document the details.",
	18: "Documentation is for users.",
	19: "Don't panic.",
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()
	var indexVerb int
	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	message := strings.TrimSuffix(string(b), "\n")
	message = strings.TrimSuffix(message, "\r")
	switch {
	case message == "random":
		indexVerb = rand.Intn(19)
		if indexVerb == 0 {
			indexVerb += 1
		}
		conn.Write([]byte(verbs[indexVerb]))
	case message == "all":
		for _, v := range verbs {
			conn.Write([]byte(v))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", adress+port)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	fmt.Printf("Start listen on %v%v\n", adress, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go connectionHandler(conn)
	}
}
