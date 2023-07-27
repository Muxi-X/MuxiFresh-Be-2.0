package handler

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/logic"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"
	"MuXiFresh-Be-2.0/common/greet/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AuthPassChangeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthPassChangeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAuthPassChangeLogic(r.Context(), svcCtx)
		resp, err := l.AuthPassChange(&req)
		response.Response(w, resp, err)

	}
}
