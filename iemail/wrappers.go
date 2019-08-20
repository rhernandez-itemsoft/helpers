package iemail

import (
	iemailstt "github.com/rhernandez-itemsoft/helpers/iemail/structs"
	"strings"
)

//WrappForgot reemplaza en el template de forgto: las variables por valores
func WrappForgot(message string, data iemailstt.WrappForgot) string {
	message = strings.ReplaceAll(message, "${fullname}", data.Fullname)
	message = strings.ReplaceAll(message, "${code}", data.Code)
	return message
}
