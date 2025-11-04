package data

import (
	"context"

	"borg/pkg/db"
)

type PostRepository interface {
	Repository[db.Post, db.AddPostParams, db.UpdatePostParams]
}

type postRepository struct {
	*db.Queries
}

func newPostRepository(q *db.Queries) PostRepository {
	return &postRepository{q}
}

func (r *postRepository) Create(ctx context.Context, post db.AddPostParams) error {
	return r.AddPost(ctx, post)
}

func (r *postRepository) Delete(ctx context.Context, id int32) error {
	return r.DeletePost(ctx, id)
}

func (r *postRepository) GetAll(ctx context.Context) ([]db.Post, error) {
	return r.GetPosts(ctx)
}

func (r *postRepository) GetByID(ctx context.Context, id int32) (db.Post, error) {
	return r.GetPost(ctx, id)
}

func (r *postRepository) Update(ctx context.Context, post db.UpdatePostParams) error {
	return r.UpdatePost(ctx, post)
}
