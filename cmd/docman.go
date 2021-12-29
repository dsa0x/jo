package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dsa0x/docman"
)

func init() {
	docman.NewConfig()
	flag.StringVar(&docman.Cfg.PORT, "port", "", "port to listen on")
	flag.StringVar(&docman.Cfg.DbHost, "dbhost", "", "database host")
	flag.StringVar(&docman.Cfg.DbPort, "dbport", "", "database port")
	flag.StringVar(&docman.Cfg.DbUser, "dbuser", "", "database user")
	flag.StringVar(&docman.Cfg.DbName, "dbname", "", "database name")
	flag.StringVar(&docman.Cfg.DbPassword, "dbpassword", "", "database password")
}

func main() {
	flag.Parse()
	err := docman.NewDB()
	if err != nil {
		panic(err)
	}
	r := docman.NewWeb()
	fmt.Println("Server listening on port:", docman.Cfg.PORT)
	log.Fatal(http.ListenAndServe(":"+docman.Cfg.PORT, r))
}
