package service

import (
	"context"
	"encoding/json"
	"errors"
	"tugas-6/internal/entity"
	"tugas-6/internal/repository"
	"tugas-6/pkg/cache"
	"tugas-6/pkg/token"
)

type TodoService interface {
	CreateTodo(ctx context.Context,userID uint, title string) (*entity.Todo, error)
	GetTodos(ctx context.Context,userID uint) ([]entity.Todo, error)
	UpdateTodoRoleCheck(userID, todoID uint, title string, done bool, role string) error
	DeleteTodoRoleCheck(userID, todoID uint, role string) error
}

type todoService struct {
	repo repository.TodoRepository
	tokenUseCase   token.TokenUseCase
	cacheable      cache.Cacheable
}

func NewTodoService(
	repo repository.TodoRepository, 
	tokenUseCase token.TokenUseCase,
	cacheable cache.Cacheable,
	) TodoService {
	return &todoService{repo, tokenUseCase, cacheable}
}

func (s *todoService) CreateTodo(ctx context.Context,userID uint, title string) (*entity.Todo, error) {
	todo := &entity.Todo{UserID: userID, Title: title}
	err := s.repo.Create(ctx, todo)
	return todo, err
}

func (s *todoService) GetTodos(ctx context.Context,userID uint) (result []entity.Todo, err error) {
	keyFindAll := "tugas-6:users:find-all"
	data := s.cacheable.Get(keyFindAll)
	if data == ""{
		result, err = s.repo.GetAll(ctx, userID)
		if err != nil{
			return nil, err
		}
	}else{
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *todoService) UpdateTodoRoleCheck(userID, todoID uint, title string, done bool, role string) error {
	todo, err := s.repo.GetByID(todoID)
	if err != nil || (role != "admin" && todo.UserID != userID) {
		return errors.New("unauthorized or not found")
	}
	todo.Title = title
	todo.Done = done
	return s.repo.Update(todo)
}

func (s *todoService) DeleteTodoRoleCheck(userID, todoID uint, role string) error {
	todo, err := s.repo.GetByID(todoID)
	if err != nil || (role != "admin" && todo.UserID != userID) {
		return errors.New("unauthorized or not found")
	}
	return s.repo.Delete(todoID)
}
