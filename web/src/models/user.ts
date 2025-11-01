export interface User {
  username: string;
  password_hash: string;
  bio?: string | null;
  followers_count?: number | null;
  following_count?: number | null;
  is_admin?: boolean | null;
}
