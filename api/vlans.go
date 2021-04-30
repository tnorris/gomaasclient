package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// VLANs represents the MaaS Vlans endpoint
type VLANs interface {
	Get(fabricID int) ([]entity.VLAN, error)
	Post(fabricID int, params *entity.VLANParams) (*entity.VLAN, error)
}
