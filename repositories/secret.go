package repositories

import (
	"github.com/didintri196/mikrotik-restfull/domain/interfaces"
	"github.com/didintri196/mikrotik-restfull/domain/models"
	"github.com/didintri196/mikrotik-restfull/utils/routeros"
)

type SecretRepository struct {
	*routeros.RouterOS
}

func (repo SecretRepository) Browse() (secrets []models.Secret, err error) {
	err = repo.Command("/ppp/secret").Scan(&secrets).Error
	return
}

func (repo SecretRepository) Add(secret models.Secret) (err error) {
	return repo.Command("/ppp/secret").Add(&secret).Error
}

func NewSecretRepository(routerOS *routeros.RouterOS) interfaces.ISecretRepository {
	return &SecretRepository{
		RouterOS: routerOS,
	}
}