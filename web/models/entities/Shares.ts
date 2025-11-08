import { Column, Entity, Index, JoinColumn, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';
import { Posts } from './Posts';
import { Users } from './Users';

@Index('shares_pkey', ['id'], { unique: true })
@Index('shares_post_id_user_id_key', ['postId', 'userId'], { unique: true })
@Entity('shares', { schema: 'public' })
export class Shares {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('integer', { name: 'post_id', unique: true })
  postId: number;

  @Column('integer', { name: 'user_id', unique: true })
  userId: number;

  @Column('timestamp without time zone', {
    name: 'created_at',
    nullable: true,
    default: () => 'CURRENT_TIMESTAMP',
  })
  createdAt: Date | null;

  @ManyToOne(
    () => Posts,
    (posts) => posts.shares,
  )
  @JoinColumn([{ name: 'post_id', referencedColumnName: 'id' }])
  post: Posts;

  @ManyToOne(
    () => Users,
    (users) => users.shares,
  )
  @JoinColumn([{ name: 'user_id', referencedColumnName: 'id' }])
  user: Users;
}
