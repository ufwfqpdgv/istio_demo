package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/cihub/seelog"

	"github.com/robfig/cron"
	"github.com/zhanglanhui/go-utils/utils"
	"samh_ticket_rank_0.1/recommend/go/reload"
	"samh_ticket_rank_0.1/server/routers"
)

//version:
//algorithm:samh_ticket_rank
//versions:0.1.1
func main() {
	defer func() {
		if rec := recover(); rec != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			err := fmt.Errorf("rec:%s\nstack:%s", rec, buf[:n])
			utils.CheckErrSendEmail(err)
		}
	}()
	sConfigFile := "config/conf_ticket_rank.yaml"
	ConfigEngine := utils.ConfigEngine{}
	ConfigEngine.Load(sConfigFile)
	port := ConfigEngine.GetString("Smh_ticket_rank_server.Port")
	if port == "" {
		panic("no port in conf.ini")
	}
	c := cron.New()
	c.AddFunc("0 0 11 * * ?", func() {
		seelog.Debugf("%s:reload new comic rank\n", time.Now().Format(time.RFC3339))
		reload.ReloadRank()
	})
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        router,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	c.Start()
	s.ListenAndServe()
}
