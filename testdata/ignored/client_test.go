package client

import (
	"fmt"
	"log"
)

// Create a client with baseURL http://localhost:6060
// and call GET http://localhost:6060/users/me
func ExampleClient_Get() {
	cl := New("http://localhost:6060")

	resp, err := cl.Get("users", "me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}
