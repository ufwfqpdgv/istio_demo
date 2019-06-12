package collection_api

import (
	"fmt"
	"github.com/zhanglanhui/go-utils/utils"
	"samh_ticket_rank_0.1/recommend/go/database"
	"samh_ticket_rank_0.1/recommend/go/define"
	"samh_ticket_rank_0.1/recommend/go/reload"
)



func init() {
	ConfigEngine := utils.ConfigEngine{}
	ConfigEngine.Load(define.SConfigFile)
	define.Product_id = ConfigEngine.GetString("Product_id")
	var yamlConfig utils.ConfigEngine
	var err error
	err = yamlConfig.Load(define.SConfigFile)
	utils.CheckFatalErr(err)
	yamlConfig.SentryRavenInit("Sentry_dsn")
	define.CidTicketDb = yamlConfig.GetDbInfoFromConf("Comic_info_Mssql")
	define.CidRankDb = yamlConfig.GetDbInfoFromConf("Comic_rank_mysql")


	utils.LoggerSetup(define.SSeelogFile)
	define.Cid_rank = make(map[string]string, 0)
	database.InitUserCollectionSql()
	reload.ReloadRank()
	if 0 == len(define.Cid_rank) {
		utils.CheckFatalErr(fmt.Errorf("cid rank map is empty"))
	}
}
