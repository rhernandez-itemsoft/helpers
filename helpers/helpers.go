package helpers

import (
	"itemsoftmx/helpers/iresponse"
	configstt "itemsoftmx/isystem/config/structs"

	"github.com/go-xorm/xorm"
)

//Helpers Permite definir los objetos que serán injectados en este controlador
type Helpers struct {
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine

	//contiene la configuración general
	Conf configstt.GralConfigStt
}

//New crea una nueva instancia de  Helpers
func New(db *xorm.Engine, conf configstt.GralConfigStt, _iresp iresponse.Definition) Helpers {
	return Helpers{
		DB:   db,
		Conf: conf,
	}
}
