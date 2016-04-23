package main
import ("github.com/gin-gonic/gin"
	"github.com/alivinco/green-route/middleware"
	"net/http"
	"github.com/alivinco/green-route/datasource"
	"flag"
	"log"
)


func main() {

	ds_host := "http://lego.fiicha.net:8080"
	flag.StringVar(&ds_host,"ds_host","http://lego.fiicha.net:8080","Datasource host")
	flag.Parse()
	log.Println("Darasource host = %s",ds_host)
	router := gin.Default()
	router.Static("/greenr/static","./static")
	router.LoadHTMLGlob("templates/**/*")
	aMid := middleware.NewAuthMiddleware(router)

	appRoot := router.Group("/greenr/ui/")
	appRoot.Use(aMid.RequestHandler())
	{
		appRoot.GET("/profile",func(c *gin.Context) {
			c.Get("UserData")
			user,_:=c.Get("UserData")
        		c.HTML(http.StatusOK, "profile.html", user)
		})
		appRoot.GET("/start",func(c *gin.Context) {
        		c.HTML(http.StatusOK, "start.html", gin.H{})
		})
	}

	appRootApi := router.Group("/greenr/api/")
	{
		appRootApi.GET("/start",)
		appRootApi.GET("/proxy",func(c *gin.Context) {
			var resp []byte
			var err error
			if c.Query("ds")=="datanorge"{
				resp ,err = datasource.GetDataNorge(ds_host,c.Query("source"),c.Query("longitude"),c.Query("latitude"),c.Query("radius"),c.Query("date"),c.Query("school"))
			}

        		if err != nil {
				c.String(http.StatusInternalServerError,"%s",err.Error())

			}else {
				c.Data(http.StatusOK, "text/json", resp)
			}

		})
	}


	router.Run(":7000")
}
