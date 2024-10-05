package main

// #include "Windows.h"
import (
	//"C"
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ConsoleMode()
	loggerUser = CreateLogUser()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		_ = fmt.Errorf("unable to start server")
	}
	go UpdateScreen()
	for {
		conn, err := ln.Accept()
		if err != nil {
			_ = fmt.Errorf("unable to accept connection")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	HandleLog("New Conn Start", Green)
	var userp *User
	SetTotalConn(1)
	input := bufio.NewScanner(conn)
	if input.Scan() {
		if input.Text() == "start" {
			HandleLog("Conn Verified", Green)
			if input.Scan() {
				userp = RegisterUser(input.Text(), conn)
				fmt.Fprintf(conn, "welcome")
				HandleLog("New User : "+userp.name, Green)
			}
		} else {
			HandleLog("Conn Not Verified", Red)
			conn.Close()
			SetTotalConn(-1)
			Terminate(userp)
			fmt.Print(input.Text())
			return
		}
	}
	for input.Scan() {
		HandleMessage(userp, input.Text(), &conn)
	}
	HandleLog("Conn Lost", Red)
	SetTotalConn(-1)

	Terminate(userp)

}

func RegisterUser(name string, conn net.Conn) *User {
	new_user := UserExist(name)
	mu_cl.Lock()
	if new_user == nil {
		new_user = new(User)
		new_user.name = name
		new_user.logqueue = LogQueue{count: queuelimit}
		os.Mkdir("files\\"+name, os.ModePerm)
		cl = append(cl, new_user)
	}
	new_user.remoteaddress = conn.RemoteAddr().String()
	new_user.state = true
	h, m, _ := time.Now().Clock()
	ye, mo, da := time.Now().Date()
	new_user.starttime = fmt.Sprintf("%2d:%2d", h, m)
	new_user.filepath = fmt.Sprintf("files\\%s\\%.4d-%.2d-%.2d_%.2d.log", name, ye, int(mo), da, h)
	mu_cl.Unlock()
	return new_user
}

func Terminate(userp *User) {
	if userp != nil {
		mu_cl.Lock()
		userp.state = false
		h, m, _ := time.Now().Clock()
		userp.starttime = fmt.Sprintf("%2d:%2d", h, m)
		mu_cl.Unlock()
	}
}
func HandleMessage(userp *User, m string, con *net.Conn) {
	//TODO : check message type
	if true {
		new_m := LogMessage{date: 1, message: m}
		userp.logqueue.Add(&new_m)

		SaveToFile(&userp.filepath, &m)
	}
	fmt.Fprintf(*con, "Received")
}

func UserExist(name string) *User {
	mu_cl.Lock()
	defer mu_cl.Unlock()
	for _, u := range cl {
		if u.name == name {
			fmt.Print("exist")
			return u
		}
	}
	return nil
}
