package actions

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	c.Session().Set("token", user.AccessToken)

	rTo := c.Session().GetOnce("login_redirect_to")
	rToStr, ok := rTo.(string)
	if !ok || rToStr == "" {
		rToStr = "/"
	}

	// Do something with the user, maybe register them/sign them in
	// return c.Render(200, r.JSON(user))
	return c.Redirect(http.StatusFound, rToStr)
}

func AuthLogout(c buffalo.Context) error {
	c.Session().GetOnce("token")
	return c.Redirect(http.StatusFound, "/")
}

func IsAuth() buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			t := c.Session().Get("token")
			c.Logger().Errorf("token: %s", t)
			if t == nil {
				c.Logger().Errorf("URL: %+v", c.Request().URL.String())
				c.Session().Set("login_redirect_to", c.Request().URL.String())
				return c.Redirect(http.StatusFound, "/auth/github")
				// return c.Redirect() .Error(http.StatusForbidden, errors.New("unauthorized"))
			}
			return h(c)
		}
	}
}
