package utils

import (
	"log"
	"os"
)

var (
	// Log ...
	Log *log.Logger
	// PricesUtils stores the stdout/stderr output log
	PricesUtils = "./prices.log"
)

func init() {
	// set location of log file
	var logpath = PricesUtils

	// flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	// Log.Println("LogFile : " + logpath)
}
