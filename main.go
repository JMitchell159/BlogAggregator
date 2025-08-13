package main

import (
	"fmt"
	
	"github.com/JMitchell159/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	err = cfg.SetUser("JMitchell159")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", cfg.DB_URL, *cfg.Current_User_Name)
}