package constants

import "os"

var (
	JWT_SECRET     string
	BCRYPT_COST    string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	MYSQL_HOST     string
	MYSQL_PORT     string
	ADMIN_USERNAME string
	ADMIN_PASSWORD string
)

func init() {
	JWT_SECRET = os.Getenv("JWT_SECRET")
	BCRYPT_COST = os.Getenv("BCRYPT_COST")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
}
