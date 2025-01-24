package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func BookmarksRoute(rg *gin.RouterGroup, bookmarksController *controllers.BookmarksController) {
	bookmarks := rg.Group("/bookmarks")
	bookmarks.POST("/:jobId", bookmarksController.AddBookmark)
	bookmarks.DELETE("/:jobId", bookmarksController.RemoveBookmark)
	bookmarks.GET("/", bookmarksController.GetBookmarks)
}
