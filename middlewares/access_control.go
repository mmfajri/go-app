package middlewares

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

//Authorize object access based on current user that has been authorized to take an action (based on casbin rule)
func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		sub, existed := c.Get("userID")
		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"msg" :"User hasn't logged in yet"})
			return
		}
		 
		//Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to load policy from DB"})
			return
		}
		
		//Casbin enforces policy
		ok, err := enforcer.Enforce(fmt.Sprint(sub), obj, act)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Error occured when authorizing user"})
			return	
		}

		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"msg": "The user not authorized to this endpoint"})
			return
		}
		c.Next()
	}
}
