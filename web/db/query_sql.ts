import { Sql } from "postgres";

export const addUserQuery = `-- name: AddUser :exec
INSERT INTO users (
  username,
  password_hash,
  bio,
  followers_count,
  following_count,
  is_admin,
  origin
) VALUES ($1, $2, $3, $4, $5, $6, $7)`;

export interface AddUserArgs {
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    origin: string | null;
}

export async function addUser(sql: Sql, args: AddUserArgs): Promise<void> {
    await sql.unsafe(addUserQuery, [args.username, args.passwordHash, args.bio, args.followersCount, args.followingCount, args.isAdmin, args.origin]);
}

export const getUserQuery = `-- name: GetUser :one
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at, origin FROM users WHERE id = $1`;

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
    origin: string | null;
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
        updatedAt: row[8],
        origin: row[9]
    };
}

export const getFollowedUsersQuery = `-- name: GetFollowedUsers :many
SELECT u.id, u.username, u.password_hash, u.bio, u.followers_count, u.following_count, u.is_admin, u.created_at, u.updated_at, u.origin FROM users u JOIN followers f ON u.id = f.following_id WHERE f.follower_id = $1`;

export interface GetFollowedUsersArgs {
    followerId: number;
}

export interface GetFollowedUsersRow {
    id: number;
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    createdAt: Date | null;
    updatedAt: Date | null;
    origin: string | null;
}

export async function getFollowedUsers(sql: Sql, args: GetFollowedUsersArgs): Promise<GetFollowedUsersRow[]> {
    return (await sql.unsafe(getFollowedUsersQuery, [args.followerId]).values()).map(row => ({
        id: row[0],
        username: row[1],
        passwordHash: row[2],
        bio: row[3],
        followersCount: row[4],
        followingCount: row[5],
        isAdmin: row[6],
        createdAt: row[7],
        updatedAt: row[8],
        origin: row[9]
    }));
}

export const getFollowingUsersQuery = `-- name: GetFollowingUsers :many
SELECT u.id, u.username, u.password_hash, u.bio, u.followers_count, u.following_count, u.is_admin, u.created_at, u.updated_at, u.origin FROM users u JOIN followers f ON u.id = f.follower_id WHERE f.following_id = $1`;

export interface GetFollowingUsersArgs {
    followingId: number;
}

export interface GetFollowingUsersRow {
    id: number;
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    createdAt: Date | null;
    updatedAt: Date | null;
    origin: string | null;
}

