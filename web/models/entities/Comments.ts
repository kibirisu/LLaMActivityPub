import { Column, Entity, Index, JoinColumn, ManyToOne, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { Posts } from './Posts';
import { Users } from './Users';

@Index('comments_pkey', ['id'], { unique: true })
@Entity('comments', { schema: 'public' })
export class Comments {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('text', { name: 'content' })
  content: string;

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

  @ManyToOne(
    () => Comments,
    (comments) => comments.comments,
  )
  @JoinColumn([{ name: 'parent_id', referencedColumnName: 'id' }])
  parent: Comments;

  @OneToMany(
    () => Comments,
    (comments) => comments.parent,
  )
  comments: Comments[];

  @ManyToOne(
    () => Posts,
    (posts) => posts.comments,
  )
  @JoinColumn([{ name: 'post_id', referencedColumnName: 'id' }])
  post: Posts;

  @ManyToOne(
    () => Users,
    (users) => users.comments,
  )
  @JoinColumn([{ name: 'user_id', referencedColumnName: 'id' }])
  user: Users;
}
