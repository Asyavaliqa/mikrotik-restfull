package interfaces

import "github.com/didintri196/mikrotik-restfull/domain/models"

type IResourceRepository interface {
	Read(resource *models.Resource) (err error)
}
