package grapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	gril "github.com/zeuce/gril"
)

var fullLogFile string
var logP string

// SetupLogging will log add logHandler() to router and create log file with supplied args
// @param logDir - The directory to save your log file to
// @param logFile - The log file name
// @param logPrefix - The prefix displayed in console / log file
func SetupLogging(logDir string, logFile string, logPrefix string) {
	Router.Use(logHandler)
	fullLogFile = logDir + "/" + logFile
	logP = logPrefix
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)

	}
	if _, err := os.Stat(fullLogFile); os.IsNotExist(err) {
		ioutil.WriteFile(fullLogFile, []byte(""), 0744)
	}
}

// logHandler will log everytime a request is made to server
func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(fullLogFile, os.O_APPEND|os.O_RDWR, 0744)
		defer f.Close()
		gril.Check(err)
		logger := log.New(f, logP+" ", log.LstdFlags)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		fmt.Printf("%s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
