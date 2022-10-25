package services

import (
	"errors"
	"github.com/HiBang15/sample-gorm.git/constant"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/dto"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/repository"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/transformers"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type UserService struct {
	UserRepo        repository.UserRepository
	UserTransformer *transformers.UserTransformer
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo:        *repository.NewUserRepository(),
		UserTransformer: transformers.NewUserTransformer(),
	}
}

func (userService *UserService) CreateUser(
	request *dto.CreateUserRequest,
) (user *dto.User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(constant.ErrHashCode)
	}
	request.Password = string(hashedPassword)

	// create verify code

	userEntities := &entities.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    string(hashedPassword),
		PhoneNumber: request.PhoneNumber,
	}

	err = userService.UserRepo.Create(userEntities)
	if err != nil {
		if err.Error() == constant.ErrDuplicateEmailMessage {
			return nil, errors.New(constant.ErrorEmailExist)
		}
		if err.Error() == constant.ErrDuplicatePhoneNumber {
			return nil, errors.New(constant.ErrorPhoneNumberExist)

		}
		return nil, errors.New(constant.ErrCreateUserFail)
	}

	return userService.UserTransformer.UserEntityToDto(userEntities), nil
}

func (userService *UserService) GetUsers() ([]entities.User, error) {
	users, err := userService.UserRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userService *UserService) GetUser(id string) (*dto.User, error) {
	user, err := userService.UserRepo.GetUser(id)

	if err != nil {
		return nil, err

	}
	return userService.UserTransformer.UserEntityToDto(user), nil
}

func (userService *UserService) DeleteUser(id string) error {
	err := userService.UserRepo.DeleteUser(id)

	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) UpdateUser(id string, request *dto.UpdateUserRequest) (*dto.User, error) {
	//generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(constant.ErrHashCode)
	}
	request.Password = string(hashedPassword)

	userEntities := &entities.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    string(hashedPassword),
		PhoneNumber: request.PhoneNumber,
	}
	if err := userService.UserRepo.UpdateUser(id, userEntities); err != nil {
		return nil, err
	}
	user, err := userService.UserRepo.GetUser(id)

	return userService.UserTransformer.UserEntityToDto(user), nil
}

func (userService *UserService) GetUsersWithPagination(limit, pageNumber int) ([]entities.User, error) {
	users, err := userService.UserRepo.GetUsersWithPagination(limit, pageNumber)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *UserService) GetUserWithSearch(limit, pageNumber, query string) ([]dto.User, int, int, int, error) {
	var limitInt int
	var pageNumberInt int

	if num, err := strconv.Atoi(limit); err != nil {
		limitInt = -1
		pageNumberInt = -1
	} else {
		limitInt = num
	}

	if num, err := strconv.Atoi(pageNumber); err != nil {
		limitInt = -1
		pageNumberInt = -1
	} else {
		pageNumberInt = num
	}

	query = strings.ReplaceAll(query, " ", "")

	users, len, err := userService.UserRepo.GetUserWithSearch(limitInt, pageNumberInt, query)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	var userDto []dto.User
	for _, user := range users {
		userDto = append(userDto, *userService.UserTransformer.UserEntityToDto(&user))
	}
	return userDto, limitInt, pageNumberInt, len, nil
}
