package api

import (
	"fgame/fgame/charge_server/api/aotian"
	"fgame/fgame/charge_server/api/baji"
	"fgame/fgame/charge_server/api/bocai"
	"fgame/fgame/charge_server/api/chenxi"
	"fgame/fgame/charge_server/api/chenxi2"
	"fgame/fgame/charge_server/api/disui"
	"fgame/fgame/charge_server/api/feifan"
	"fgame/fgame/charge_server/api/feiyang"
	"fgame/fgame/charge_server/api/fengyuan"
	"fgame/fgame/charge_server/api/henggewan"
	"fgame/fgame/charge_server/api/jiandaowenqing"
	"fgame/fgame/charge_server/api/jianghu"
	"fgame/fgame/charge_server/api/jiuling"
	"fgame/fgame/charge_server/api/jiumeng"
	"fgame/fgame/charge_server/api/jiutian"
	"fgame/fgame/charge_server/api/judu"
	"fgame/fgame/charge_server/api/lieyan"
	"fgame/fgame/charge_server/api/lingmeng"
	"fgame/fgame/charge_server/api/longyu"
	"fgame/fgame/charge_server/api/luoliwan"
	"fgame/fgame/charge_server/api/menghuan"
	"fgame/fgame/charge_server/api/menghuan2"
	"fgame/fgame/charge_server/api/mengyuanwenxian"
	"fgame/fgame/charge_server/api/mingjian"
	"fgame/fgame/charge_server/api/mofangyouxifengmo"
	"fgame/fgame/charge_server/api/niuchayoufutu"
	"fgame/fgame/charge_server/api/paiquyouxixianlu"
	"fgame/fgame/charge_server/api/piaomiao"
	"fgame/fgame/charge_server/api/qia"
	"fgame/fgame/charge_server/api/qiling"
	"fgame/fgame/charge_server/api/qiling2"
	"fgame/fgame/charge_server/api/qiling3"
	"fgame/fgame/charge_server/api/qiyiyouzhangjian"
	"fgame/fgame/charge_server/api/ruozi"
	"fgame/fgame/charge_server/api/shenyu"
	"fgame/fgame/charge_server/api/shunyou"
	"fgame/fgame/charge_server/api/taifeng"
	"fgame/fgame/charge_server/api/taigu"
	"fgame/fgame/charge_server/api/tianji"
	"fgame/fgame/charge_server/api/tianshen"
	"fgame/fgame/charge_server/api/tianshu"
	"fgame/fgame/charge_server/api/tianxing"
	"fgame/fgame/charge_server/api/wanjie"
	"fgame/fgame/charge_server/api/wenjian"
	"fgame/fgame/charge_server/api/xiakexing"
	"fgame/fgame/charge_server/api/xianfan"
	"fgame/fgame/charge_server/api/xiaoyao"
	"fgame/fgame/charge_server/api/xiaoyaojianghu"
	"fgame/fgame/charge_server/api/xinfeng"
	"fgame/fgame/charge_server/api/xingyue"
	"fgame/fgame/charge_server/api/xiongwei"
	"fgame/fgame/charge_server/api/xixiyou"
	"fgame/fgame/charge_server/api/yaojing"
	"fgame/fgame/charge_server/api/yeyushengge"
	"fgame/fgame/charge_server/api/yiyun"
	"fgame/fgame/charge_server/api/youmeng"
	"fgame/fgame/charge_server/api/zhangjian"
	"fgame/fgame/charge_server/api/zhengfu"
	"fgame/fgame/charge_server/api/zhutianxing"
	"fgame/fgame/charge_server/api/zuowan"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	chargePath = "/charge"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(chargePath).Subrouter()
	henggewan.Router(sr)
	zhengfu.Router(sr)
	feifan.Router(sr)
	feiyang.Router(sr)
	shunyou.Router(sr)
	judu.Router(sr)
	mengyuanwenxian.Router(sr)
	luoliwan.Router(sr)
	xinfeng.Router(sr)
	xingyue.Router(sr)
	bocai.Router(sr)
	xixiyou.Router(sr)
	qia.Router(sr)
	zuowan.Router(sr)
	qiling.Router(sr)
	taifeng.Router(sr)
	baji.Router(sr)
	zhutianxing.Router(sr)
	jiuling.Router(sr)
	aotian.Router(sr)
	piaomiao.Router(sr)
	yeyushengge.Router(sr)
	zhangjian.Router(sr)
	tianshen.Router(sr)
	shenyu.Router(sr)
	wenjian.Router(sr)
	xiongwei.Router(sr)
	disui.Router(sr)
	ruozi.Router(sr)
	jiandaowenqing.Router(sr)
	chenxi.Router(sr)
	tianxing.Router(sr)
	tianji.Router(sr)
	yaojing.Router(sr)
	xiakexing.Router(sr)
	tianshu.Router(sr)
	youmeng.Router(sr)
	jiumeng.Router(sr)
	longyu.Router(sr)
	chenxi2.Router(sr)
	xianfan.Router(sr)
	jianghu.Router(sr)
	fengyuan.Router(sr)
	qiling2.Router(sr)
	qiling3.Router(sr)
	menghuan.Router(sr)
	xiaoyao.Router(sr)
	menghuan2.Router(sr)
	xiaoyaojianghu.Router(sr)
	taigu.Router(sr)
	yiyun.Router(sr)
	wanjie.Router(sr)
	lingmeng.Router(sr)
	lieyan.Router(sr)
	jiutian.Router(sr)
	mingjian.Router(sr)
	qiyiyouzhangjian.Router(sr)
	niuchayoufutu.Router(sr)
	mofangyouxifengmo.Router(sr)
	paiquyouxixianlu.Router(sr)
	sr.Path("/get").HandlerFunc(http.HandlerFunc(handleGetOrder))
	// sr.Path("/test").HandlerFunc(http.HandlerFunc(handleTestCharge))
}
