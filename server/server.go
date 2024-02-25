package main

import (
	"fmt"
	"net"
)

func Handle_connect(con net.Conn){
	buf := make([]byte, 1024)
	for{
		con.Write([]byte("Hello, what's your name?\n"))
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		 }
		con.Write(append([]byte("Goodbye, "), buf[:n]...))
	}

}

func main(){
	Listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		con, err := Listener.Accept()
		if err != nil{
			fmt.Println(err)
		}
		go Handle_connect(con)

	}

}