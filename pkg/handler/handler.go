package handler 

import(
    "github.com/depressed-kid/todo-app/pkg/service"
    "github.com/gin-gonic/gin"
) 



type Handler struct {
    services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{
	services: services,
    }
}




func (this *Handler) InitRoutes() *gin.Engine {
    router := gin.New()

    auth := router.Group("/auth") 
    {
	auth.POST("/sign-up", this.signUp)
	auth.POST("/sign-in", this.signIn)
    }

    api := router.Group("/api") 
    {
	lists := api.Group("/lists") 
	{
	    lists.POST("/", this.createList)
	    lists.GET("/", this.getAllLists)
	    lists.GET("/:id", this.getListById)
	    lists.PUT("/:id", this.updateList)
	    lists.DELETE("/:id", this.deleteList)

	    items := lists.Group(":id/items") 
	    {
		items.POST("/", this.createItem)
		items.GET("/", this.getAllItems)
		items.GET("/:item_id", this.getItemById)
		items.PUT("/:item_id", this.updateItem)
		items.DELETE("/:item_id", this.deleteItem)
	    }

	}



    }

    return router

}




