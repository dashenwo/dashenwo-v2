package sessions

import (
	"context"
	"github.com/dashenwo/dashenwo/v2/console/account/global"
	"github.com/dashenwo/go-library/session"
	"github.com/micro/go-micro/v2/server"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/http/httptest"
)

type sessionKey struct{}

func NewSessionWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			r, _ := http.NewRequestWithContext(ctx, "", "", nil)
			// 把metadata转换为http的header
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				header := http.Header{}
				for key, value := range md {
					header.Add(key, value[0])
				}
				r.Header = header
			}
			w := httptest.NewRecorder()
			s := global.SessionManage.New(
				session.Request(r),
				session.Writer(w),
				session.Scheme("grpc"),
			)
			ctx = context.WithValue(ctx, sessionKey{}, s)
			return h(ctx, req, rsp)
		}
	}
}

func GetSession(ctx context.Context) *session.Session {
	return ctx.Value(sessionKey{}).(*session.Session)
}
