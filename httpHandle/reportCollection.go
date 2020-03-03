package httpHandle

import (
	"bugReport/db"
	"bugReport/define"
	"net/http"
	"strconv"

	"github.com/coderguang/GameEngine_go/sglog"
)

func logicHandle(w http.ResponseWriter, r *http.Request, flag chan bool) {
	r.ParseForm()

	defer func() {
		flag <- true
	}()

	data := new(define.SReportData)
	data.AppName = r.FormValue(define.HTTP_ARGS_APP)
	data.Platform = r.FormValue(define.HTTP_ARGS_PLATFORM)
	data.Token = r.FormValue(define.HTTP_ARGS_TOKEN)
	reportType, err := strconv.Atoi(r.FormValue(define.HTTP_ARGS_REPORT_TYPE))
	if err != nil {
		sglog.Error("transform report type error,err:", err)
	} else {
		data.ReportType = reportType
	}
	errcode, err := strconv.Atoi(r.FormValue(define.HTTP_ARGS_ERROR_CODE))
	if err != nil {
		sglog.Error("transform errcode error,err:", err)
	} else {
		data.ErrCode = errcode
	}
	data.Data = r.FormValue(define.HTTP_ARGS_ERROR_DATA)
	data.ReportDesc = r.FormValue(define.HTTP_ARGS_ERROR_DESC)

	err = db.SaveReportToDb(data)
	if err != nil {
		sglog.Error("save report to db error")
	}
	sglog.Info("get new report:", data)
}
