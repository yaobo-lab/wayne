package oauth2

import (
	"fmt"

	"golang.org/x/net/context"

	"wayne/internal/api/auth"
	"wayne/internal/model"
	pkgoauth "wayne/pkg/oauth2"

	"wayne/pkg/logger"
)

func init() {
	auth.Register(model.AuthTypeOAuth2, &OAuth2Auth{})
}

type OAuth2Auth struct{}

func (*OAuth2Auth) Authenticate(m model.AuthModel) (*model.User, error) {
	oauther := pkgoauth.OAutherMap[m.OAuth2Name]

	code := m.OAuth2Code

	token, err := oauther.Exchange(context.Background(), code)
	if err != nil {
		logger.Errorf("oauth2 get token by code (%s) error.%v", code, err)
		return nil, fmt.Errorf("oauth2 get token by code (%s) error.%v", code, err)
	}
	userinfo, err := oauther.UserInfo(token.AccessToken)
	if err != nil {
		logger.Errorf("oauth2 get user by token (%s) error.%v", token.AccessToken, err)
		return nil, fmt.Errorf("oauth2 get user by token (%s) error.%v", token.AccessToken, err)
	}
	userModel := model.User{
		Name:    userinfo.Name,
		Email:   userinfo.Email,
		Display: userinfo.Display,
	}

	return &userModel, nil
}
