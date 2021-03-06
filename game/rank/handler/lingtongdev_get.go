package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/pbutil"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RANK_LINGTONGDEV_GET_TYPE), dispatch.HandlerFunc(handleRankLingTongDevGet))
}

//处理灵童养成类排行榜信息
func handleRankLingTongDevGet(s session.Session, msg interface{}) (err error) {
	log.Debug("rank:处理获取灵童养成类排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankLingTongDevGet := msg.(*uipb.CSRankLingTongDevGet)
	classType := csRankLingTongDevGet.GetClassType()
	page := csRankLingTongDevGet.GetPage()
	isArea := csRankLingTongDevGet.GetIsArea()

	err = rankLingTongDevGet(tpl, types.LingTongDevSysType(classType), page, isArea)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"page":      page,
				"isArea":    isArea,
				"error":     err,
			}).Error("rank:处理获取灵童养成类排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("rank:处理获取灵童养成类排行榜消息完成")
	return nil

}

//获取灵童养成类排行榜界面信息的逻辑
func rankLingTongDevGet(pl player.Player, sysType types.LingTongDevSysType, page int32, isArea bool) (err error) {
	if !sysType.Vaild() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"sysType":  sysType,
			"page":     page,
			"isArea":   isArea,
		}).Warn("rank:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if page < 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"sysType":  sysType,
			"page":     page,
			"isArea":   isArea,
		}).Warn("rank:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	var classType ranktypes.RankClassType
	if isArea {
		classType = ranktypes.RankClassTypeArea
	} else {
		classType = ranktypes.RankClassTypeLocal
	}

	rankType := ranktypes.LongTongDevTypeToRankType(sysType)
	lingTongDevList, rankTime := rank.GetRankService().GetOrderListByPage(rankType, classType, 0, page)
	scRankLingTongDevGet := pbutil.BuildSCRankLingTongDevGet(int32(sysType), isArea, page, lingTongDevList, rankTime)
	pl.SendMsg(scRankLingTongDevGet)
	return
}
