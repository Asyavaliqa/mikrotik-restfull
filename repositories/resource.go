package repositories

import (
	"github.com/didintri196/mikorm"
	"github.com/didintri196/mikrotik-restfull/domain/interfaces"
	"github.com/didintri196/mikrotik-restfull/domain/models"
)

type ResourceRepository struct {
	MikORM *mikorm.MikORM
}

func (repo ResourceRepository) Read(resource *models.Resource) (err error) {
	return repo.MikORM.Command("/system/resource").Print(resource).Error
}

func NewResourceRepository(mikORM *mikorm.MikORM) interfaces.IResourceRepository {
	return &ResourceRepository{
		MikORM: mikORM,
	}
}