export async function getFollowingUsers(sql: Sql, args: GetFollowingUsersArgs): Promise<GetFollowingUsersRow[]> {
    return (await sql.unsafe(getFollowingUsersQuery, [args.followingId]).values()).map(row => ({
        id: row[0],
        username: row[1],
        passwordHash: row[2],
        bio: row[3],
        followersCount: row[4],
        followingCount: row[5],
        isAdmin: row[6],
        createdAt: row[7],
        updatedAt: row[8],
        origin: row[9]
    }));
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

export const addPostQuery = `-- name: AddPost :exec
INSERT INTO posts (user_id, content) VALUES ($1, $2)`;

export interface AddPostArgs {
    userId: number;
    content: string;
}

export async function addPost(sql: Sql, args: AddPostArgs): Promise<void> {
    await sql.unsafe(addPostQuery, [args.userId, args.content]);
}

export const getPostsByUserIDQuery = `-- name: GetPostsByUserID :many
SELECT id, user_id, content, like_count, share_count, comment_count, created_at, updated_at FROM posts WHERE user_id = $1`;

export interface GetPostsByUserIDArgs {
    userId: number;
}

export interface GetPostsByUserIDRow {
    id: number;
    userId: number;
    content: string;
    likeCount: number | null;
    shareCount: number | null;
    commentCount: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getPostsByUserID(sql: Sql, args: GetPostsByUserIDArgs): Promise<GetPostsByUserIDRow[]> {
    return (await sql.unsafe(getPostsByUserIDQuery, [args.userId]).values()).map(row => ({
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

export const addCommentQuery = `-- name: AddComment :exec
INSERT INTO comments (post_id, user_id, content, parent_id) VALUES ($1, $2, $3, $4)`;

export interface AddCommentArgs {
    postId: number;
    userId: number;
    content: string;
    parentId: number | null;
}

export async function addComment(sql: Sql, args: AddCommentArgs): Promise<void> {
    await sql.unsafe(addCommentQuery, [args.postId, args.userId, args.content, args.parentId]);
}

export const getPostCommentsQuery = `-- name: GetPostComments :many
SELECT c.id, c.post_id, c.user_id, c.content, c.parent_id, c.created_at, c.updated_at FROM comments c JOIN posts p ON c.post_id = p.id WHERE p.id = $1`;

export interface GetPostCommentsArgs {
    id: number;
}

export interface GetPostCommentsRow {
    id: number;
    postId: number;
    userId: number;
    content: string;
    parentId: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getPostComments(sql: Sql, args: GetPostCommentsArgs): Promise<GetPostCommentsRow[]> {
    return (await sql.unsafe(getPostCommentsQuery, [args.id]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        content: row[3],
        parentId: row[4],
        createdAt: row[5],
        updatedAt: row[6]
    }));
}

export const getUserCommentsQuery = `-- name: GetUserComments :many
SELECT c.id, c.post_id, c.user_id, c.content, c.parent_id, c.created_at, c.updated_at FROM comments c JOIN users u ON c.user_id = u.id WHERE u.id = $1`;

export interface GetUserCommentsArgs {
    id: number;
}

export interface GetUserCommentsRow {
    id: number;
    postId: number;
    userId: number;
    content: string;
    parentId: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getUserComments(sql: Sql, args: GetUserCommentsArgs): Promise<GetUserCommentsRow[]> {
    return (await sql.unsafe(getUserCommentsQuery, [args.id]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        content: row[3],
        parentId: row[4],
        createdAt: row[5],
        updatedAt: row[6]
    }));
}

export const deleteCommentQuery = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1`;

export interface DeleteCommentArgs {
    id: number;
}

export async function deleteComment(sql: Sql, args: DeleteCommentArgs): Promise<void> {
    await sql.unsafe(deleteCommentQuery, [args.id]);
}

export const addLikeQuery = `-- name: AddLike :exec
INSERT INTO likes (post_id, user_id) VALUES ($1, $2)`;

export interface AddLikeArgs {
    postId: number;
    userId: number;
}

export async function addLike(sql: Sql, args: AddLikeArgs): Promise<void> {
    await sql.unsafe(addLikeQuery, [args.postId, args.userId]);
}

export const getLikeByIDQuery = `-- name: GetLikeByID :one
SELECT id, post_id, user_id, created_at FROM likes WHERE id = $1`;

export interface GetLikeByIDArgs {
    id: number;
}

export interface GetLikeByIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getLikeByID(sql: Sql, args: GetLikeByIDArgs): Promise<GetLikeByIDRow | null> {
    const rows = await sql.unsafe(getLikeByIDQuery, [args.id]).values();
    if (rows.length !== 1) {
        return null;
    }
    const row = rows[0];
    return {
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    };
}

export const getLikesByPostIDQuery = `-- name: GetLikesByPostID :many
SELECT id, post_id, user_id, created_at FROM likes WHERE post_id = $1`;

export interface GetLikesByPostIDArgs {
    postId: number;
}

export interface GetLikesByPostIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getLikesByPostID(sql: Sql, args: GetLikesByPostIDArgs): Promise<GetLikesByPostIDRow[]> {
    return (await sql.unsafe(getLikesByPostIDQuery, [args.postId]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    }));
}

export const getLikesByUserIDQuery = `-- name: GetLikesByUserID :many
SELECT id, post_id, user_id, created_at FROM likes WHERE user_id = $1`;

export interface GetLikesByUserIDArgs {
    userId: number;
}

export interface GetLikesByUserIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getLikesByUserID(sql: Sql, args: GetLikesByUserIDArgs): Promise<GetLikesByUserIDRow[]> {
    return (await sql.unsafe(getLikesByUserIDQuery, [args.userId]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    }));
}

export const deleteLikeQuery = `-- name: DeleteLike :exec
DELETE FROM likes WHERE id = $1`;

export interface DeleteLikeArgs {
    id: number;
}

export async function deleteLike(sql: Sql, args: DeleteLikeArgs): Promise<void> {
    await sql.unsafe(deleteLikeQuery, [args.id]);
}

export const addShareQuery = `-- name: AddShare :exec
INSERT INTO shares (post_id, user_id) VALUES ($1, $2)`;

export interface AddShareArgs {
    postId: number;
    userId: number;
}

export async function addShare(sql: Sql, args: AddShareArgs): Promise<void> {
    await sql.unsafe(addShareQuery, [args.postId, args.userId]);
}

export const getShareByIDQuery = `-- name: GetShareByID :one
SELECT id, post_id, user_id, created_at FROM shares WHERE id = $1`;

export interface GetShareByIDArgs {
    id: number;
}

export interface GetShareByIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getShareByID(sql: Sql, args: GetShareByIDArgs): Promise<GetShareByIDRow | null> {
    const rows = await sql.unsafe(getShareByIDQuery, [args.id]).values();
    if (rows.length !== 1) {
        return null;
    }
    const row = rows[0];
    return {
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    };
}

export const getSharesByPostIDQuery = `-- name: GetSharesByPostID :many
SELECT id, post_id, user_id, created_at FROM shares WHERE post_id = $1`;

export interface GetSharesByPostIDArgs {
    postId: number;
}

export interface GetSharesByPostIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getSharesByPostID(sql: Sql, args: GetSharesByPostIDArgs): Promise<GetSharesByPostIDRow[]> {
    return (await sql.unsafe(getSharesByPostIDQuery, [args.postId]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    }));
}

export const getShareByUserIDQuery = `-- name: GetShareByUserID :many
SELECT id, post_id, user_id, created_at FROM shares WHERE user_id = $1`;

export interface GetShareByUserIDArgs {
    userId: number;
}

export interface GetShareByUserIDRow {
    id: number;
    postId: number;
    userId: number;
    createdAt: Date | null;
}

export async function getShareByUserID(sql: Sql, args: GetShareByUserIDArgs): Promise<GetShareByUserIDRow[]> {
    return (await sql.unsafe(getShareByUserIDQuery, [args.userId]).values()).map(row => ({
        id: row[0],
        postId: row[1],
        userId: row[2],
        createdAt: row[3]
    }));
}

export const deleteShareQuery = `-- name: DeleteShare :exec
DELETE FROM shares WHERE id = $1`;

export interface DeleteShareArgs {
    id: number;
}

export async function deleteShare(sql: Sql, args: DeleteShareArgs): Promise<void> {
    await sql.unsafe(deleteShareQuery, [args.id]);
}

export const getPostsByOriginQuery = `-- name: GetPostsByOrigin :many
SELECT p.id, p.user_id, p.content, p.like_count, p.share_count, p.comment_count, p.created_at, p.updated_at FROM posts p JOIN users u ON p.user_id = u.id WHERE u.origin = $1`;

export interface GetPostsByOriginArgs {
    origin: string | null;
}

export interface GetPostsByOriginRow {
    id: number;
    userId: number;
    content: string;
    likeCount: number | null;
    shareCount: number | null;
    commentCount: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export async function getPostsByOrigin(sql: Sql, args: GetPostsByOriginArgs): Promise<GetPostsByOriginRow[]> {
    return (await sql.unsafe(getPostsByOriginQuery, [args.origin]).values()).map(row => ({
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

export const getAllUsersQuery = `-- name: GetAllUsers :many
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at, origin FROM users`;

export interface GetAllUsersRow {
    id: number;
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
    createdAt: Date | null;
    updatedAt: Date | null;
    origin: string | null;
}

export async function getAllUsers(sql: Sql): Promise<GetAllUsersRow[]> {
    return (await sql.unsafe(getAllUsersQuery, []).values()).map(row => ({
        id: row[0],
        username: row[1],
        passwordHash: row[2],
        bio: row[3],
        followersCount: row[4],
        followingCount: row[5],
        isAdmin: row[6],
        createdAt: row[7],
        updatedAt: row[8],
        origin: row[9]
    }));
}

