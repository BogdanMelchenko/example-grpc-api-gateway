package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/BogdanMelchenko/example/proto/logic"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, os.Getenv("LOGIC_ADDRESS"), opts)
	//err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, "logic:9090", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(os.Getenv("ENDPOINT_HOST"), mux)
	//return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
