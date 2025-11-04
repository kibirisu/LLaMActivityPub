import { Sql } from "postgres";

export const getUsersQuery = `-- name: GetUsers :many
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at FROM users`;

export interface GetUsersRow {
    id: number;
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getUsers(sql: Sql): Promise<GetUsersRow[]> {
    return (await sql.unsafe(getUsersQuery, []).values()).map(row => ({
        id: row[0],
        username: row[1],
        passwordHash: row[2],
        bio: row[3],
        followersCount: row[4],
        followingCount: row[5],
        isAdmin: row[6],
        createdAt: row[7],
        updatedAt: row[8]
    }));
}

export const addUserQuery = `-- name: AddUser :exec
INSERT INTO users (
  username,
  password_hash,
  bio,
  followers_count,
  following_count,
  is_admin
) VALUES ($1, $2, $3, $4, $5, $6)`;

export interface AddUserArgs {
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
}

export async function addUser(sql: Sql, args: AddUserArgs): Promise<void> {
    await sql.unsafe(addUserQuery, [args.username, args.passwordHash, args.bio, args.followersCount, args.followingCount, args.isAdmin]);
}

export const getUserQuery = `-- name: GetUser :one
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at FROM users WHERE id = $1`;

export interface GetUserArgs {
    id: number;
}

export interface GetUserRow {
    id: number;
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getUser(sql: Sql, args: GetUserArgs): Promise<GetUserRow | null> {
    const rows = await sql.unsafe(getUserQuery, [args.id]).values();
    if (rows.length !== 1) {
        return null;
    }
    const row = rows[0];
    return {
        id: row[0],
        username: row[1],
        passwordHash: row[2],
        bio: row[3],
        followersCount: row[4],
        followingCount: row[5],
        isAdmin: row[6],
        createdAt: row[7],
        updatedAt: row[8]
    };
}

export const updateUserQuery = `-- name: UpdateUser :exec
UPDATE users SET password_hash = $2, bio = $3, followers_count = $4, following_count = $5, is_admin = $6 WHERE id = $1`;

export interface UpdateUserArgs {
    id: number;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
}

export async function updateUser(sql: Sql, args: UpdateUserArgs): Promise<void> {
    await sql.unsafe(updateUserQuery, [args.id, args.passwordHash, args.bio, args.followersCount, args.followingCount, args.isAdmin]);
}

export const deleteUserQuery = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1`;

export interface DeleteUserArgs {
    id: number;
}

export async function deleteUser(sql: Sql, args: DeleteUserArgs): Promise<void> {
    await sql.unsafe(deleteUserQuery, [args.id]);
}

export const getPostsQuery = `-- name: GetPosts :many
SELECT id, user_id, content, like_count, share_count, comment_count, created_at, updated_at FROM posts`;

export interface GetPostsRow {
    id: number;
    userId: number;
    content: string;
    likeCount: number | null;
    shareCount: number | null;
    commentCount: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getPosts(sql: Sql): Promise<GetPostsRow[]> {
    return (await sql.unsafe(getPostsQuery, []).values()).map(row => ({
        id: row[0],
        userId: row[1],
        content: row[2],
        likeCount: row[3],
        shareCount: row[4],
        commentCount: row[5],
        createdAt: row[6],
        updatedAt: row[7]
    }));
}

export const addPostQuery = `-- name: AddPost :exec
INSERT INTO posts (user_id, content) VALUES ($1, $2)`;

export interface AddPostArgs {
    userId: number;
    content: string;
}

export async function addPost(sql: Sql, args: AddPostArgs): Promise<void> {
    await sql.unsafe(addPostQuery, [args.userId, args.content]);
}

export const getPostQuery = `-- name: GetPost :one
SELECT id, user_id, content, like_count, share_count, comment_count, created_at, updated_at FROM posts WHERE id = $1`;

export interface GetPostArgs {
    id: number;
}

export interface GetPostRow {
    id: number;
    userId: number;
    content: string;
    likeCount: number | null;
    shareCount: number | null;
    commentCount: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getPost(sql: Sql, args: GetPostArgs): Promise<GetPostRow | null> {
    const rows = await sql.unsafe(getPostQuery, [args.id]).values();
    if (rows.length !== 1) {
        return null;
    }
    const row = rows[0];
    return {
        id: row[0],
        userId: row[1],
        content: row[2],
        likeCount: row[3],
        shareCount: row[4],
        commentCount: row[5],
        createdAt: row[6],
        updatedAt: row[7]
    };
}

export const updatePostQuery = `-- name: UpdatePost :exec
UPDATE posts SET content = $2, like_count = $3, share_count = $4, comment_count = $5 WHERE id = $1`;

export interface UpdatePostArgs {
    id: number;
    content: string;
    likeCount: number | null;
    shareCount: number | null;
    commentCount: number | null;
}

export async function updatePost(sql: Sql, args: UpdatePostArgs): Promise<void> {
    await sql.unsafe(updatePostQuery, [args.id, args.content, args.likeCount, args.shareCount, args.commentCount]);
}

export const deletePostQuery = `-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1`;

export interface DeletePostArgs {
    id: number;
}

export async function deletePost(sql: Sql, args: DeletePostArgs): Promise<void> {
    await sql.unsafe(deletePostQuery, [args.id]);
}

