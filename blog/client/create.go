package main

import (
	"context"
	"log"

	pb "github.com/Reuben-Stephen-John/grpc-mongodb/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Reuben",
		Title: "My First Blog",
		Content: "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n",err)
	}

	log.Printf("Blog has been created: %s\n",res.Id)
	return res.Id
}