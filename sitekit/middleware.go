package sitekit

import (
	"context"
	"net/http"

	"github.com/digimakergo/digimaker/core/log/httplog"
	"github.com/digimakergo/digimaker/core/util"
)

func InitRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := util.AnonymousUser()
		ctx = context.WithValue(ctx, util.CtxKeyUserID, userID)

		//start http log
		r = httplog.InitLog(r, ctx, userID)

		next.ServeHTTP(w, r)

		//write http log
		httplog.LogRequest(w, r)
	})
}
