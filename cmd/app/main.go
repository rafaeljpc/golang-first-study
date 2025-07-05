package main

import (
	"fmt"

	"github.com/rafaeljpc/golang-first-study/internal/di"
)

func main() {
    container := di.NewContainer()
    
    fmt.Println("Server started on ", container.ApiServer.Addr)
}
