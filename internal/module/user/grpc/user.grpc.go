package grpc

import (
	context "context"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/services"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/transformers"
	pb "github.com/HiBang15/sample-gorm.git/proto/pb"
	"log"
	"strconv"
)

type UserGrpcServer struct {
	userService *services.UserService
}

func (server *UserGrpcServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("call getUser api")
	idString := strconv.Itoa(int(request.GetId()))
	log.Printf("id String", idString)
	user, err := server.userService.GetUser(idString)
	if err != nil {
		log.Println("err GetUser")
	}
	res := server.userService.UserTransformer.UserDtoToDPB(user)
	return res, nil
}
func (server *UserGrpcServer) GetUsers(req *pb.GetUsersRequest, stream pb.UserGRPCService_GetUsersServer) error {
	log.Println("call GetUsers api")
	users, err := server.userService.GetUsers()
	if err != nil {
		log.Println("err GetUser")
	}
	for _, user := range users {
		stream.Send(&pb.GetUserResponse{
			Id:          user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		},
		)
	}
	return nil
}
func (server *UserGrpcServer) SearchUser(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println("call SearchUser api")
	userDTO, limit, pageNumber, len, err := server.userService.GetUserWithSearch(req.GetLimit(), req.GetPageNumber(), req.GetQuery())
	if err != nil {
		log.Println("err SearchUser")
		return nil, err
	}
	var usersPB []*pb.GetUserResponse

	for _, user := range userDTO {
		usersPB = append(usersPB, server.userService.UserTransformer.UserDtoToDPB(&user))
	}
	return &pb.SearchResponse{Users: usersPB, Limit: int32(limit), PageNumber: int32(pageNumber), Total: int32(len)}, nil
}

func (server *UserGrpcServer) PostUser(ctx context.Context, request *pb.PostUserRequest) (*pb.GetUserResponse, error) {
	log.Println("call PostUser api")
	userTransformer := transformers.NewUserTransformer()
	userReqDto := userTransformer.CreateUserPBtoDto(request)
	errValid := userTransformer.VerifyCreateUserRequest(userReqDto)
	if errValid != nil {
		log.Println("err PostUser")
	}
	user, err := server.userService.CreateUser(userReqDto)
	if err != nil {
		return nil, err
	}
	res := server.userService.UserTransformer.UserDtoToDPB(user)
	return res, nil
}
func (server *UserGrpcServer) DeleteUser(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Println("call DeleteUser api")
	//idString := strconv.Itoa(int(request.GetId()))
	//err := server.userService.DeleteUser(idString)

	//if err != nil {
	//	log.Println("err DeleteUser")
	//	return &pb.DeleteResponse{
	//		Result: "Can del user",
	//	}, err
	//}
	return &pb.DeleteResponse{
		Result: "del user success",
	}, nil
}
func NewUserGrpcServer() *UserGrpcServer {
	return &UserGrpcServer{userService: services.NewUserService()}
}
