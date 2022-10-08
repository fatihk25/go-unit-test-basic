package repository

import "basic-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
