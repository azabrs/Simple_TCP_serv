package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"log"
)

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
	   log.Fatal(err)
	}
 }

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		fmt.Println(err)
	}
	for{
		go copyTo(os.Stdout , conn)
		copyTo(conn, os.Stdin)
	}

}