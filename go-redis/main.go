package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	redClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	mainContext := context.Background()

	// check redis connection
	ping, err := redClient.Ping(mainContext).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("server send ping and recieve ", ping)

	// set cache value
	type Person struct {
		ID   string `json:"id"`
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	person1Id := uuid.NewString()

	person1, err := json.Marshal(&Person{
		ID:   person1Id,
		Name: "Elon",
		Age:  40,
	})
	if err = redClient.Set(mainContext, person1Id, person1, 0).Err(); err != nil {
		fmt.Println("could not set data to cache", err.Error())
	}

	// get cache value from redis
	val, err := redClient.Get(mainContext, person1Id).Result()
	if err != nil {
		fmt.Println("could get cache data for ", person1Id, err.Error())
	}

	fmt.Printf("data for %s is: %+v \n", person1Id, val)

	// invalidate redis cache
	delCache, err := redClient.Del(mainContext, person1Id).Result()
	if err != nil {
		fmt.Println("could not delete cache for", person1Id, err.Error())
	}

	fmt.Println("deleted cache:", delCache)
}
