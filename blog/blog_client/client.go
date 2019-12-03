package main

import (
	"context"
	"fmt"
	"grpc-go/blog/blogpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client...")
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect to server ", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "Mainak Dhar",
		Title:    "Aati kikore badhe",
		Content:  "Ki r boli bara!!!",
	}

	resp, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		fmt.Println("Issue while creating blog request", err)
	}
	fmt.Println("Blog entry created! ", resp)
	blogID := resp.GetBlog().GetId()

	fmt.Println("Reading the blog...")
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogResp, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error while reading...", readBlogErr)
	}
	fmt.Printf("Blog output:: ", readBlogResp)

	newBlog := &blogpb.Blog{
		Id: blogID,
		AuthorId: "Ratna Dhar",
		Title:    "Jhat ki kore jalaye!!",
		Content:  "Ami eta janina...",
	}

	updateBlogResp, updateErr := c.UpdateBlog(ctx context.Context, &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Println("error while updating", updateErr)
	}
	fmt.Println("Blog was read", updateBlogResp)
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: \n", deleteRes)
}
