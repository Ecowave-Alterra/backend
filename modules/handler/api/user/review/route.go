package review

import "github.com/labstack/echo/v4"

func (rh *ReviewHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	reviewGroup := e.Group("/user/review")
	// reviewGroup.Use(jwtMiddleware)
	reviewGroup.POST("/:id", rh.CreateReview)
}
