package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func InitializeRoutes(r *gin.RouterGroup) {
	db["hoge"] = "unko"
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		userName := c.Params.ByName("name")

		if v, ok := db[userName]; ok {
			c.JSON(http.StatusOK, gin.H{"value": v})
		} else {
			c.JSON(http.StatusOK, gin.H{"value": "not found"})
		}
	})

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":      "bar",
		"autin":    "1234",
		"sawasaki": "narumi",
	}))

	authorized.POST("", func(c *gin.Context) {
		// contextからuserキーの値を取得する
		// この値はBasicAuthが返すmiddlewareがsetする
		userName := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[userName] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "failed"})
		}
	})
}
