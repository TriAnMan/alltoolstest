package handler

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/TriAnMan/alltoolstest/domain"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type handler struct {
	tree domain.Collection
	log  logrus.FieldLogger
}

func New(tree domain.Collection, log logrus.FieldLogger) *handler {
	return &handler{
		tree: tree,
		log:  log,
	}
}

func (h *handler) generateUnique() string {
	b := make([]byte, 15)
	_, err := rand.Read(b)
	if err != nil {
		h.log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func (h *handler) constructLog(handler string, r *http.Request) *logrus.Entry {
	return h.log.
		WithField("subsystem", "http").
		WithField("handler", handler).
		WithField("remote-addr", r.RemoteAddr).
		WithField("method", r.Method).
		WithField("url", r.URL).
		WithField("req-id", h.generateUnique())
}

func sendError(w http.ResponseWriter, logEntry *logrus.Entry, httpStatus int, err error) {
	http.Error(w, "", httpStatus)

	logEntry.
		WithField("action", "response").
		WithField("http-status", httpStatus).
		WithError(err).
		Warning()
}

func sendOk(w http.ResponseWriter, logEntry *logrus.Entry) {
	http.Error(w, "", http.StatusOK)

	logEntry.
		WithField("action", "response").
		WithField("http-status", http.StatusOK).
		Info()
}

func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	logEntry := h.constructLog("search", r)
	logEntry.WithField("action", "request").Info()

	if r.Method != http.MethodGet {
		sendError(w, logEntry, http.StatusMethodNotAllowed, nil)
		return
	}

	valStr := r.FormValue("val")
	if valStr == "" {
		sendError(w, logEntry, http.StatusBadRequest, nil)
		return
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		sendError(w, logEntry, http.StatusBadRequest, err)
		return
	}

	if !h.tree.Has(val) {
		sendError(w, logEntry, http.StatusNotFound, nil)
		return
	}

	sendOk(w, logEntry)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	logEntry := h.constructLog("delete", r)
	logEntry.WithField("action", "request").Info()

	if r.Method != http.MethodDelete {
		sendError(w, logEntry, http.StatusMethodNotAllowed, nil)
		return
	}

	valStr := r.FormValue("val")
	if valStr == "" {
		sendError(w, logEntry, http.StatusBadRequest, nil)
		return
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		sendError(w, logEntry, http.StatusBadRequest, err)
		return
	}

	h.tree.Del(val)

	sendOk(w, logEntry)
}

func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	logEntry := h.constructLog("insert", r)
	logEntry.WithField("action", "request").Info()

	if r.Method != http.MethodPost {
		sendError(w, logEntry, http.StatusMethodNotAllowed, nil)
		return
	}

	valStr := r.FormValue("val")
	if valStr == "" {
		sendError(w, logEntry, http.StatusBadRequest, nil)
		return
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		sendError(w, logEntry, http.StatusBadRequest, err)
		return
	}

	h.tree.Put(val)

	sendOk(w, logEntry)
}

func (h *handler) Default(w http.ResponseWriter, r *http.Request) {
	logEntry := h.constructLog("default", r)
	logEntry.WithField("action", "request").Info()
	sendError(w, logEntry, http.StatusNotFound, nil)
}
