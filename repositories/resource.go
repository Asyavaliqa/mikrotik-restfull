package repositories

import (
	"github.com/didintri196/mikrotik-restfull/domain/interfaces"
	"github.com/didintri196/mikrotik-restfull/domain/models"
	"github.com/didintri196/mikrotik-restfull/utils/routeros"
)

type ResourceRepository struct {
	*routeros.RouterOS
}

func (repo ResourceRepository) Read(resource *models.Resource) (err error) {
	return repo.Command("/system/resource").Print(resource).Error
}

func NewResourceRepository(routerOS *routeros.RouterOS) interfaces.IResourceRepository {
	return &ResourceRepository{
		RouterOS: routerOS,
	}
}
