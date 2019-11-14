package httpservice

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func HandlePanicFunc(log logrus.FieldLogger, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	defer func() {
		err := recover()
		if err != nil {
			buf := make([]byte, 1<<20)
			n := runtime.Stack(buf, true)
			log.Fatalf("panic: %v\n\n%s", err, buf[:n])
		}
	}()
	http.HandleFunc(pattern, handler)
}
