package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/hamedhaghi/goforum/config"
	"github.com/hamedhaghi/goforum/handlers"
	"github.com/hamedhaghi/goforum/services"
)

func Run(c *config.Config) {
	if c == nil {
		log.Fatal("Config is nil")
		return
	}
	threadService := services.NewThreadService(c.Storage)
	postService := services.NewPostService(c.Storage)
	commentService := services.NewCommentService(c.Storage)
	c.Router.Use(middleware.Logger)
	c.Router.Use(corsConfig())
	c.Router.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	c.Router.Use(middleware.CleanPath)
	c.Router.Use(middleware.GetHead)
	c.Router.Use(middleware.Heartbeat("/"))
	c.Router.Use(middleware.Recoverer)
	c.Router.Use(httprate.LimitByIP(100, 1*time.Minute))

	c.Router.Route("/threads", func(r chi.Router) {
		r.Get("/", handlers.ThreadList(threadService))
		r.Get("/{id}", handlers.ThreadShow(threadService))
		r.Post("/", handlers.ThreadStore(threadService))
		r.Put("/{id}", handlers.ThreadUpdate(threadService))
		r.Delete("/{id}", handlers.ThreadDelete(threadService))
	})
	c.Router.Route("/posts", func(r chi.Router) {
		r.Get("/", handlers.PostList(postService))
		r.Get("/{id}", handlers.PostShow(postService))
		r.Post("/", handlers.PostStore(postService))
		r.Put("/{id}", handlers.PostUpdate(postService))
		r.Delete("/{id}", handlers.PostDelete(postService))
	})
	c.Router.Route("/comments", func(r chi.Router) {
		r.Get("/", handlers.CommentList(commentService))
		r.Get("/{id}", handlers.CommentShow(commentService))
		r.Post("/", handlers.CommentStore(commentService))
		r.Put("/{id}", handlers.CommentUpdate(commentService))
		r.Delete("/{id}", handlers.CommentDelete(commentService))
	})
	err := http.ListenAndServe(
		fmt.Sprintf(
			"%s:%s",
			c.HttpServer,
			c.HttpPort,
		),
		c.Router,
	)
	if err != nil {
		log.Fatalf("error running server: %v \n", err)
	}

	fmt.Printf("Server running at: %s:%s \n", c.HttpServer, c.HttpPort)
}

func corsConfig() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
