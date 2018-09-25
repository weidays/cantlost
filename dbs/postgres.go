/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-10-13 21:09:02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 09:40:39
***********************************************/
package dbs

import (
	"fmt"
	"log"
	"projects/cantlost/libs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PGMaster *gorm.DB

func init() {
	var err error
	username := libs.Conf.Read("postgres", "username")
	// password := libs.Conf.Read("postgres", "password")
	dataname := libs.Conf.Read("postgres", "dataname")
	port := libs.Conf.Read("postgres", "port")
	host := libs.Conf.Read("postgres", "host")
	dns := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dataname + " sslmode=disable"
	fmt.Println(dns)
	PGMaster, err = gorm.Open("postgres", dns)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = PGMaster.DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	PGMaster.DB().SetMaxIdleConns(20)
	PGMaster.DB().SetMaxOpenConns(120)
}
