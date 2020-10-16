package main

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	Trace *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func initLog() {
	// Log file.
	file, err := os.OpenFile(path.Join(workPath, NAME+".log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)

	Trace = log.New(mw, "["+NAME+"] [trace] ", log.Lshortfile|log.Ldate|log.Ltime)
	Debug = log.New(mw, "["+NAME+"] [debug] ", log.Ldate|log.Ltime)
	Info = log.New(mw, "["+NAME+"] [info ] ", log.Ldate|log.Ltime)
	Warn = log.New(mw, "["+NAME+"] [warn ] ", log.Ldate|log.Ltime)
	Error = log.New(mw, "["+NAME+"] [error] ", log.Lshortfile|log.Ldate|log.Ltime)

	// Trace.Printf("message")
	// Debug.Printf("message")
	// Info.Printf("message")
	// Warn.Printf("message")
	// Error.Printf("message")
}

// todo -remove
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

// todo - end
