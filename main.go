package main

import (
	"fmt"
	"log"

	"github.com/ejaz0/blog_aggreator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	if err := cfg.SetUser("ejaz"); err != nil {
		log.Fatal(err)
	}
	cfg2, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg2.DBURL)
	fmt.Printf("%+v\n", cfg2)
}
