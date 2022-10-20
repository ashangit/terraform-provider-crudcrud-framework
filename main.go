package main

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/crudcrud"
)

func main() {
	client := crudcrud.CrudcrudClient{Endpoint: "https://crudcrud.com/api/c6d5829262e24691988c3298126df5da"}
	unicorn := crudcrud.Unicorn{Name: "nico", Age: 5, Colour: "blue"}

	if err := client.Create(&unicorn); err != nil {
		panic(err)
	}
	fmt.Printf("Create: %v\n", unicorn)

	unicorn2, err := client.Get(unicorn.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get: %v\n", unicorn2)

	unicorn.Name = "david"
	if err := client.Update(unicorn); err != nil {
		panic(err)
	}
	fmt.Printf("Update: %v\n", unicorn)

	unicorn3, err := client.Get(unicorn.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get: %v\n", unicorn3)

	if err := client.Delete(unicorn.Id); err != nil {
		panic(err)
	}
	fmt.Printf("Create: %v\n", unicorn)

	_, err = client.Get(unicorn.Id)
	if err != nil {
		panic(err)
	}

	//var debug bool
	//
	//flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	//flag.Parse()
	//
	//opts := providerserver.ServeOpts{
	//	// TODO: Update this string with the published name of your provider.
	//	Address: "registry.terraform.io/hashicorp/scaffolding",
	//	Debug:   debug,
	//}
	//
	//err := providerserver.Serve(context.Background(), provider.New(version), opts)
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
}
