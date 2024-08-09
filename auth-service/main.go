package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dilshodforever/fooddalivary-auth/api"
	"github.com/dilshodforever/fooddalivary-auth/api/handler"
	pb "github.com/dilshodforever/fooddalivary-auth/genprotos/auth"
	pbu "github.com/dilshodforever/fooddalivary-auth/genprotos/user"
	"github.com/dilshodforever/fooddalivary-auth/service"
	"github.com/dilshodforever/fooddalivary-auth/storage/postgres"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, service.NewAuthService(db))
	pbu.RegisterUserServiceServer(s, service.NewUserService(db))
	log.Printf("server listening at %v", liss.Addr())
	go func() {
		if err := s.Serve(liss); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	UserConn, err := grpc.NewClient(fmt.Sprintf("auth-service%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	redisstorage := handler.NewInMemoryStorage(rdb)
	aus := pb.NewAuthServiceClient(UserConn)
	us := pbu.NewUserServiceClient(UserConn)
	h := handler.NewHandler(aus, us, redisstorage)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8081")
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
