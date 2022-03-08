package main

import "net"

func handler(c net.Conn) {
	c.Write([]byte("Gangstas paradise mr Clifford"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":6060")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
