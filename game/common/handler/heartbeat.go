package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HEARTBEAT_TYPE), dispatch.HandlerFunc(handleHeartbeat))
}

//心跳
func handleHeartbeat(s session.Session, msg interface{}) error {
	log.Debug("common:处理心跳信息")

	gcs := gamesession.SessionInContext(s.Context())
	flag := gcs.Ping()
	if !flag {
		log.WithFields(
			log.Fields{}).Warn("common:处理心跳信息,ping失败")
		return nil
	}
	pl := gcs.Player()
	tpl := pl.(player.Player)

	scHeartBeat := pbutil.BuildSCHeartBeat()
	err := tpl.SendMsg(scHeartBeat)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("common:处理心跳信息,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("common:处理心跳信息,完成")
	return nil
}
