package controller

import (
	"fmt"
	"strings"

	"paper/server-go/config"
	"paper/server-go/model"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func queryString(q map[string]string) string {
	s := []string{}
	for k, v := range q {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, "&")
}

func OauthGoogle(c *gin.Context) {

	base := "https://accounts.google.com/o/oauth2/auth?"
	url := base + queryString(map[string]string{
		"client_id":              config.Oauth_id_google,
		"redirect_uri":           fmt.Sprintf("%v://%v/oauth-callback", config.Protocol, config.Hostname),
		"response_type":          "code",
		"scope":                  "https://www.googleapis.com/auth/plus.me",
		"include_granted_scopes": "true",
	})

	c.Redirect(301, url)
}

func OauthCallback(c *gin.Context) {

	defer func() {
		err := recover()
		if err != nil {
			println("Error: ", err.(string))
			c.String(500, "Server error")
		}
	}()

	authorization_code := c.Query("code")

	getAccesssToken := func() string {
		_, body, errs := gorequest.New().
			Post("https://www.googleapis.com/oauth2/v4/token").
			Set("Content-Type", "application/x-www-form-urlencoded").
			Send(queryString(map[string]string{
				"code":          authorization_code,
				"client_id":     config.Oauth_id_google,
				"client_secret": config.Oauth_secret_google,
				"redirect_uri":  fmt.Sprintf("%v://%v/oauth-callback", config.Protocol, config.Hostname),
				"grant_type":    "authorization_code",
			})).
			End()

		if errs != nil {
			panic(errs[0].Error())
		}

		access_token, err := jsonparser.GetString([]byte(body), "access_token")

		if err != nil {
			panic(err.Error())
		}

		return access_token
	}

	fetchApi := func(access_token string) string {
		_, body, errs := gorequest.New().
			Get("https://www.googleapis.com/oauth2/v2/userinfo").
			Set("Authorization", "Bearer "+access_token).
			End()

		if errs != nil {
			panic(errs[0].Error())
		}

		google_id, err := jsonparser.GetString([]byte(body), "id")

		if err != nil {
			panic(err.Error())
		}
		return google_id

	}

	findOrCreateUser := func(google_id string) model.User {

		user, err := model.FindOrCreateUserByGoogleId(google_id)

		if err != nil {
			panic(err.Error())
		}

		return user
	}

	sendResponse := func(authtoken string) {
		c.String(200, authtoken)
	}

	access_token := getAccesssToken()

	google_id := fetchApi(access_token)

	user := findOrCreateUser(google_id)

	sendResponse(user.Google_ID)

}
