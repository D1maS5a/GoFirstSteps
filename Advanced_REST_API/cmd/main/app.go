package main

import (
	"advanREST/internal/config"
	"advanREST/internal/user"
	"advanREST/pkg/logging"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

//	func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//		name := params.ByName("name")
//		w.Write([]byte(fmt.Sprintf("Hello %s", name)))
//	}

func main() {
	logger := logging.GetLogger()
	// log.Println("create router")
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	// log.Println("register user handler")
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	// router.HandlerFunc("GET", "/", IndexHandler)
	// router.GET("/:name", IndexHandler)

	logger := logging.GetLogger()
	// log.Println("start application")
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		// /path/to/binary
		// Dir() -- /path/to
		logger.Info("detect app port")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// log.Println("server is listening port 0.0.0.0:1234")

	// log.Fatalln(server.Serve(listener))
	logger.Fatal(server.Serve(listener))
}
