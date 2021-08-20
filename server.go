package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"time"

	//"bufio"
	"fmt"
	"log"
	"net"
	//"strings"
)


const (
	EXIT_COMMAND = "exit"
)

// Read message from a net.Conn
func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			// if the error is an End Of File this is still good
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}

// Write message to a net.Conn
// Return the number of bytes returned
func Write(conn net.Conn, encoded string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(encoded)
	if err == nil {
		err = writer.Flush()
	}
	return number, err
}

func handle(n int, cnn map[int] net.Conn ) {
	var message string
	var clientNo int
	conn := cnn[n]
	defer conn.Close()
	log.Printf("Now listnen: %s \n", conn.RemoteAddr().String())
	buf := make([]byte,256)
	for {

		fmt.Println("Read...")
		readed_len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("ReadErr:", err.Error())
			break
		}
		if readed_len > 0 {
			data_len,data := DecodeProto(buf)
			if data == EXIT_COMMAND {
				log.Println("Listener: Exit!")
				os.Exit(0)
			}

			clientNo = -1
			log.Printf("#%d: Response: %s len=%d\n",n, data, len(data))

			if data_len > 0 && data[0] == '#' {
				fmt.Sscanf(data,"#%d", &clientNo, &message)
				pos := strings.Index(data," ")
				message := data[pos:]
				fmt.Printf("Parser:ClientNo=%d message:%s", n, message )
				data = fmt.Sprintf("%d:%s", n, message )
				_, b := EncodeProto(data)
				fmt.Println(data)
				cnn[clientNo].Write(b)
				continue
			}

			data = fmt.Sprintf("#%d:%s", n, data )
			_, b := EncodeProto(data)
			for key,val := range(cnn)  {
				if n == key {
					continue
				}
				write_len, err := val.Write(b)
				if err != nil {
					fmt.Println(key, err)
					delete(cnn, key)
				} else {
					fmt.Printf( "#%d: writed %d\n", key, write_len)
				}
			}

			//log.Printf("Listener: Wrote %d byte(s) to %s \n", num, conn.RemoteAddr().String())
		}
	}
	fmt.Println("Close connection", n, "free connection:", len(cnn))
}

func main() {
	//var cnn map[int]  net.Conn    //
	cnn := make( map[int] net.Conn, 255)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Use fatal to exit if the listener fails to start
		log.Fatal(err)
	}
	defer listener.Close()

	i := 0
	for {
		if len(cnn) > 254 {
			log.Println("Connection pool is full")
			time.Sleep(time.Second )
		}
		conn, err := listener.Accept()
		if err != nil {
			// Print the error using a log.Fatal would exit the server
			log.Println(err)
			continue
		}
		out := fmt.Sprintf("client #%d\n", i)
		conn.Write([]byte(out))
		cnn[i] = conn
		go handle(i, cnn)
		// Using a go routine to handle the connection
		i++



	}
}

