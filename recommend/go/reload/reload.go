package reload

import (
	"github.com/cihub/seelog"
	"samh_ticket_rank_0.1/recommend/go/database"
	"samh_ticket_rank_0.1/recommend/go/define"
)


// 只更新月票排行
func ReloadRank() {
	new_cid_rank := database.GetCidRank(define.CidRankDb)

	if nil == new_cid_rank {
		seelog.Error("failed to reload cid rank map")
		return
	}
	define.Cid_rank = new_cid_rank // 赋值
}