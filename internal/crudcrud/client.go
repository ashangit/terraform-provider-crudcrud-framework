package crudcrud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const resource string = "unicorn"

type Unicorn struct {
	Id, Name, Color string
	Age             int
}

type CrudcrudClient struct {
	Endpoint string
}

func (c CrudcrudClient) Get(id string) (Unicorn, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", c.Endpoint, resource, id))
	if err != nil {
		return Unicorn{}, fmt.Errorf("Failed to get unicorn %s: %w", id, err)
	}
	fmt.Printf("Response: %v\n", resp)

	// TODO manage error
	data, err := io.ReadAll(resp.Body)
	fmt.Printf("Results: %v\n", data)
	if err != nil {
		return Unicorn{}, fmt.Errorf("Failed to read body: %w", err)
	}

	var unicorn Unicorn
	json.Unmarshal(data, &unicorn)
	unicorn.Id = id

	return unicorn, nil
}

func (c CrudcrudClient) Create(unicorn *Unicorn) error {
	json_data, err := json.Marshal(unicorn)
	if err != nil {
		return fmt.Errorf("Failed to json marshal unicorn %s: %w", unicorn, err)
	}

	// TODO manage error
	resp, err := http.Post(fmt.Sprintf("%s/%s", c.Endpoint, resource), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return fmt.Errorf("Failed to create unicorn %s: %w", unicorn, err)
	}
	fmt.Printf("Create Response: %v\n", resp)

	location := resp.Header.Get("Location")
	locations := strings.Split(location, "/")
	unicorn.Id = locations[len(locations)-1]

	return nil
}

func (c CrudcrudClient) Update(unicorn Unicorn) error {
	json_data, err := json.Marshal(unicorn)
	if err != nil {
		return fmt.Errorf("Failed to json marshal unicorn %s: %w", unicorn, err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", c.Endpoint, resource, unicorn.Id), bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// TODO manage error
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("Update Response: %v\n", resp)

	return nil
}

func (c CrudcrudClient) Delete(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.Endpoint, resource, id), nil)
	if err != nil {
		return err
	}

	// TODO manage error

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("Delete Response: %v\n", resp)

	return nil
}
