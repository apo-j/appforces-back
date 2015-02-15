package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DbContext struct{
	Dbmap *gorp.DbMap
}

func NewDbContext() DbContext{
	db, err := sql.Open("mysql", "root:root@/appforces_dev")
	//db, err := sql.Open("mysql", "ebroot:ebroot@tcp(aa148tj059wcdvz.cxllfavw7ze3.eu-west-1.rds.amazonaws.com:3306)/appforces_dev")
	if err != nil { log.Fatal(err) }

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(App{}, "app")
	dbmap.AddTableWithName(AppConfig{}, "appconfig")
	return DbContext{Dbmap: dbmap}
}





