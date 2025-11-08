import { Column, Entity, Index, JoinColumn, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';
import { Users } from './Users';

@Index('followers_follower_id_following_id_key', ['followerId', 'followingId'], { unique: true })
@Index('followers_pkey', ['id'], { unique: true })
@Entity('followers', { schema: 'public' })
export class Followers {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('integer', { name: 'follower_id', unique: true })
  followerId: number;

  @Column('integer', { name: 'following_id', unique: true })
  followingId: number;

  @Column('timestamp without time zone', {
    name: 'created_at',
    nullable: true,
    default: () => 'CURRENT_TIMESTAMP',
  })
  createdAt: Date | null;

  @ManyToOne(
    () => Users,
    (users) => users.followers,
  )
  @JoinColumn([{ name: 'follower_id', referencedColumnName: 'id' }])
  follower: Users;

  @ManyToOne(
    () => Users,
    (users) => users.followers2,
  )
  @JoinColumn([{ name: 'following_id', referencedColumnName: 'id' }])
  following: Users;
}
