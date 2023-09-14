package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/larturi/golang-ecommerce/bd"
	"github.com/larturi/golang-ecommerce/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	isValidateCode, isValidateMsg := validations(body, User)

	if isValidateCode != 200 {
		return isValidateCode, isValidateMsg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, Id int) (int, string) {
	var t models.Category

	isValidateCode, isValidateMsg := validations(body, User)

	if isValidateCode != 200 {
		return isValidateCode, isValidateMsg
	}

	t.CategID = Id

	fmt.Println("RouterCategory > ", t.CategID, t.CategName, t.CategPath)

	err2 := bd.UpdateCategory(t)

	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el update de la categoría " + strconv.Itoa(Id) + " > " + err2.Error()
	}

	return 200, "Update OK"
}

func validations(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	return 200, "OK"
}
