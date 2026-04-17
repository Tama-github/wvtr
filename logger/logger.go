package logger

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	DumpLog *log.Logger
	ErrLog  *log.Logger
)

func init() {
	// set location of log file
	path, err := os.Getwd()
	if err != nil {
		return
	}
	date := time.Now().Unix()
	var dumplogpath = path + fmt.Sprintf("/tmp/logs/dump.%d.log", date)
	var errlogpath = path + fmt.Sprintf("/tmp/logs/error.%d.log", date)

	flag.Parse()
	var dumpFile, err1 = os.Create(dumplogpath)
	var errFile, err2 = os.Create(errlogpath)
	// defer dumpFile.Close()
	// defer errFile.Close()

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	DumpLog = log.New(io.MultiWriter(dumpFile, os.Stdout), "", log.Ldate|log.LstdFlags|log.Lshortfile)

	ErrLog = log.New(io.MultiWriter(errFile, os.Stdout), "", log.Ldate|log.LstdFlags|log.Lshortfile)
}
