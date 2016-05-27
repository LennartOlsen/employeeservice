package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"github.com/lennartolsen/employeeservice/employee"
)

func main() {
	var projID = "ggcompanypage"
	//Bootstrap Datastore client
	ctx := context.Background()

	Client, err := datastore.NewClient(ctx, projID)
	defer Client.Close()

	if(err != nil){
		log.Fatalf("Error connection to datastore : %s", err)
	}

	log.Print("We are flying");

	var repos = employee.Repository{Client}

	if( err != nil ){
		log.Printf("There was an error %v", err)
	}

	var controller = employee.Controller{&repos}

	router := httprouter.New()

	router.GET("/employees", controller.GetAll)
	router.GET("/employees/:id", controller.GetById)
	router.POST("/employees", controller.Create)
	router.PUT("/employees/:id", controller.Put)

	log.Fatal(http.ListenAndServe(":8082", router))
}
