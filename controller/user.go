package controller

import (
	"encoding/json"
	"gorm-example/global"
	"gorm-example/logic"
	"gorm-example/utils"
	"net/http"
)

func JoinFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().UserDepList(global.DB.WithContext(req.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.WriteHeader(200)
	w.Write(jData)
}

func PreloadFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().PreloadUserDep(global.DB.WithContext(req.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)

	w.Write(jData)
}

func PreloadsFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, pagination := logic.NewUserLogic().PreloadUserDeps(global.DB.WithContext(req.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	jData, _ := json.Marshal(res)
	w.Write(jData)
}
