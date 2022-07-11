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

func JoinFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().UserDepList(global.DB, 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.Write(jData)
}

func PreloadFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().PreloadUserDep(global.DB, 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.Write(jData)
}

func PreloadsFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().PreloadUserDeps(global.DB, 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.Write(jData)
}

func main() {
	http.HandleFunc("/join", JoinFunc)
	http.HandleFunc("/preload", PreloadFunc)
	http.HandleFunc("/preloads", PreloadsFunc)

	port := 8001
	fmt.Printf("连接查询: http://127.0.0.1:%d/join\n", port)
	fmt.Printf("预加载: http://127.0.0.1:%d/preload\n", port)
	fmt.Printf("预加载多对多: http://127.0.0.1:%d/preloads\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
