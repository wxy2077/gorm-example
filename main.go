package main

import (
	"flag"
	"fmt"
	"gorm-example/config"
	"gorm-example/controller"
	"gorm-example/global"
	"gorm-example/initialize"
	"gorm-example/middleware"
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

	_, err := initialize.InitJaeger(global.Config.Runtime)
	if err != nil {
		fmt.Printf("初始化jager出错:%s", err.Error())
	}

	global.DB = initialize.GormMysql(global.Config.MainMySQL)

}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

type Router struct {
	middleware []Middleware // 中间件列表
	mux        *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Use(middleware ...Middleware) {
	r.middleware = append(r.middleware, middleware...)
}

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	for _, m := range r.middleware {
		handler = m(handler)
	}

	r.mux.HandleFunc(pattern, handler)
}

func (r *Router) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}

func main() {

	router := NewRouter()

	router.Use(middleware.TraceMiddleware)

	router.HandleFunc("/join", controller.JoinFunc)
	router.HandleFunc("/preload", controller.PreloadFunc)
	router.HandleFunc("/preloads", controller.PreloadsFunc)

	_ = router.ListenAndServe(fmt.Sprintf(":%d", global.Config.Runtime.HttpPort))
}
