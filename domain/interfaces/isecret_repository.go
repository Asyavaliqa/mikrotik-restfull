package interfaces

import "github.com/didintri196/mikrotik-restfull/domain/models"

type ISecretRepository interface {
	Add(secret models.Secret) (err error)

	Browse() (secrets []models.Secret, err error)

	Read(filter models.Secret) (secrets []models.Secret, err error)

	Edit(filter models.Secret, data models.Secret) (err error)
}
