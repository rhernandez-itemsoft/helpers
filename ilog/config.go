package ilog

import (
	"itemsoftmx/helpers/icommon"
	ilogstt "itemsoftmx/helpers/ilog/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() ilogstt.Config {
	return ilogstt.Config{
		Enable: true,
		//Cuando SaveTo = File, necesitamos tener un nombre de archivo, junto con la ruta donde se almacenrá el log
		FileName: icommon.GetPath("logs/errors_go/"),

		//Permite mostrar mensajes en consola
		PrintConsole: true,

		//Formato de la fecha en que se registro el log
		TimeFormat: "2006-01-02 15:04:05.9999999 Z0700 MST",
	}
}
