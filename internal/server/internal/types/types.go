// Code generated by goctl. DO NOT EDIT.
package types

type BatchFinancialNewsReq struct {
}

type BatchFinancialNewsResp struct {
	*BaseResp
}

type BatchStockInfoReq struct {
}

type BatchStockInfoResp struct {
	*BaseResp
}

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
