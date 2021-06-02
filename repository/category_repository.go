package repository

import "gotesting/entity"

type CategoryRepository interface{
	FindById(id string) *entity.Category
}
	