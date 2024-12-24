package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func BookmarksRoute(rg *gin.RouterGroup, bookmarksController *controllers.BookmarksController) {
	bookmarks := rg.Group("/bookmarks")
	bookmarks.POST("/:job_id", bookmarksController.AddBookmark)
	bookmarks.DELETE("/:job_id", bookmarksController.RemoveBookmark)
	bookmarks.GET("/", bookmarksController.GetBookmarks)
}
