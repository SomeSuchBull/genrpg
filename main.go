/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "net/http/pprof"

	"github.com/genrpg/cmd"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "")
	// })
	// log.Fatal(http.ListenAndServe(":8000", nil))
	cmd.Execute()
}
