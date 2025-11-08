import { Column, Entity, Index, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { Comments } from './Comments';
import { Followers } from './Followers';
import { Likes } from './Likes';
import { Posts } from './Posts';
import { Shares } from './Shares';

@Index('users_pkey', ['id'], { unique: true })
@Index('users_username_key', ['username'], { unique: true })
@Entity('users', { schema: 'public' })
export class Users {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('character varying', { name: 'username', unique: true, length: 50 })
  username: string;

  @Column('character varying', { name: 'password_hash', length: 255 })
  passwordHash: string;

  @Column('text', { name: 'bio', nullable: true })
  bio: string | null;

  @Column('integer', {
    name: 'followers_count',
    nullable: true,
    default: () => '0',
  })
  followersCount: number | null;

  @Column('integer', {
    name: 'following_count',
    nullable: true,
    default: () => '0',
  })
  followingCount: number | null;

  @Column('boolean', {
    name: 'is_admin',
    nullable: true,
    default: () => 'false',
  })
  isAdmin: boolean | null;

  @Column('timestamp without time zone', {
    name: 'created_at',
    nullable: true,
    default: () => 'CURRENT_TIMESTAMP',
  })
  createdAt: Date | null;

  @Column('timestamp without time zone', {
    name: 'updated_at',
    nullable: true,
    default: () => 'CURRENT_TIMESTAMP',
  })
  updatedAt: Date | null;

  @OneToMany(
    () => Comments,
    (comments) => comments.user,
  )
  comments: Comments[];

  @OneToMany(
    () => Followers,
    (followers) => followers.follower,
  )
  followers: Followers[];

  @OneToMany(
    () => Followers,
    (followers) => followers.following,
  )
  followers2: Followers[];

  @OneToMany(
    () => Likes,
    (likes) => likes.user,
  )
  likes: Likes[];

  @OneToMany(
    () => Posts,
    (posts) => posts.user,
  )
  posts: Posts[];

  @OneToMany(
    () => Shares,
    (shares) => shares.user,
  )
  shares: Shares[];
}
