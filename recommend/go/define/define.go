package define

import "github.com/zhanglanhui/go-utils/utils"

var Cid_rank map[string]string
var CidTicketDb *utils.MysqlDbInfo
var CidRankDb *utils.MysqlDbInfo

var CollectionMin int
var Product_id string
var SConfigFile = "config/conf_ticket_rank.yaml"
var SSeelogFile = "config/seelog_samh_ticket.xml"

type Num_rank struct {
	Num  string `json:"thismonth_ticket_num"`
	Rank string `json:"Rank"`
}

type JsonOut struct {
	Status string   `json:"status"`
	Msg    string   `json:"msg"`
	Data   Num_rank `json:"data"`
}
