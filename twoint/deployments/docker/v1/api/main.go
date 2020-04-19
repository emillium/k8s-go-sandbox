package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	v1 "github.com/emillium/k8s-go-sandbox/twoint/pkg/api/v1"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  29-Jan-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func main() {
	//	Connect to Add service
	conn, err := grpc.Dial("twoint-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	twoIntClient := v1.NewTwoIntServiceClient(conn)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/add/{a}/{b}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UFT-8")

		vars := mux.Vars(r)
		a, err := strconv.ParseUint(vars["a"], 10, 64)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid parameter A")
		}
		b, err := strconv.ParseUint(vars["b"], 10, 64)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid parameter B")
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &v1.AddRequest{Api: "v1", TwoInt: &v1.TwoInt{A: a, B: b}}

		if resp, err := twoIntClient.Add(ctx, req); err == nil {
			msg := fmt.Sprintf("Summation is %d", resp.Result)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("GET")

	fmt.Println("Application is running on : 8080 .....")
	http.ListenAndServe(":8080", routes)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}
