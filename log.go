package main

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func initLog() {
	// Log file.
	logFile, err := os.OpenFile(path.Join(workPath, NAME+".log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// Log configuration.
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetPrefix("[" + NAME + "] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lmsgprefix)
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Ldate | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func checkError(err error) bool {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		function, file, line, _ := runtime.Caller(1)
		log.Printf("[error] [%s] [%s:%d] %v", filepath.Base(file), runtime.FuncForPC(function).Name(), line, err)
		return true
	}
	return false
}

func Error(err error) {
	function, file, line, _ := runtime.Caller(1)
	log.Printf("[error] [%s] [%s:%d] %v", filepath.Base(file), runtime.FuncForPC(function).Name(), line, err)
}

func debug(format string, a ...interface{}) {
	log.Printf("[debug] "+format, a...)
}

func warn(format string, a ...interface{}) {
	log.Printf("[warn] "+format, a...)
}

func trace(format string, a ...interface{}) {
	log.Printf("[trace] "+format, a...)
}
