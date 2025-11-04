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

