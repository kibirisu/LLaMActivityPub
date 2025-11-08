import { Column, Entity, Index, PrimaryGeneratedColumn } from 'typeorm';

@Index('goose_db_version_pkey', ['id'], { unique: true })
@Entity('goose_db_version', { schema: 'public' })
export class GooseDbVersion {
  @PrimaryGeneratedColumn({ type: 'integer', name: 'id' })
  id: number;

  @Column('bigint', { name: 'version_id' })
  versionId: string;

  @Column('boolean', { name: 'is_applied' })
  isApplied: boolean;

  @Column('timestamp without time zone', {
    name: 'tstamp',
    default: () => 'now()',
  })
  tstamp: Date;
}
