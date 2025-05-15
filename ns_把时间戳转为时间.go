package timezh

import "time"

var NS = &ns把时间戳转为时间{}

type ns把时间戳转为时间 struct{}

func (T *ns把时间戳转为时间) Get纳秒时间(ns纳秒数 int64) time.Time {
	return time.Unix(ns纳秒数/1e9, ns纳秒数%1e9)
}

func (T *ns把时间戳转为时间) Get微秒时间(ms微秒数 int64) time.Time {
	return time.UnixMicro(ms微秒数)
}

func (T *ns把时间戳转为时间) Get毫秒时间(ms微秒数 int64) time.Time {
	return time.UnixMilli(ms微秒数)
}

func (T *ns把时间戳转为时间) Get秒数时间(sn总秒数 int64) time.Time {
	return time.Unix(sn总秒数, 0)
}
