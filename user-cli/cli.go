package main

import (
	"log"
	"os"

	pb "github.com/justcy/shippy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"golang.org/x/net/context"
)


func main() {

	cmd.Init()

	// 创建 user-service 微服务的客户端
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// 设置命令行参数
	//service := micro.NewService(
	//	micro.Flags(
	//		cli.StringFlag{
	//			Name:  "name",
	//			Usage: "You full name",
	//		},
	//		cli.StringFlag{
	//			Name:  "email",
	//			Usage: "Your email",
	//		},
	//		cli.StringFlag{
	//			Name:  "password",
	//			Usage: "Your password",
	//		},
	//		cli.StringFlag{
	//			Name: "company",
	//			Usage: "Your company",
	//		},
	//	),
	//)
	//service.Init(
	//	micro.Action(func(c *cli.Context) {
	//		name := c.String("name")
	//		email := c.String("email")
	//		password := c.String("password")
	//		company := c.String("company")
	//
	//		r, err := client.Create(context.TODO(), &pb.User{
	//			Name: name,
	//			Email: email,
	//			Password: password,
	//			Company: company,
	//		})
	//		if err != nil {
	//			log.Fatalf("Could not create: %v", err)
	//		}
	//		log.Printf("Created: %v", r.User.Id)
	//
	//		getAll, err := client.GetAll(context.Background(), &pb.Request{})
	//		if err != nil {
	//			log.Fatalf("Could not list users: %v", err)
	//		}
	//		for _, v := range getAll.Users {
	//			log.Println(v)
	//		}
	//
	//		os.Exit(0)
	//	}),
	//)
	//
	//// 启动客户端
	//if err := service.Run(); err != nil {
	//	log.Println(err)
	//}

	// 暂时将用户信息写死在代码中
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	resp, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}
	log.Println("created: ", resp.User.Id)

	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("call GetAll error: %v", err)
	}
	for _, u := range allResp.Users {
		log.Printf("%v\n", u)
	}

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	log.Println("token: ", authResp.Token)

	// 直接退出即可
	os.Exit(0)


}
