package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type PackageRepository interface {
	CreatePackage(p entities.Package) (entities.Package, error)
	GetPackageByID(id int) (entities.Package, error)
	GetAllPackagesByOrderID(id int) ([]entities.Package, error)
	GetAllPackages() ([]entities.Package, error)
}
