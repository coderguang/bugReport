package define

import (
	"strconv"
	"time"
)

type SReportData struct {
	Id         int `gorm:"primary_key;AUTO_INCREMENT"`
	AppName    string
	Platform   string
	Token      string
	ReportType int
	ErrCode    int
	Data       string
	ReportDesc string
	ReportDt   time.Time
}

func (data *SReportData) String() string {
	str := "\n=======report start==========\n" +
		"\napp:" + data.AppName +
		"\nplatform:" + data.Platform +
		"\ntoken:" + data.Token +
		"\ntype:" + strconv.Itoa(data.ReportType) +
		"\nerrcode:" + strconv.Itoa(data.ErrCode) +
		"\nreportDt:" + data.ReportDt.String() +
		"\nreportDesc:" + data.ReportDesc +
		"\n===========end===========\n\n"
	return str
}
