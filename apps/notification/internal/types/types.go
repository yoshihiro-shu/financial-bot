// Code generated by goctl. DO NOT EDIT.
package types

type NotificationFinancialNewsReq struct {
}

type NotificationFinancialNewsResp struct {
	*BaseResp
}

type NotificationStockInfoReq struct {
}

type NotificationStockInfoResp struct {
	*BaseResp
}

type NotificationCryptoCurrencyInfoReq struct {
}

type NotificationCryptoCurrencyInfoResp struct {
	*BaseResp
}

type HealtCheckReq struct {
}

type HealtCheckResp struct {
	Code int `json:"code"`
}

type BaseReq struct {
}

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
