package base

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"wayne/internal/model"
	util "wayne/pkg"
	"wayne/pkg/dto"
)

var (
	PublishRequestMessageMethodFilter = []string{
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}
)

type LoggedInController struct {
	ParamBuilderController
	User *model.User
}

func (c *LoggedInController) Prepare() {
	authString := c.Ctx.Input.Header("Authorization")

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		c.CustomAbort(http.StatusUnauthorized, "Token invalid!")
	}
	tokenString := kv[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return util.RsaPublicKey, nil
	})

	errResult := dto.ErrorResult{}
	switch err.(type) {
	case nil: // no error
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
	c.User, err = model.UserModel.GetUserDetail(aud)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
}
