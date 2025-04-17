// go run main.go

package main

import "net/http"

func main() {
	println("\n  https://localhost:8443\n")
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != http.ErrServerClosed {
		panic(err)
	}
}
