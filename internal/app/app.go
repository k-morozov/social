package app

import (
	"fmt"

	"github.com/k-morozov/social/internal/service/http_service"
)

func Run() {
	fmt.Println("Hello world!")

	server, err := http_service.NewServiceHTTP()
}
