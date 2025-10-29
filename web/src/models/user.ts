export interface User {
  id: number;
  username: string;
  password_hash: string;
  bio?: string | null;
  followersCount?: number | null;
  followingCount?: number | null;
  isAdmin?: boolean | null;
  createdAt?: string | null;
  updatedAt?: string | null;
}
