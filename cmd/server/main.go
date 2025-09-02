// Here we will set up the server
package main

import (
	log "log"
	"net"
	"user_management/api/proto"
	"user_management/pkg/handler"
	"user_management/pkg/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	proto.RegisterUserServiceServer(grpcServer, userHandler)

	//Setup gin server after making the httpHandler in Handler folder
	r := gin.Default()
	httpHandler := handler.NewHTTPHandler(userService)
	r.POST("/users", httpHandler.CreateUser)
	r.GET("/users/:id", httpHandler.GetUser)
	r.GET("/users", httpHandler.ListUsers)
	r.PUT("/users/:id", httpHandler.UpdateUser)
	r.DELETE("/users/:id", httpHandler.DeleteUser)

	//WE want to run the gin server and grpc server concurrently, so we will use goroutine
	go func() {
		log.Println("Grpc Server running on port: 8080")
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("Http Server running on port: 8000")
	err = r.Run(":8000")
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
