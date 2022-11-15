package database

import "github.com/LulianoM/api-geoprocess/src/models"

var Tables = []interface{}{
	&models.Metadata{},
	&models.Seller{},
	&models.Service{},
	&models.Operation{},
	&models.Products{},
	&models.Contact{},
	&models.Commercials{},
}
