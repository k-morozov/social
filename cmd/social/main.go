package main

import (
	"fmt"

	"github.com/k-morozov/social/internal/service"
)

func main() {
	srv, err := service.NewServiceHTTP(
		service.OptionSocial(),
		service.OptionLogger(),
	)
	if err != nil {
		fmt.Println("Failed: {}", err)
		return
	}

	if err = srv.ListenAndServe(); err != nil {
		fmt.Println("Failed: {}", err)
		return
	}
}
