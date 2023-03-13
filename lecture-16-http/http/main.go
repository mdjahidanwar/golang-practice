package main

import (
	"fmt"
	"net"
	"os"
)





func main() {
//net->listen- generate listener and error--> err print if any 
	nl, err := net.Listen("tcp", ":8888") //1 to 65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1=stop with error
	}
// listener gives Accept-- 2 datas -- nets.conn and error --> print error 
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		//continue
	}

	bs := make([]byte, 1024) //text


// read the request by Read method, parameter will be bytes, it is a package to handle texts, 1500 byte is maximum and there are some unneccessary datas, good practice is 1024
// returns 2 value int and error, print to get exact how many bytes it gets 
	n, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
		//continue
	}
	fmt.Println(n)
	// convert byte string to regular string 
	regstr := string(bs) //conversion-- different datatype exchange 
	fmt.Println(regstr)
// a body for web page 
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>hello world </title>
	</head>
	<body>
          helloeeeeeeee		
	</body>
	</html>`

	resp := "HTTP/1.1 200 OK\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"
	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))
	//conn.Close()
}
