package main

import (
	"flag"
	"github.com/TriAnMan/alltoolstest/infrastructure/file"
	"github.com/TriAnMan/alltoolstest/infrastructure/httpservice"
	"github.com/TriAnMan/alltoolstest/interface/handler"
	"github.com/TriAnMan/alltoolstest/usecase/wrapper/bst"
	log "github.com/sirupsen/logrus"
	syslog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	port := flag.String("port", "19100", "http service port")
	initFile := flag.String("init-file", "./init.json", "data to initialize BST")
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, syscall.SIGTERM) // doesn't work on windows

	log.WithField("subsystem", "system").Info("loading data")
	initData := file.Load(*initFile)

	tree := bst.New(log.StandardLogger(), initData)

	log.WithField("subsystem", "system").Info("starting server")
	go func() {
		w := log.StandardLogger().Writer()
		defer func() { _ = w.Close() }()

		h := handler.New(tree, log.StandardLogger())
		httpservice.HandlePanicFunc(log.StandardLogger(), "/search", h.Search)
		httpservice.HandlePanicFunc(log.StandardLogger(), "/insert", h.Insert)
		httpservice.HandlePanicFunc(log.StandardLogger(), "/delete", h.Delete)
		httpservice.HandlePanicFunc(log.StandardLogger(), "/", h.Default)
		server := &http.Server{Addr: ":" + *port, ErrorLog: syslog.New(w, "", 0)}
		log.Panic(server.ListenAndServe())
	}()

	<-interrupt
	log.WithField("subsystem", "system").Info("termination")
}
