package main

import (
	"Ainotes/config"
	"Ainotes/controller"
	"Ainotes/middleware"
	"Ainotes/repository"
	"Ainotes/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	resourceRepo := repository.NewResourceRepository(db)
	noteRepo := repository.NewNoteRepository(db)
	quizRepo := repository.NewQuizRepository(db)
	flashcardRepo := repository.NewFlashcardRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	resourceService := service.NewResourceService(resourceRepo)
	noteService := service.NewNoteService(noteRepo)
	quizService := service.NewQuizService(quizRepo)
	flashcardService := service.NewFlashcardService(flashcardRepo)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	resourceController := controller.NewResourceController(resourceService)
	noteController := controller.NewNoteController(noteService)
	quizController := controller.NewQuizController(quizService)
	flashcardController := controller.NewFlashcardController(flashcardService)

	// Set up router
	r := gin.Default()

	// Apply Supabase Auth middleware to all /api routes
	api := r.Group("/api", middleware.SupabaseAuthMiddleware())
	{
		users := api.Group("/users")

		users.GET("/", userController.FindAll)
		users.GET("/:id", userController.FindByID)
		users.POST("/", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)

		resources := api.Group("/resources")

		resources.GET("/", resourceController.FindAll)
		resources.GET("/:id", resourceController.FindByID)
		resources.POST("/", resourceController.Create)
		resources.PUT("/:id", resourceController.Update)
		resources.DELETE("/:id", resourceController.Delete)

		notes := api.Group("/notes")

		notes.GET("/", noteController.FindAll)
		notes.GET("/:id", noteController.FindByID)
		notes.POST("/", noteController.Create)
		notes.PUT("/:id", noteController.Update)
		notes.DELETE("/:id", noteController.Delete)

		quizzes := api.Group("/quizzes")

		quizzes.GET("/", quizController.FindAll)
		quizzes.GET("/:id", quizController.FindByID)
		quizzes.POST("/", quizController.Create)
		quizzes.PUT("/:id", quizController.Update)
		quizzes.DELETE("/:id", quizController.Delete)

		flashcards := api.Group("/flashcards")

		flashcards.GET("/", flashcardController.FindAll)
		flashcards.GET("/:id", flashcardController.FindByID)
		flashcards.POST("/", flashcardController.Create)
		flashcards.PUT("/:id", flashcardController.Update)
		flashcards.DELETE("/:id", flashcardController.Delete)
	}

	// Run server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
