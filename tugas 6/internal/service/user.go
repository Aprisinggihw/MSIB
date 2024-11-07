package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
	"tugas-6/internal/entity"
	"tugas-6/internal/repository"
	"tugas-6/pkg/cache"
	"tugas-6/pkg/token"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	Register(ctx context.Context, req *entity.UserReg) error
	Login(ctx context.Context, username, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	tokenUseCase   token.TokenUseCase
	cacheable      cache.Cacheable
}

func NewUserService(
	userRepository repository.UserRepository,
	tokenUseCase token.TokenUseCase,
	cacheable cache.Cacheable,
) UserService {
	return &userService{userRepository, tokenUseCase, cacheable}
}

func (s *userService) FindAll(ctx context.Context) (result []entity.User, err error) {
	keyFindAll := "tugas-6:users:find-all"
	data := s.cacheable.Get(keyFindAll)
	if data == "" {
		result, err = s.userRepository.FindAll(ctx)
		if err != nil {
			return nil, err
		}

		marshalledData, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		err = s.cacheable.Set(keyFindAll, marshalledData, 5*time.Minute)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Logika registrasi user
func (s *userService) Register(ctx context.Context, req *entity.UserReg) error {
	// Periksa apakah username sudah ada
	_, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err == nil {
		return errors.New("username already exists")
	}
	// Invalidate cache "tugas-6:users:find-all"
	keyFindAll := "tugas-6:users:find-all"
	err = s.cacheable.Delete(keyFindAll) // Menghapus cache lama
	if err != nil {
		return errors.New("falied deleting key cache")
	}
	return s.userRepository.CreateUser(ctx, req)
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("username or password invalid")
	}

	if user.Password != password {
		return "", errors.New("username or password invalid")
	}

	expiredTime := time.Now().Local().Add(time.Minute * 5)

	claims := token.JwtCustomClaims{
		UserID:   uint(user.ID),
		Username: user.Username,
		Role:     user.Role,
		FullName: user.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tugas-6",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("ada kesalahan di server")
	}
	keyGetTodos := "tugas-6:todos:get-todos"
	err = s.cacheable.Delete(keyGetTodos)
	if err != nil {
		return "", errors.New("falied deleting key cache")
	}
	return token, nil
}
