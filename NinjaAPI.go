package main 

import (
	"fmt"
	"github.com/franela/goreq"
)

	// returns the user information
	func NinjaGetUser(accessToken string, NU *UserItem) (error) {
		res, err := goreq.Request{
		    Method:     "GET",
	    	Uri:        "https://api.ninja.is/rest/v0/user?user_access_token="+accessToken,
	    	Accept: 	"application/json",
		}.Do()
		if err != nil {
			fmt.Println("func NinjaGetUser, Error: ",  err)
		} else {
			// fmt.Println(res, err)
			// fmt.Println(res.Response, err)
			res.Body.FromJsonTo(&NU)
		}
		
		return err
	}
