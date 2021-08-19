package main

import (
	"bufio"
	"fmt"
	//"io"
	"net"
	"os"
	"time"
)


func readConsole(ch chan []byte) {
	// ввод данных с консоли
	for {
		fmt.Print(">")
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(line) > 250 {
			fmt.Println("Error message is very lagre")
			continue
		}
		_, b := EncodeProto(line[:len(line)-1])
		//fmt.Println("after encoded len:", len, b[:len])
		ch <- b
	}
}

func readSock(conn net.Conn ) {
	// прием данных из сокета
	buf := make([]byte,256)
	for {
		readed_len, err := conn.Read(buf)
		if err != nil {
			panic(err.Error())
		}
		if readed_len > 0  {
			_, content :=  DecodeProto(buf)
			fmt.Printf("->%s\n",  content)
			//fmt.Print(">")
		}
	}
}

func main(){
	ch := make(chan []byte)
	// Подключаемся к сокету
	defer close(ch)

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")

	buf := make([]byte,256)
	readed_len, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}
	if readed_len > 0 {
		fmt.Print("Connect from server: " , string(buf[:readed_len]))

	}
	go readConsole(ch)
	go readSock(conn)

	for {
		val, ok := <- ch
		if ok {

			len_writed, err := conn .Write(val)
			fmt.Println(val)
			if err != nil {
				fmt.Println("Write:", err.Error())
				break
			}
			fmt.Println("Send bytes:", len_writed)


		} else {
			time.Sleep(time.Second * 2)
		}

	}
	fmt.Println("Finished...")
	conn.Close()
}
