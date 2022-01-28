package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/supernova0730/project/uuid"
)

func main() {
	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	connMap := &sync.Map{}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		id := uuid.GenerateUUID(12)
		connMap.Store(id, conn)

		go handleConn(id, conn, connMap)
	}
}

func handleConn(id string, c net.Conn, connMap *sync.Map) {
	defer func() {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
		connMap.Delete(id)
	}()

	for {
		input, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		connMap.Range(func(key, value interface{}) bool {
			if keyID, ok := key.(string); ok {
				if keyID == id {
					return true
				}
			}
			if conn, ok := value.(net.Conn); ok {
				if _, err := conn.Write([]byte(input)); err != nil {
					log.Println(err)
				}
			}

			return true
		})
	}
}
