package auth2


type AuthProvider interface {
    GetName()  string
    GetToken(code string) (string, error)
    GetUserInfo(token string)  (UserInfo, error)
    GetCallBackUrl(base string)  (string, error)
}
