package main

import ( )

type UserItem struct {
	Email         string      `json:"email"`
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Preferences   interface{} `json:"preferences"`
	PusherChannel string      `json:"pusherChannel"`
	PusherKey     string      `json:"pusherKey"`
}
