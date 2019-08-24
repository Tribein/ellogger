// ellogger project main.go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type LogRecord struct {
	AppName string `json:"appname"`
	MyID    string `json:"hostname"`
	LineId  uint32 `json:"lineid"`
	Data    string `json:"message"`
}

/*
func setConnection(hostport string) (bool, net.Conn) {
	fmt.Println("connecting to " + hostport)
	logconn, err := net.Dial("tcp", hostport)
	_ = logconn
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		res, logconn := setConnection(hostport)
		_ = res
		return true, logconn
	}
	return true, logconn
}

func sendData(conn net.Conn, input *LogRecord) {

	//logRecJSON, _ := json.Marshal(input)
	fmt.Println("sending test")
	_, err := conn.Write([]byte("gotest2"))
	if err != nil {
		log.Println(err)
		if conn != nil {
			conn.Close()
		}
	}
}
*/
func main() {
	var lineid uint32 = 0
	logRec := &LogRecord{
		AppName: os.Args[2],
		MyID:    os.Args[1],
		Data:    ""}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if v := strings.TrimSpace(scanner.Text()); len(v) > 0 {
			lineid++
			logRec.LineId = lineid
			logRec.Data = v
			logRecJSON, _ := json.Marshal(logRec)
			fmt.Println(string(logRecJSON))
			if lineid > 4294967293 {
				lineid = 0
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
