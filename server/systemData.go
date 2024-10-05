package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var mu_lo sync.Mutex
var loggerUser *User

var mu_cl sync.Mutex
var cl []*User

var mu_tc sync.Mutex
var total_conn int = 0

func SetTotalConn(i int) {
	mu_tc.Lock()
	total_conn += i
	mu_tc.Unlock()
}
func GetTotalConn() int {
	mu_tc.Lock()
	defer mu_tc.Unlock()
	return total_conn
}

func SaveToFile(filepath *string, m *string) {
	file, err := os.OpenFile(*filepath, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Could not open %s", *filepath)
		return
	}
	defer file.Close()
	_, err2 := file.WriteString(*m)
	file.WriteString("\n")
	if err2 != nil {
		fmt.Println("Could not write text to example.txt")
	}
}

func HandleLog(m string, color int) {
	mu_lo.Lock()
	defer mu_lo.Unlock()
	m = fmt.Sprintf("%25s", m)
	if color == Green {
		m = "\x1B[32m" + m + "\x1B[37m"
	} else if color == Red {
		m = "\x1B[31m" + m + "\x1B[37m"
	}
	new_m := LogMessage{date: 1, message: m}
	loggerUser.logqueue.Add(&new_m)
}

func CreateLogUser() *User {
	new_user := new(User)
	new_user.name = "Logger"
	new_user.logqueue = LogQueue{count: queuelimit}
	os.Mkdir("files\\"+"Logger", os.ModePerm)
	cl = append(cl, new_user)
	new_user.state = true
	h, m, _ := time.Now().Clock()
	ye, mo, da := time.Now().Date()
	new_user.starttime = fmt.Sprintf("%2d:%2d", h, m)
	new_user.filepath = fmt.Sprintf("files\\%s\\%.4d-%.2d-%.2d_%.2d.log", "Logger", ye, int(mo), da, h)
	return new_user
}
