package database

import (
	"fmt"
	"github.com/zhanglanhui/go-utils/utils"
	"samh_ticket_rank_0.1/recommend/go/define"
)

var sql_num, sql_rank string

func GetCidTicket(db *utils.MysqlDbInfo, comicId string) (num string) {
	num = "0"
	rows := db.SqlDataDb.QueryRow(sql_num, comicId)
	err := rows.Scan(&num)
	utils.CheckCommonErr(err)
	return num
}

func GetCidRank(db *utils.MysqlDbInfo) (cid_rank map[string]string) {
	cid_rank = make(map[string]string)
	rows, err := db.SqlDataDb.Query(sql_rank)
	if err != nil {
		utils.CheckCommonErr(err)
		rows, err = db.SqlDataDb.Query(sql_rank) // 若有错误则再次请求
		if err != nil {
			return nil
		}
	}
	defer rows.Close()
	for rows.Next() {
		var tmpId, tmpRank string
		err := rows.Scan(&tmpId, &tmpRank)
		utils.CheckCommonErr(err)
		if err == nil {
			cid_rank[tmpId] = tmpRank
		} else {
			utils.CheckErrSendEmail(fmt.Errorf("error is %s\ntmpId: %s\ntmpRank: %s", err, tmpId, tmpRank))
			return nil
		}
	}
	return cid_rank
}

func InitUserCollectionSql() {
	sql_num = fmt.Sprintf(`SELECT thismonth_ticket_num
		FROM %s 
		WHERE product_id=%s and comic_id=?;`,
		define.CidTicketDb.TableName["Comic"], define.Product_id)

	sql_rank = fmt.Sprintf(`SELECT comic_id, FP_rank
	FROM
	(
		SELECT comic_id, count_num, @FP_rank :=  @FP_rank + 1 AS FP_rank
	FROM %v JOIN (SELECT @FP_rank := 0) AS vars
	ORDER BY count_num DESC
	) ranked`, define.CidRankDb.TableName["Month"])
}
