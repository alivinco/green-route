package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/apexskier/httpauth"
	//"github.com/gorilla/mux"
	"strings"
	"fmt"
	"log"
)

type AuthMiddleware struct {
	router *gin.Engine
	backend httpauth.AuthBackend
	aaa *httpauth.Authorizer

}

func NewAuthMiddleware(router *gin.Engine)(*AuthMiddleware) {

	backend, err := httpauth.NewGobFileAuthBackend ("auth.db")
	if err != nil {
		panic(err)
	}
	auth := AuthMiddleware{router:router,backend:&backend}
	auth.SetDefaults()
	auth.SetRoutes()

	return &auth
}

func (a *AuthMiddleware) SetDefaults(){
	// create some default roles
	var err error
	roles := make(map[string]httpauth.Role)
	roles["user"] = 30
	roles["admin"] = 80
	aaa, err := httpauth.NewAuthorizer(a.backend, []byte("cookie-encryption-key"), "user", roles)
	a.aaa = &aaa
	// create a default user
	username := "admin"
	defaultUser := httpauth.UserData{Username: username, Role: "admin"}
	err = a.backend.SaveUser(defaultUser)
	if err != nil {
		panic(err)
	}
	// Update user with a password and email address
	err = a.aaa.Update(nil, nil, username, "adminadmin", "admin@localhost.com")
	if err != nil {
		panic(err)
	}
}

func (a *AuthMiddleware) SetRoutes(){
	a.router.GET("/greenr/ui/login",func(c *gin.Context) {
        	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
	})
	a.router.POST("/greenr/ui/login",a.postLogin)
	a.router.GET("/greenr/ui/user/change",func(c *gin.Context) {
        	c.HTML(http.StatusOK, "change.html", gin.H{"title": "Login"})
	})
	a.router.Any("/greenr/ui/logout",a.handleLogout)

}

func (a *AuthMiddleware) RequestHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		rw := c.Writer
		req := c.Request
		//log.Println("New request")
		//c.Header("Pragma","no-cache")
		//c.Header("Cache-Control","no-cache")
		if err := a.aaa.Authorize(rw, req, true); err != nil {
			log.Println(err)
			http.Redirect(rw, req, "/greenr/ui/login", http.StatusSeeOther)
			return
		}
		if user, err := a.aaa.CurrentUser(rw, req); err == nil {
			c.Set("UserData",user)
		}

	}
}



func (a *AuthMiddleware) postLogin(c *gin.Context) {
	rw := c.Writer
	req := c.Request
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	if err := a.aaa.Login(rw, req, username, password, "/greenr/ui/start"); err == nil || (err != nil && strings.Contains(err.Error(), "already authenticated")) {
		http.Redirect(rw, req, "/greenr/ui/start", http.StatusSeeOther)
	} else if err != nil {
		fmt.Println(err)
		http.Redirect(rw, req, "/greenr/ui/login", http.StatusSeeOther)
	}
}

func (a *AuthMiddleware) handleLogout(c *gin.Context) {
	rw := c.Writer
	req := c.Request
	if err := a.aaa.Logout(rw, req); err != nil {
		fmt.Println(err)
		// this shouldn't happen
		return
	}
	http.Redirect(rw, req, "/greenr/ui/login", http.StatusSeeOther)

}