package service

import (
	"context"
	"io"

	"github.com/waldrey/goexpert/grpc/internal/database"
	"github.com/waldrey/goexpert/grpc/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var categoriesResponse []*pb.Category
	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{
		Categories: categoriesResponse,
	}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}
