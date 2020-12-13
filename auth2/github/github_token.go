package github

import (
    "fmt"
    "github.com/slowmoon/base/auth2"
)

const (
    redirectUrl  = "https://github.com/login/oauth/authorize"
    accessToken  = "https://github.com/login/oauth/access_token"
    apiUser      = "https://api.github.com/user"
)

type githubToken struct {
    clientId   string
    secret     string
    callbackUrl string
}


type  options  func(token *githubToken)

func WithClientId(id  string) options {
    return func(token *githubToken) {
        token.clientId =  id
    }
}
func WithSecret(secret string) options {
    return func(token *githubToken) {
        token.secret = secret
    }
}
func WithCallbackUrl(callback string) options {
    return func(token *githubToken) {
        token.callbackUrl = callback
    }
}

func NewGithub(opts ...options) (*githubToken, error){
    var g githubToken
    for _, opt := range opts {
        opt(&g)
    }
    if g.clientId == "" || g.secret == "" || g.callbackUrl == "" {
        panic("missing required params")
    }
    return  &g, nil
}

func (g *githubToken)GetToken(code string) (string, error) {
    return "", nil
}

func (g *githubToken)GetUserInfo(token string)  (auth2.UserInfo, error) {
    return auth2.UserInfo{}, nil
}

func (g *githubToken) GetCallBackUrl(base string) (string, error) {
    return  fmt.Sprintf("base?client_id=%s&redirect_uri=%s", g.clientId, g.callbackUrl), nil
}

func (g *githubToken) GetName() string{
    return  "github"
}
