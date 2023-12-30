package main

import (
    "fmt"
    "log"
    "net/http"
    "crypto/sha256"
    "crypto/subtle"
	)

func handler(w http.ResponseWriter, r *http.Request) {
    username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte("admin"))
			expectedPasswordHash := sha256.Sum256([]byte("password"))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
			
			fmt.Println(username, password)
			if usernameMatch && passwordMatch {
				fmt.Fprintf(w, "Hello, Admin")
				return
			}
		}

    fmt.Fprintf(w, "Hello, guest")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
