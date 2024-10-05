package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// var cl map[string]*user = make(map[string]*user)
var key string = "start"
var name string = fmt.Sprintf("client-%d", rand.Int31n(100))
var SendPackage chan string = make(chan string, 100)

var ip string = "127.0.0.1:8080"

func main() {
	fmt.Println(name)
	var conn net.Conn
	var err error
	var buffer []byte = make([]byte, 1024)
	go SendData()
	for {
		conn, err = net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("Error : Will try 10 sec later")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("Connected")
			if !Login(&conn, &buffer) {
				fmt.Print("Login failed")
				break
			}
			fmt.Println("Logged in")
			for msg := range SendPackage {
				fmt.Println("sss")
				i, e := fmt.Fprintf(conn, "%s\n", msg)
				if e == nil {
					fmt.Printf("%d bytes sent\n", i)
				} else {
					break
				}
				//read answer
			}
		}
		if conn != nil {
			conn.Close()
		}
	}
}

func Login(conn *net.Conn, buffer *[]byte) bool {
	fmt.Fprintf(*conn, "%s\n", key)
	fmt.Fprintf(*conn, "%s\n", name)
	(*conn).Read(*buffer)
	if string((*buffer)[:7]) == "welcome" {
		fmt.Printf("Success : %s", (*buffer)[:7])
		return true
	} else {
		fmt.Printf("Fail : %s", (*buffer)[:7])
		return false
	}
}

func SendData() {
	for {
		time.Sleep(2 * time.Second)
		h, m, s := time.Now().Clock()
		//store
		SendPackage <- fmt.Sprintf("%d:%d:%d", h, m, s)
	}
}
