import { Column, Entity, Index, JoinColumn, ManyToOne, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { Comments } from './Comments';
import { Likes } from './Likes';
import { Users } from './Users';
import { Shares } from './Shares';

@Index('posts_pkey', ['id'], { unique: true })
@Entity('posts', { schema: 'public' })
export class Posts {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('text', { name: 'content' })
  content: string;

  @Column('integer', { name: 'like_count', nullable: true, default: () => '0' })
  likeCount: number | null;

  @Column('integer', {
    name: 'share_count',
    nullable: true,
    default: () => '0',
  })
  shareCount: number | null;

  @Column('integer', {
    name: 'comment_count',
    nullable: true,
    default: () => '0',
  })
  commentCount: number | null;

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
    (comments) => comments.post,
  )
  comments: Comments[];

  @OneToMany(
    () => Likes,
    (likes) => likes.post,
  )
  likes: Likes[];

  @ManyToOne(
    () => Users,
    (users) => users.posts,
  )
  @JoinColumn([{ name: 'user_id', referencedColumnName: 'id' }])
  user: Users;

  @OneToMany(
    () => Shares,
    (shares) => shares.post,
  )
  shares: Shares[];
}
