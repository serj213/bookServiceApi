package main

import (
	"fmt"

	"github.com/serj213/bookServiceApi/internal/config"
)

func main(){
	cfg, err := config.Deal()
	if err != nil {
		panic(err)
	}

	fmt.Println("cfg ", cfg)

}