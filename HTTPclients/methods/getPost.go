package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"bytes"
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

func createUser(url, apiKey string, data User) (User, error) {
    
    // encode our User struct as json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	
	// create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

    // set request headers
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

    // decode the json data from the response
	// into a new User struct
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func main() {
	userToCreate := User{
		Role:       "Junior Developer",
		Experience: 2,
		Remote:     true,
		User: struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Age      int    `json:"age"`
		}{
			Name:     "Dan",
			Location: "NOR",
			Age:      29,
		},
	}

	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	fmt.Println("Retrieving user data...")
	
	//getting user from server via GET request
	userDataFirst, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataFirst)
	fmt.Println("---")

	fmt.Println("Creating new character...")
	
	//POST new user on server
	creationResponse, err := createUser(url, apiKey, userToCreate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	
	
	jsonData, _ := json.Marshal(creationResponse)
	fmt.Printf("Creation response body: %s\n", string(jsonData))
	fmt.Println("---")

	fmt.Println("Retrieving user data...")
	
	//getting posted user from server
	userDataSecond, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataSecond)
	fmt.Println("---")	
}
