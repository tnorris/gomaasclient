package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Fabric interface {
	Get(id int) (*entity.Fabric, error)
	Update(id int, fabricParams *entity.FabricParams) (*entity.Fabric, error)
	Delete(id int) error
}
