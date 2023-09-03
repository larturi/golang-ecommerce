package bd

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/larturi/golang-ecommerce/awssecret"
	"github.com/larturi/golang-ecommerce/models"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = awssecret.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "sql10642491"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	// fmt.Println(dsn)
	return dsn
}
