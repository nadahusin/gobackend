package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type goods struct{}

// func reqs() {
// 	mainRoute := mux.NewRouter()
// 	// mainRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Hello Worlds"))
// 	// })

// 	mainRoute.HandleFunc("/hello", sampleHandler).Methods("GET")
// 	mainRoute.HandleFunc("/poduct/{username}", paramsHendler).Methods("GET")
// 	mainRoute.HandleFunc("/users", queryHendler).Methods("GET")
// 	mainRoute.HandleFunc("/food", bodyHendler).Methods("POST")

// 	fmt.Println("app runn on port 8082")
// 	err := http.ListenAndServe(":8082", mainRoute)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// // func sampleHandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Write([]byte("Hello Worlds"))
// // }

// // // fatch params
// // func paramsHendler(w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)

// // 	fmt.Fprintf(w, "Params : %v ", vars["username"])
// // }

// // // fatch query urkl
// // func queryHendler(w http.ResponseWriter, r *http.Request) {
// // 	vars := r.URL.Query()

// // 	fmt.Fprintf(w, "Params : %v ", vars)
// // }

// // // fatch query urkl
// // func bodyHendler(w http.ResponseWriter, r *http.Request) {
// // 	var good goods

// // 	json.NewDecoder(r.Body).Decode(&good)

// // 	fmt.Println(good)
// // }
