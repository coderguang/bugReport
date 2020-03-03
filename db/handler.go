package db

import "bugReport/define"

func SaveReportToDb(data *define.SReportData) error {
	return globalDb.Create(data).Error
}
