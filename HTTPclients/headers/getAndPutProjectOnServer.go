package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)


type Project struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Assignees int    `json:"assignees"`
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

func getProjectResponse(apiKey, url string) (Project, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Project{}, err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return Project{}, err
	}

	return project, nil
}

func putProject(apiKey, url string, project Project) error {
	client := &http.Client{}
	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/projects/52fdfc07-2182-454f-963f-5f0f9a621d72"
	apiKey := generateKey()

	oldProject, err := getProjectResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting old project:", err)
		return
	}
	fmt.Println("Got old project:")
	fmt.Printf("- title: %s\n", oldProject.Title)
	fmt.Printf("- assignees: %d\n", oldProject.Assignees)
	fmt.Println("--------------------------------")

	newProjectData := Project{
		ID:        "52fdfc07-2182-454f-963f-5f0f9a621d72",
		Title:     "Product Roadmap 2025",
		Completed: false,
		Assignees: 1,
	}

	if err := putProject(apiKey, url, newProjectData); err != nil {
		fmt.Println("Error updating project:", err)
		return
	}
	fmt.Println("Project updated!")
	fmt.Println("---")

	//newApiKey := generateKey()
	newProject, err := getProjectResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting new project:", err)
		return
	}
	fmt.Println("Got new project:")
	fmt.Printf("- title: %s\n", newProject.Title)
	fmt.Printf("- assignees: %d\n", newProject.Assignees)
	fmt.Println("--------------------------------")
}
