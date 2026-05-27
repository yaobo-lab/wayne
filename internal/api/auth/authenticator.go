package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"

	"encoding/json"
	"wayne/internal/api/base"
	"wayne/internal/model"
	rsakey "wayne/pkg"
	"wayne/pkg/dto"
	"wayne/pkg/hack"
	selfoauth "wayne/pkg/oauth2"

	"wayne/pkg/logger"
)

// Authenticator provides interface to authenticate user credentials.
type Authenticator interface {
	// Authenticate ...
	Authenticate(m model.AuthModel) (*model.User, error)
}

var registry = make(map[string]Authenticator)

// Register add different authenticators to registry map.
func Register(name string, authenticator Authenticator) {
	if _, dup := registry[name]; dup {
		logger.Infof("authenticator: %s has been registered", name)
		return
	}
	registry[name] = authenticator
}

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

type LoginResult struct {
	Token string `json:"token"`
}

// swagger:route GET /login/{type}/{name} auth reqLoginAuth
// type is login type <br/>
// name when login type is oauth2 used for oauth2 type
// responses:
//   200: respSuccessDescription

// swagger:route POST /login/{type}/{name} auth reqLoginAuth
// type is login type <br/>
// name when login type is oauth2 used for oauth2 type
// responses:
//
//	200: respSuccessDescription
func (c *AuthController) Login() {
	var authModel model.AuthModel
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &authModel)

	authType := c.Ctx.Input.Param(":type")
	oauth2Name := c.Ctx.Input.Param(":name")
	next := c.Ctx.Input.Query("next")
	if authType == "" || authModel.Username == "admin" {
		authType = model.AuthTypeDB
	}

	authenticator, ok := registry[authType]
	if !ok {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Ctx.Output.Body(hack.Slice(fmt.Sprintf("auth type (%s) is not supported.", authType)))
		return
	}

	if authType == model.AuthTypeOAuth2 {
		oauther, ok := selfoauth.OAutherMap[oauth2Name]
		if !ok {
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Ctx.Output.Body(hack.Slice("oauth2 type is not supported."))
			return
		}
		code := c.Input().Get("code")
		if code == "" {
			c.Ctx.Redirect(http.StatusFound, oauther.AuthCodeURL(next, oauth2.AccessTypeOnline))
			return
		}
		authModel.OAuth2Code = code
		authModel.OAuth2Name = oauth2Name
		state := c.Ctx.Input.Query("state")
		if state != "" {
			next = state
		}

	}

	user, err := authenticator.Authenticate(authModel)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Ctx.Output.Body(hack.Slice(fmt.Sprintf("Login failed. %v", err)))
		return
	}

	now := time.Now()
	user.LastIp = c.Ctx.Input.IP()
	user.LastLogin = &now
	user, err = model.UserModel.EnsureUser(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body(hack.Slice(err.Error()))
		return
	}

	// default token exp time is 3600s.
	expSecond := beego.AppConfig.DefaultInt64("TokenLifeTime", 86400)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		// 签发者
		"iss": "wayne",
		// 签发时间
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(expSecond) * time.Second).Unix(),
		"aud": user.Name,
	})

	apiToken, err := token.SignedString(rsakey.RsaPrivateKey)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body(hack.Slice(err.Error()))
		return
	}

	if next != "" {
		// if oauth type is oauth, set token for client.
		if authType == model.AuthTypeOAuth2 {
			next = next + "&sid=" + apiToken
		}
		c.Redirect(next, http.StatusFound)
		return
	}

	loginResult := LoginResult{
		Token: apiToken,
	}
	c.Data["json"] = base.Result{Data: loginResult}
	c.ServeJSON()
}

// swagger:route GET /logout auth reqLogoutAuth
//
// logout
func (c *AuthController) Logout() {

}

// swagger:route GET /currentuser auth reqCurrentUserAuth
// get current user
// responses:
//
//	200: respSuccessDescription
func (c *AuthController) CurrentUser() {
	c.Controller.Prepare()
	authString := c.Ctx.Input.Header("Authorization")

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		c.CustomAbort(http.StatusUnauthorized, "Token Invalid ! ")
	}
	tokenString := kv[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return rsakey.RsaPublicKey, nil
	})

	errResult := dto.ErrorResult{}
	switch err.(type) {
	case nil:
		if !token.Valid { // but may still be invalid
			errResult.Code = http.StatusUnauthorized
			errResult.Msg = "Token Invalid ! "
		}

	case *jwt.ValidationError: // something was wrong during the validation
		errResult.Code = http.StatusUnauthorized
		errResult.Msg = err.Error()

	default: // something else went wrong
		errResult.Code = http.StatusInternalServerError
		errResult.Msg = err.Error()
	}

	if err != nil {
		c.CustomAbort(errResult.Code, errResult.Msg)
	}

	claim := token.Claims.(jwt.MapClaims)
	aud := claim["aud"].(string)
	user, err := model.UserModel.GetUserDetail(aud)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	c.Data["json"] = base.Result{Data: user}
	c.ServeJSON()
}
