package middleware

import (
	"QtCloudPan/pkg/utils"
	"context"
	"net/http"
	"time"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 从请求头中获取 token
		token := r.Header.Get("Authorization")
		if token == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		// 验证 token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		//  claims中提取过期时间
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token expired")
			return
		}
		// 将用户信息写入上下文
		ctx := context.WithValue(r.Context(), "user", claims)
		next(w, r.WithContext(ctx))
	}
}
