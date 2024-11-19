package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"log"
	"errors"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

func getUsers(url, apiKey string) ([]User, error) {
    
    //creating new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

    //set apikey header
	req.Header.Set("X-API-Key", apiKey)

    //making GET request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	
	//decoding response body to JSON
	decoder := json.NewDecoder(res.Body)
	
	//decode JSON to User struct
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

    //creating new request object
	req, err := http.NewRequest("DELETE", fullURL, nil)
    if err != nil {
    	return err
    }
    
    //set apikey header
	req.Header.Set("X-API-Key", apiKey)
    
    //making DELETE request
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	
	if res.StatusCode > 299 {
		return errors.New("request to delete location unsuccessful")
	}
	return nil
}


func main() {
	userId := "0194fdc2-fa2f-4cc0-81d3-ff12045b73c8"
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	users, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logging user records:")
	logUsers(users)
	fmt.Println("---")

	err := deleteUser(url, userId, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted user with id: %s\n", userId)
	fmt.Println("---")

	newUsers, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(newUsers)
	fmt.Println("---")
}
