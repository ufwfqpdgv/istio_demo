package collection_api

import (
	"github.com/cihub/seelog"
	"math/rand"
	"samh_ticket_rank_0.1/recommend/go/database"
	"samh_ticket_rank_0.1/recommend/go/define"
	"strconv"
)


func CidTicketApi(ComicId string) define.JsonOut {
	//获取ComicId对应的漫画信息
	var ok bool
	var ticket, rank string
	ticket = database.GetCidTicket(define.CidTicketDb, ComicId) //获取对应漫画ID的月票数和排名
	rank, ok = define.Cid_rank[ComicId]
	var json define.JsonOut
	var data define.Num_rank

	json.Msg = "ok"
	json.Status = "0"
	if "0" != ticket {
		data.Num = ticket
	} else {
		num := rand.Intn(3) + 10
		data.Num = strconv.Itoa(num)
		seelog.Warn(ComicId, " has no ticket num")
	}
	if ok {
		data.Rank = rank
	} else {
		data.Rank = "1111"
		seelog.Warn(ComicId, " has no ticket rank")
	}
	json.Data = data
	return json
}
