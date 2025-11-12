package router

import (
	"io/fs"
	"log"
	"net/http"

	"borg/internal/domain"
	"borg/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(ds domain.DataStore) http.Handler {
	assets, err := web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/", func(h chi.Router) {
		h.Get("/*", rootHandler(assets))
		h.Get("/static/*", assetHandler(assets))
	})
	r.Mount("/api", apiRouter(ds))

	return r
}

func rootHandler(assets fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, assets, "index.html")
	}
}

func assetHandler(assets fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(assets).ServeHTTP(w, r)
	}
}

func apiRouter(ds domain.DataStore) http.Handler {
	r := chi.NewRouter()
	r.Mount("/user", userRouter(ds))
	r.Mount("/post", postRouter(ds))
	return r
}

func userRouter(ds domain.DataStore) http.Handler {
	repo, opts := ds.UserRepository(), ds.Opts()
	r := chi.NewRouter()
	r.Post("/", create(repo, opts))
	r.Put("/", update(repo, opts))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(idCtx)
		r.Get("/", getByID(repo, opts))
		r.Delete("/", delete(repo))
	})
	r.With(idCtx).Get("/{id}/posts", getByUserID(ds.PostRepository(), opts))
	r.With(idCtx).Get("/{id}/shares", getByUserID(ds.ShareRepository(), opts))
	r.With(idCtx).Get("/{id}/likes", getByUserID(ds.LikeRepository(), opts))
	r.With(idCtx).Get("/{id}/comments", getByUserID(ds.CommentRepository(), opts))
	r.With(idCtx).Get("/{id}/following", getFollowing(repo, opts))
	r.With(idCtx).Get("/{id}/followers", getFollowers(repo, opts))
	return r
}

func postRouter(ds domain.DataStore) http.Handler {
	repo, opts := ds.PostRepository(), ds.Opts()
	r := chi.NewRouter()
	r.Post("/", create(repo, opts))
	r.Put("/", update(repo, opts))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(idCtx)
		r.Get("/", getByID(repo, opts))
		r.Delete("/", delete(repo))
	})
	r.With(idCtx).Get("/{id}/shares", getByPostID(ds.ShareRepository(), opts))
	r.With(idCtx).Get("/{id}/likes", getByPostID(ds.LikeRepository(), opts))
	r.With(idCtx).Get("/{id}/comments", getByPostID(ds.CommentRepository(), opts))
	// TODO consider using different route
	r.Post("/comments", create(ds.CommentRepository(), opts))
	r.Delete("/comments", delete(ds.CommentRepository()))
	return r
}
