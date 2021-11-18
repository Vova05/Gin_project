package middlewares

//import (
//	"github.com/gin-gonic/gin"
//)
//

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() (gin.HandlerFunc){
	return gin.BasicAuth(gin.Accounts{
		"pragmatic": "reviews",
	})
}
func BasicAuth2() (gin.HandlerFunc){
	return gin.BasicAuth(gin.Accounts{
		"pragmatic2": "reviews",
	})
}
