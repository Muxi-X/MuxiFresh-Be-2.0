package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/types"
	"greet/response"
	"net/http"
)

func CheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), svcCtx)
		resp, err := l.Check(&req)
		response.Response(w, resp, err)

	}
}