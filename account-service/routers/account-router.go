package routers

import (
	"account-service/middleware"
	"account-service/service"

	"github.com/gin-gonic/gin"
)

func AccountRouter(r *gin.RouterGroup, accountService *service.AccountService) {
	r.GET("/accounts", accountService.GetAccounts)
	r.GET("/fetch-acc/:acc_num", accountService.GetAccoutById)

	// Adding new User
	r.POST("/register", accountService.AddAccount)

	// Update user details
	r.PATCH("/update_acc/:acc_num", accountService.UpdateUser)

	// Delete User
	r.DELETE("/delete-user/:email",
		middleware.FetchUserUUID(),
		middleware.DeleteAuthUser(),
		accountService.DeleteUser,
	)
}
