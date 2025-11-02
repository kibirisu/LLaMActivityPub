import { Sql } from "postgres";

export const getUsersQueryQuery = `-- name: GetUsersQuery :many
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at FROM users`;

export interface GetUsersQueryRow {
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

export async function getUsersQuery(sql: Sql): Promise<GetUsersQueryRow[]> {
    return (await sql.unsafe(getUsersQueryQuery, []).values()).map(row => ({
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

export const addUserQueryQuery = `-- name: AddUserQuery :exec
INSERT INTO users (
  username,
  password_hash,
  bio,
  followers_count,
  following_count,
  is_admin
) VALUES ($1, $2, $3, $4, $5, $6)`;

export interface AddUserQueryArgs {
    username: string;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
}

export async function addUserQuery(sql: Sql, args: AddUserQueryArgs): Promise<void> {
    await sql.unsafe(addUserQueryQuery, [args.username, args.passwordHash, args.bio, args.followersCount, args.followingCount, args.isAdmin]);
}

export const getUserQueryQuery = `-- name: GetUserQuery :one
SELECT id, username, password_hash, bio, followers_count, following_count, is_admin, created_at, updated_at FROM users WHERE id = $1`;

export interface GetUserQueryArgs {
    id: number;
}

export interface GetUserQueryRow {
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

export async function getUserQuery(sql: Sql, args: GetUserQueryArgs): Promise<GetUserQueryRow | null> {
    const rows = await sql.unsafe(getUserQueryQuery, [args.id]).values();
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

export const updateUserQueryQuery = `-- name: UpdateUserQuery :exec
UPDATE users SET password_hash = $2, bio = $3, followers_count = $4, following_count = $5, is_admin = $6 WHERE id = $1`;

export interface UpdateUserQueryArgs {
    id: number;
    passwordHash: string;
    bio: string | null;
    followersCount: number | null;
    followingCount: number | null;
    isAdmin: boolean | null;
}

export async function updateUserQuery(sql: Sql, args: UpdateUserQueryArgs): Promise<void> {
    await sql.unsafe(updateUserQueryQuery, [args.id, args.passwordHash, args.bio, args.followersCount, args.followingCount, args.isAdmin]);
}

export const deleteUserQueryQuery = `-- name: DeleteUserQuery :exec
DELETE FROM users WHERE id = $1`;

export interface DeleteUserQueryArgs {
    id: number;
}

export async function deleteUserQuery(sql: Sql, args: DeleteUserQueryArgs): Promise<void> {
    await sql.unsafe(deleteUserQueryQuery, [args.id]);
}

