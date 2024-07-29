package main

import (
	"advanREST/internal/user"
	"advanREST/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

//	func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//		name := params.ByName("name")
//		w.Write([]byte(fmt.Sprintf("Hello %s", name)))
//	}
//
// 14:30
func main() {
	logger := logging.GetLogger()
	// log.Println("create router")
	logger.Info("create router")
	router := httprouter.New()

	// log.Println("register user handler")
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	// router.HandlerFunc("GET", "/", IndexHandler)
	// router.GET("/:name", IndexHandler)

	logger := logging.GetLogger()
	// log.Println("start application")
	logger.Info("start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// log.Println("server is listening port 0.0.0.0:1234")
	logger.Info("server is listening port 0.0.0.0:1234")
	// log.Fatalln(server.Serve(listener))
	logger.Fatal(server.Serve(listener))
}
