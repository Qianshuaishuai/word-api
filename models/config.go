package models

import (
	"context"
	"net/http"
	"os"
	"regexp"

	loglib "github.com/HYY-yu/LogLib"
	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	opentracing "github.com/opentracing/opentracing-go"
)

//Models包结构
// m0config.go : 程序入口，用于配置初始化。
// m*data.go : 资源结构定义文件，以及存放一些不公开的方法。

type Mconfig struct {
	LogLevel int

	LogFile string

	dBHost     string
	dBName     string
	dBUsername string
	dBPassword string
	dBMaxIdle  int
	dBMaxConn  int

	SnowFlakDomain           string
	SnowFlakAuthUser         string
	SnowFlakAuthUserSecurity string

	LoginServerDomain         string
	PrepareDomain             string
	PrepareApiAccesstoken     string
	ApiToken                  string
	CircuitDynamicCoverDomain string

	// 默认Response消息
	ConfigMyResponse map[int]string

	UseElasticSearch bool
	ElasticUrl       string
	ElasticUsername  string
	ElasticPassword  string
	ElasticIndexName string

	WxSessionAPIURL string
	WxPayAPIURL     string
	WxAppID         string
	WxAppSecret     string
	WxGrantType     string

	// 是否在Docker容器中，影响是否把日志打到标准输出或者文件。
	INDOCKER string

	//错误处理
	RecoverPanic bool

	ImgAudioResourceDomain string
	BeikeDomain            string

	TraceOpen       int
	TraceRemoteAddr string
	TraceServerName string
}

var (
	MyConfig     Mconfig
	dbOrmDefault *gorm.DB

	Regexp_dbConfig_Sql                = regexp.MustCompile(`\?`)
	Regexp_dbConfig_NumericPlaceHolder = regexp.MustCompile(`\$\d+`)
)

const (
	timeLayoutStr = "2006-01-02 15:04:05"
)

const (
	//公共响应码
	RESP_OK        = 10000 // 200
	RESP_ERR       = 10001 // 400
	RESP_PARAM_ERR = 10002 // 200
	RESP_TOKEN_ERR = 10003 // 403
	RESP_NO_ACCESS = 10004 // 400
	RESP_NO_RESULT = 10005 //

	//自定义响应码 --200
	RESP_RESOURCE_NOT_FOUND       = 13300
	RESP_RESOURCE_API_TOKEN_ERROR = 13301
)

//AccessToken 相关
const (
	ROLE_STUDENT = 1
	ROLE_TEACHER = 2

	PLATFORM_ANDROID = 1
	PLATFORM_WEB     = 2
	PLATFORM_WEBCHAT = 3
)

//电子书包资源定义
const (
	TYPE_POETRY_DETAIl = -59
)

func init() {
	DREAMENV := os.Getenv("DREAMENV")
	if len(DREAMENV) <= 0 {
		DREAMENV = "PROD"
	}
	appName := "word-api"

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return
	}

	MyConfig = Mconfig{}
	if appConf != nil {
		MyConfig.INDOCKER = os.Getenv("INDOCKER")

		levelStr := appConf.String(appName + "::LogLevel")
		switch levelStr {
		case "DEBUG":
			MyConfig.LogLevel = loglib.LevelDebug
		case "INFO":
			MyConfig.LogLevel = loglib.LevelInfo
		case "ERROR":
			MyConfig.LogLevel = loglib.LevelError
		}

		MyConfig.LogFile = appConf.String(DREAMENV + "::LogFile")
		MyConfig.dBHost = appConf.String(DREAMENV + "::dbHost")
		MyConfig.dBName = appConf.String(DREAMENV + "::dbName")
		MyConfig.dBUsername = appConf.String(DREAMENV + "::dBUsername")
		MyConfig.dBPassword = appConf.String(DREAMENV + "::dbPassword")
		MyConfig.dBMaxIdle, _ = appConf.Int(DREAMENV + "::dbMaxIdle")
		MyConfig.dBMaxConn, _ = appConf.Int(DREAMENV + "::dbMaxConn")
		MyConfig.ApiToken = appConf.String(DREAMENV + "::apiToken")
		MyConfig.TraceOpen, _ = appConf.Int(DREAMENV + "::traceOpen")
		MyConfig.TraceServerName = appConf.String(DREAMENV + "::traceServerName")
		MyConfig.RecoverPanic = appConf.DefaultBool(DREAMENV+"::recoverPanic", true)

		MyConfig.SnowFlakDomain = appConf.String(DREAMENV + "::domain")
		MyConfig.SnowFlakAuthUser = appConf.String(DREAMENV + "::authUser")
		MyConfig.SnowFlakAuthUserSecurity = appConf.String(DREAMENV + "::authUserSecurity")

		MyConfig.WxSessionAPIURL = appConf.String(DREAMENV + "::wxSessionApiUrl")
		MyConfig.WxAppID = appConf.String(DREAMENV + "::wxAppId")
		MyConfig.WxAppSecret = appConf.String(DREAMENV + "::wxAppSecret")
		MyConfig.WxGrantType = appConf.String(DREAMENV + "::wxGrantType")

	}
	getResponseConfig()

	initLog()
}

func initLog() {
	//初始化日志模块
	if Indocker() {
		loglib.InitLogger(loglib.LogConfig{LogTo: loglib.ConsoleLogs, LogLevel: MyConfig.LogLevel, LogPretty: false})
	} else {
		loglib.InitLogger(loglib.LogConfig{LogTo: loglib.FileLogs, LogPath: MyConfig.LogFile, LogLevel: MyConfig.LogLevel, LogPretty: true})
	}
}

//获取config
func getResponseConfig() {
	MyConfig.ConfigMyResponse = make(map[int]string)
	MyConfig.ConfigMyResponse[RESP_OK] = "成功"
	MyConfig.ConfigMyResponse[RESP_ERR] = "失败,未知错误"
	MyConfig.ConfigMyResponse[RESP_PARAM_ERR] = "参数错误"
	MyConfig.ConfigMyResponse[RESP_TOKEN_ERR] = "token错误"
	MyConfig.ConfigMyResponse[RESP_NO_ACCESS] = "没有访问权限"
}

//获取对应的db对象
func GetDb() *gorm.DB {
	return dbOrmDefault
}

func Indocker() bool {
	return len(MyConfig.INDOCKER) > 0
}

//获取appctx里的traceSpan
func GetTraceSpan(appCtx context.Context) (TraceSpan opentracing.Span) {
	TraceSpan = nil
	if MyConfig.TraceOpen == 1 && appCtx != nil {
		if traceSpanTmp := appCtx.Value("TraceSpan"); traceSpanTmp != nil {
			traceSpan2, ok := traceSpanTmp.(opentracing.Span)
			if ok && traceSpan2 != nil {
				TraceSpan = traceSpan2
			}
		}
	}
	return
}

//向请求投注入TraceSpan
func InjectTraceSpanToReqHeader(appCtx context.Context, reqHeader http.Header) {
	TraceSpan := GetTraceSpan(appCtx)
	if TraceSpan != nil {
		TraceSpan.Tracer().Inject(TraceSpan.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(reqHeader))
	}
}
