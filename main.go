package main 

import (
	"fmt"
	"os"
	"code.google.com/p/gcfg" // https://code.google.com/p/gcfg/
)

type Config struct {
        Ninja struct {
                APIAccessToken string
        	}
        }

func main() {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
  		fmt.Println("error:", err)
  		pwd, err2 := os.Getwd()
    	if err2 != nil {
	        fmt.Println(err)
        	os.Exit(1)
    	}
    	fmt.Println("You should place the file config.gcfg in this directory:",pwd)
  		os.Exit(1)
	}
	fmt.Println("Values from config.gcfg: APIAccessToken =", cfg.Ninja.APIAccessToken)
	
	var NinjaCurrentUser UserItem 
   	ncuerr := NinjaGetUser(cfg.Ninja.APIAccessToken, &NinjaCurrentUser)
   	if ncuerr != nil {
        fmt.Println("Error! ", ncuerr)
        os.Exit(1)
    } 
   
   	fmt.Println("Your UserID is:",NinjaCurrentUser.Name," email:",NinjaCurrentUser.Email)
	
}
