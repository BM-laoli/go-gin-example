package core_auth

// 存入 token key=user_id value = token
func saveTokenByUserId(uid string, token string) bool {

	return true
}

// 取 key = user_id, 看看这个用户的token 是否还在redis 中
func hasUserToken(uid string) string {
	return ""
}

// 换 key=user_id value = token
func getTokenByUserId(uid string, token string) bool {
	return true
}
