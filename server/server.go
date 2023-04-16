package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rajatxs/go-iamone/rpc"

	"github.com/gorilla/mux"
	gorilla_rpc "github.com/gorilla/rpc"
	gorilla_rpc_json "github.com/gorilla/rpc/json"
)

var instance *http.Server

// Returns server address
func GetAddr() string {
	return fmt.Sprintf(
		"%s:%s",
		os.Getenv("IAMONE_SERVER_HOST"),
		os.Getenv("IAMONE_SERVER_PORT"))
}

// Set of operation needs to be execute before running server
func presetup() error {
	var router = mux.NewRouter()
	var rpcs = gorilla_rpc.NewServer()

	rpcs.RegisterCodec(gorilla_rpc_json.NewCodec(), "application/json")
	rpcs.RegisterService(new(rpc.User), "user")
	rpcs.RegisterService(new(rpc.Common), "rpc")
	router.Handle("/x", rpcs)

	instance = &http.Server{
		Addr:    GetAddr(),
		Handler: router,
	}
	return nil
}

// Starts server instance
func Start() (err error) {
	if err = presetup(); err != nil {
		log.Fatal("failed to setup server", err)
		return err
	}

	log.Printf("server starting at %s", GetAddr())

	if err = instance.ListenAndServe(); err == http.ErrServerClosed {
		return nil
	} else {
		return err
	}
}

// Shutdowns server instance
func Stop() (err error) {
	err = instance.Close()

	if err == http.ErrServerClosed {
		return nil
	} else {
		return err
	}
}
