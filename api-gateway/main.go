package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/fooddalivary-api/api"
	"github.com/dilshodforever/fooddalivary-api/api/handler"
	pb "github.com/dilshodforever/fooddalivary-api/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	UserConn, err := grpc.NewClient(fmt.Sprintf("auth-service%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	ProductService, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer ProductService.Close()
	//auth := pb.NewAuthServiceClient(UserConn)
	Order := pb.NewOrderServiceClient(ProductService)
	OrderImets := pb.NewOrderItemServiceClient(ProductService)
	product:=pb.NewProductServiceClient(ProductService)
	h := handler.NewHandler(Order, OrderImets, product)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
