package route

import (
	"core/middleware"

	"github.com/labstack/echo"
)

func v1Routes(g *echo.Group, h AppModel) {
	g.GET("/health", h.Health.Check)

	auth := g.Group("/auth")
	auth.POST("/register", h.Auth.RegisterUser)
	auth.POST("/resend-otp", h.Auth.ResendOTP)
	auth.POST("/verify-otp", h.Auth.VerifyOTP)
	auth.POST("/login", h.Auth.LoginUser)
	auth.GET("/validate", h.Auth.ValidateSession, middleware.JWTVerify())
	auth.GET("/logout", h.Auth.UserLogOut, middleware.JWTVerify())
	auth.GET("/github/callback", h.Auth.GithubOAuthCallback, middleware.JWTVerify())
	auth.GET("/google", h.Auth.GoogleAuthURL)
	auth.GET("/google/callback", h.Auth.GoogleOAuthCallback)
	auth.GET("/github", h.Auth.GithubAuthURL)
	auth.GET("/github/callback", h.Auth.GithubAuthCallback)

	user := g.Group("/user", middleware.JWTVerify())
	user.GET("/get-user", h.User.GetUserName)
	user.POST("/update-profile", h.User.UpdateUserProfile)

	userList := g.Group("/user-list", middleware.JWTVerify())
	userList.GET("/get", h.UserList.GetUserList)
	userList.POST("/create", h.UserList.CreateUserList)
	userList.GET("/get_expenses", h.UserList.GetUserListExpenses)

	category := g.Group("/category", middleware.JWTVerify())
	category.POST("/create", h.Category.CreateCategory)
	category.GET("/list", h.Category.GetCategories)

	expense := g.Group("/expense", middleware.JWTVerify())
	expense.POST("", h.Expense.CreateExpense)
	expense.GET("/list", h.Expense.RecentExpenses)
}
