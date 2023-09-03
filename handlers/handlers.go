package handlers

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	// id := request.PathParameters["id"]
	// idn, _ := strconv.Atoi(id)

	// isOk, statusCode, user := validoAuthorization(path, method, headers)
	// if !isOk {
	// 	return statusCode, user
	// }

	fmt.Println("path[0:4] = " + path[0:4])

	// switch path[0:4] {
	// case "user":
	// 	return ProcesoUsers(body, path, method, user, id, request)
	// case "prod":
	// 	return ProcesoProducts(body, path, method, user, idn, request)
	// case "stoc":
	// 	return ProcesoStock(body, path, method, user, idn, request)
	// case "addr":
	// 	return ProcesoAddress(body, path, method, user, idn, request)
	// case "cate":
	// 	return ProcesoCategory(body, path, method, user, idn, request)
	// case "orde":
	// 	return ProcesoOrder(body, path, method, user, idn, request)
	// }

	return 400, "Method Invalid"
}
