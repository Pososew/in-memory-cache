package main

import (
	"fmt"

	"github.com/Pososew/in-memory-cache/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, found := cache.Get("userId")
	if found {
		fmt.Println("UserdId: ", userId)
	} else {
		fmt.Println("User is not found")
	}

	cache.Delete("userId")
	userId, found = cache.Get("userId")
	if found {
		fmt.Println("UserdId: ", userId)
	} else {
		fmt.Println("User is not found")
	}
}

