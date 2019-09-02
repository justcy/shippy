package main

import (
	"context"
	// 导如 protoc 自动生成的包
	pb "github.com/justcy/shippy/consignment-service/proto/consignment"
	vesselPb "github.com/justcy/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
)
// 仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物

	GetAll() []*pb.Consignment                                   // 获取仓库中所有的货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//
// 定义微服务
//
type service struct {
	repo Repository
	// consignment-service 作为客户端调用 vessel-service 的函数
	vesselClient vesselPb.VesselServiceClient
}

//
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	resp.Created = true
	resp.Consignment = consignment
	return nil
}

// 获取目前所有托运的货物
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response)  error {
	allConsignments := s.repo.GetAll()
	resp = &pb.Response{Consignments: allConsignments}
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		)
	vesselClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	// 解析命令行参数
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(),&service{repo,vesselClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
