// Code generated by goctl. DO NOT EDIT.
package types

type ChoiceItem struct {
	Number int64  `json:"number"`
	Data   string `json:"data"`
}

type TestReq struct {
	Authorization string       `header:"Authorization"`
	Choice        []ChoiceItem `json:"choice"`
}

type TestResp struct {
	Flag bool `json:"flag"`
}

type TestInfoReq struct {
	Authorization string `header:"Authorization"`
	UserID        string `form:"user_id"`
}

type TestInfoResp struct {
	Name        string       `json:"name"`
	Gender      string       `json:"gender"`
	Major       string       `json:"major"`
	Grade       string       `json:"grade"`
	LeQunXing   int64        `json:"le_qun_xing"`
	YouHengXing int64        `json:"you_heng_xing"`
	XingFenXing int64        `json:"xing_fen_fen_xing"`
	CongHuiXing int64        `json:"cong_hui_xing"`
	JiaoJiXing  int64        `json:"jiao_ji_xing"`
	HuaiYiXing  int64        `json:"huai_yi_xing"`
	WenDingXing int64        `json:"wen_ding_xing"`
	Choice      []ChoiceItem `json:"choice"`
}
