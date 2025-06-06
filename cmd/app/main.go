package main

import (
	"fmt"
	"mygogo/hello/internal/di"
)

func main() {
    container := di.NewContainer()
    
    fmt.Println("Server started on ", container.ApiServer.Addr)
}
