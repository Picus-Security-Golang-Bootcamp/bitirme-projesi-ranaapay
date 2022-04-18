package service

import (
	"PicusFinalCase/src/models"
	"io"
	"reflect"
	"testing"
)

func TestCategoryService_CreateCategories(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{"WithNoFile_ShouldReturnsPanic", args{file: nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("CreateCategories() =  want %v", tt.wantPanic)
				}
			}()
			catService.CreateCategories(tt.args.file)
		})
	}
}

func TestCategoryService_FindCategories(t *testing.T) {
	tests := []struct {
		name string
		want *[]models.Category
	}{
		{"WithCategories_ShouldReturnsList", &mockCatRepo.users},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catService.FindCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoryService_FindCategory(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		want      *models.Category
		wantPanic bool
	}{
		{"IdEqualsIsDeletedTrue_ShouldReturnPanic", args{id: cat2.Id}, nil, true},
		{"IdEqualsIsDeletedFalse_ShouldReturnCategory", args{id: cat1.Id}, &cat1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SequenceInt() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := catService.FindCategory(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	cat1 = models.Category{
		Base: models.Base{
			Id:        "test1",
			IsDeleted: false,
		},
		CategoryName: "test",
		Description:  "test",
		Product:      nil}
	cat2 = models.Category{
		Base: models.Base{
			Id:        "test",
			IsDeleted: true,
		},
		CategoryName: "testtwo",
		Description:  "testtwo",
		Product:      nil,
	}

	mockCatRepo = &mockCatRepository{users: []models.Category{cat1, cat2}}
	catService  = NewCategoryService(mockCatRepo)
)

type mockCatRepository struct {
	users []models.Category
}

func (m *mockCatRepository) CreateCategories(categories []models.Category) bool {
	if len(categories) == 0 {
		return false
	}
	return true
}

func (m *mockCatRepository) FindCategories() *[]models.Category {
	return &m.users
}

func (m *mockCatRepository) FindCategory(id string) *models.Category {
	for _, u := range m.users {
		if u.Id == id && u.IsDeleted == false {
			return &u
		}
	}
	return nil
}
