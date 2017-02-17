package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	index, err := ioutil.ReadFile("html/index.html")
	if err != nil {
		log.Println(err)
	}

	type Person struct {
		Name string
		Age  int
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Client subscribed")

		myPerson := Person{
			Name: "Bill",
			Age:  0,
		}

		for {
			time.Sleep(2 * time.Second)
			if myPerson.Age < 40 {
				myJSON, err := json.Marshal(myPerson)
				if err != nil {
					log.Println(err)
					return
				}

				err = conn.WriteMessage(websocket.TextMessage, myJSON)
				if err != nil {
					log.Println(err)
					break
				}

				myPerson.Age += 2
			} else {
				conn.Close()
				break
			}
		}

		log.Println("Client unsubscribed")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})

	http.ListenAndServe(":8000", nil)
}
