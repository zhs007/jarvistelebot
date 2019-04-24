package plugindtdata

import (
	"sort"

	"github.com/zhs007/jarvistelebot/jarviscrawlercore"
	"github.com/zhs007/jarvistelebot/plugins/dtdata/proto"
)

func hasBusinessInDTGameReport(game *plugindtdatapb.DTGameReport, businessid string) bool {
	for _, v := range game.Businessid {
		if v == businessid {
			return true
		}
	}

	return false
}

func hasGameInDTBusinessReport(business *plugindtdatapb.DTBusinessReport, gameCode string) bool {
	for _, v := range business.Gamecode {
		if v == gameCode {
			return true
		}
	}

	return false
}

func findDTGameReport(lstGame []*plugindtdatapb.DTGameReport, gameCode string) *plugindtdatapb.DTGameReport {
	for _, v := range lstGame {
		if v.GameCode == gameCode {
			return v
		}
	}

	return nil
}

func findDTBusinessReport(lstBusiness []*plugindtdatapb.DTBusinessReport, businessid string) *plugindtdatapb.DTBusinessReport {
	for _, v := range lstBusiness {
		if v.BusinessID == businessid {
			return v
		}
	}

	return nil
}

func countDTReportWithBusinessGameReport(reply *jarviscrawlercore.ReplyDTData, mainCurrency string,
	topNumsGame int, topNumsBusiness int) *plugindtdatapb.DTReport {

	dtreport := &plugindtdatapb.DTReport{
		MainCurrency: mainCurrency,
	}

	var lstGame []*plugindtdatapb.DTGameReport
	var lstBusiness []*plugindtdatapb.DTBusinessReport

	for _, v := range reply.GameReports {
		if v.Currency == mainCurrency {
			dtreport.TotalBet += v.TotalBet / 100.0
			dtreport.TotalWin += v.TotalWin / 100.0
			dtreport.SpinNums += int64(v.GameNums)

			cg := findDTGameReport(lstGame, v.Gamecode)
			if cg == nil {
				cg = &plugindtdatapb.DTGameReport{
					GameCode: v.Gamecode,
				}

				lstGame = append(lstGame, cg)
			}

			cg.TotalBet += v.TotalBet / 100.0
			cg.TotalWin += v.TotalWin / 100.0
			cg.SpinNums += int64(v.GameNums)

			if !hasBusinessInDTGameReport(cg, v.Businessid) {
				cg.Businessid = append(cg.Businessid, v.Businessid)

				cg.BusinessNums++
			}

			cb := findDTBusinessReport(lstBusiness, v.Businessid)
			if cb == nil {
				cb = &plugindtdatapb.DTBusinessReport{
					BusinessID: v.Businessid,
				}

				lstBusiness = append(lstBusiness, cb)
			}

			cb.TotalBet += v.TotalBet / 100.0
			cb.TotalWin += v.TotalWin / 100.0
			cb.SpinNums += int64(v.GameNums)

			if !hasGameInDTBusinessReport(cb, v.Gamecode) {
				cb.Gamecode = append(cb.Gamecode, v.Gamecode)

				cb.GameNums++
			}
		}
	}

	dtreport.GameNums = int32(len(lstGame))
	dtreport.BusinessNums = int32(len(lstBusiness))

	sort.Slice(lstGame, func(i, j int) bool {
		return lstGame[i].TotalBet > lstGame[j].TotalBet
	})

	sort.Slice(lstBusiness, func(i, j int) bool {
		return lstBusiness[i].TotalBet > lstBusiness[j].TotalBet
	})

	for i := 0; i < len(lstGame); i++ {
		for _, v := range lstGame[i].Businessid {
			ccb := findDTBusinessReport(lstBusiness, v)
			if ccb != nil {
				lstGame[i].BusinessReport = append(lstGame[i].BusinessReport, ccb)
			}
		}

		lstGame[i].Businessid = nil

		dtreport.TopGames = append(dtreport.TopGames, lstGame[i])
	}

	for i := 0; i < len(lstBusiness); i++ {
		for _, v := range lstBusiness[i].Gamecode {
			ccg := findDTGameReport(lstGame, v)
			if ccg != nil {
				lstBusiness[i].GameReport = append(lstBusiness[i].GameReport, ccg)
			}
		}

		lstBusiness[i].Gamecode = nil

		dtreport.TopBusiness = append(dtreport.TopBusiness, lstBusiness[i])
	}

	return dtreport
}
