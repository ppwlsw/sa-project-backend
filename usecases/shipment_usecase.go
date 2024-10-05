package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type ShipmentUsecase interface {
	CreateShipment(s entities.Shipment) (entities.Shipment, error)
	UpdateShipment(id int, s entities.Shipment) (entities.Shipment, error)
	GetShipmentByID(id int) (entities.Shipment, error)
	GetAllShipments() ([]entities.Shipment, error)
}

type ShipmentService struct {
	repo repositories.ShipmentRepository
}

func InitiateShipmentService(repo repositories.ShipmentRepository) ShipmentUsecase {
	return &ShipmentService{
		repo: repo,
	}
}

func (ss *ShipmentService) CreateShipment(s entities.Shipment) (entities.Shipment, error) {
	cs, err := ss.repo.CreateShipment(s)

	if err != nil {
		return entities.Shipment{}, err
	}

	verifiedShipment, err := ss.repo.GetShipmentByID(cs.Id)

	if err != nil {
		return entities.Shipment{}, err
	}

	return verifiedShipment, nil
}

func (ss *ShipmentService) UpdateShipment(id int, us entities.Shipment) (entities.Shipment, error) {
	updateShipment, err := ss.repo.UpdateShipment(id, us)

	if err != nil {
		return entities.Shipment{}, err
	}

	return updateShipment, nil
}

func (ss *ShipmentService) GetShipmentByID(id int) (entities.Shipment, error) {
	shipment, err := ss.repo.GetShipmentByID(id)

	if err != nil {
		return entities.Shipment{}, err
	}

	return shipment, nil
}

func (ss *ShipmentService) GetAllShipments() ([]entities.Shipment, error) {
	shipmentList, err := ss.repo.GetAllShipments()

	if err != nil {
		return nil, err
	}

	return shipmentList, nil
}
