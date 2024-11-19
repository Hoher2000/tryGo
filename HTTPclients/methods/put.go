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

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	
	//ensoding user to json
    jsondata, err := json.Marshal(data)
    if err != nil{
        return User{}, err
    }
	
	// create a new request
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsondata))
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

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id
	
    //creating new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

    //set apikey header
	req.Header.Set("X-API-Key", apiKey)

    //making GET request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	
	//decoding response body to JSON
	decoder := json.NewDecoder(res.Body)
	
	//decode JSON to User struct
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func logUser(user User) {
	fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
		user.User.Name, user.Role, user.Experience, user.Remote)
}

func main() {
	userId := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	userData, err := getUserById(baseURL, userId, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	logUser(userData)

	fmt.Printf("Updating user with id: %s\n", userData.ID)
	userData.Role = "Senior Backend Developer"
	userData.Experience = 7
	userData.Remote = true
	userData.User.Name = "Allan"

	updatedUser, err := updateUser(baseURL, userId, apiKey, userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(updatedUser)
}
