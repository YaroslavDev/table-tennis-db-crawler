package config

import "upper.io/db.v3/mysql"

var ConnectionUrl = mysql.ConnectionURL{
	Database: 	`ttdb`,
	Host: 		`localhost`,
	User: 		`root`,
	Password:	`toor`,
}