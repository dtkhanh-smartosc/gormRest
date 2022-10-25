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
	idString := strconv.Itoa(int(request.GetID()))
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

func (server *UserGrpcServer) SearchUser(req *pb.SearchRequest, stream pb.UserGRPCService_SearchUserServer) error {
	log.Println("call SearchUser api")
	//[]dto.User, int, int, int, error
	userDto, _, _, len, err := server.userService.GetUserWithSearch(req.GetLimit(), req.GetPageNumber(), req.GetQuery())
	if err != nil {
		log.Println("err SearchUser")
	}
	for _, user := range userDto {
		stream.Send(&pb.SearchResponse{
			Id:          user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Total:       int32(len),
		})
	}
	return nil

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
func NewUserGrpcServer() *UserGrpcServer {
	return &UserGrpcServer{userService: services.NewUserService()}
}
