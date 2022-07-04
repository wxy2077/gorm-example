package main

import (
	"encoding/json"
	"flag"
	"fmt"
	logic "gorm-example/03_logic"
	"gorm-example/config"
	"gorm-example/global"
	"gorm-example/initialize"
	"gorm-example/utils"
	"net/http"
)

var (
	configFile = flag.String("f", "config/config.yaml", "the config file")
)

func init() {
	flag.Parse()

	global.Config = new(config.Config)
	loadEngine := config.NewLoad()
	loadEngine.LoadCfg(*configFile, global.Config)
	global.DB = initialize.GormMysql(global.Config.MainMySQL)

}

func HandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewJoinLogic().UserDepList(global.DB, 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.Write(jData)
}

func main() {
	http.HandleFunc("/index", HandleFunc)
	port := 8001
	fmt.Printf("http://127.0.0.1:%d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
