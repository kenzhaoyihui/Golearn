package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"bufio"
	"io"
	"strings"
	"github.com/mgutz/str"
	"net/url"
	"crypto/md5"
	"encoding/hex"
)

const (
	HANDLER_DIG = "/dig?"
)

type (
	cmdParams struct {
		logFilePath string
		routineNum int
	}

	digData struct {
		time string
		url string
		refer string
		ua string
	}

	urlData struct {
		//use to transfer
		data digData
		uid string
	}

	urlNode struct {
		//use to storage
	}

	storageBlock struct {
		counterType string
		storageModel string
		unode urlNode
	}
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main()  {
	//get params
	logFilePath := flag.String("logFilePath", "/var/log/message", "log path")
	routineNum := flag.Int("routineNum", 5, "consumer number by goroutine")

	l := flag.String("l", "/tmp/log", "this program runtime log target file path")
	flag.Parse()

	params := cmdParams{*logFilePath, *routineNum}

	//print logs
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}

	log.Infoln("Exec start.")
	log.Infof("Params: logFilePath=%s, routineNum=%d", *logFilePath, *routineNum)



	//init channels, used to data transfer
	var logChannel = make(chan string, 3 * params.routineNum)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)


	//log consumer
	go readFileLinebyLine(params, logChannel)

	//handle logs

	for i:=0;i<params.routineNum;i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}


	//pv uv counter
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)


	//create the storage
	go dataStorage(storageChannel)

	time.Sleep(1000*time.Second)
}


func readFileLinebyLine(params cmdParams, logChannel chan string) error {
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("ReadFileLinebyLine can't open file: %s", params.logFilePath)
		return err
	}
	defer fd.Close()

	bufferRead := bufio.NewReader(fd)
	count := 0
	for {
		line, err := bufferRead.ReadString('\n')
		logChannel <- line
		count++

		if count%(1000*params.routineNum) == 0 {
			log.Infof("ReadFileLinebyLine line: %d", count)
		}
		if err != nil {
			if err == io.EOF {
				time.Sleep(3*time.Second)
				log.Infof("ReadFileLinebyLine wait: %d", count)
			}else{
				log.Warningf("ReadFileLinebyLine log error: %s", err)
			}
		}
	}
	return nil
}


func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) error{

	for logStr := range logChannel {
		//log string, select data
		data := cutLogFetchData(logStr)

		//uid, --> md5(refer+ua)
		hasher := md5.New()
		hasher.Write([]byte(data.refer+data.ua))

		uid := hex.EncodeToString(hasher.Sum(nil))

		uData := urlData{data, uid}

		pvChannel <- uData
		uvChannel <- uData
	}

	return nil
}

func cutLogFetchData(logStr string) digData {
	logStr = strings.TrimSpace(logStr)
	pos1 := str.IndexOf(logStr, HANDLER_DIG, 0)
	if pos1 == -1 {
		return digData{}
	}

	pos1 += len(HANDLER_DIG)
	pos2 := str.IndexOf(logStr, " HTTP/", pos1)

	d := str.Substr(logStr, pos1, pos2-pos1)

	urlInfo, err := url.Parse("http://127.0.0.1/?"+d)
	if err != nil {
		return digData{}
	}

	data := urlInfo.Query()
	return digData{
		data.Get("time"),
		data.Get("refer"),
		data.Get("url"),
		data.Get("ua"),
	}
}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func dataStorage(storageChannel chan storageBlock) {

}
