export type User = {
  id: number;
  username: string;
  bio?: string;
  followers_count: number;
  following_count: number;
  is_admin: boolean;
  created_at: string;
  updated_at: string;
  avatarColor?: string; // not in DB !!!
};
export type Post = {
  id: string;
  author: string;
  content: string;
  createdAt: string;
  likes?: number;
  replies?: number;
  reposts?: number;
};

export const samplePosts: Post[] = [
  {
    id: '1',
    author: '@ada',
    content: '# This\n is a comment, guys.',
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(),
    likes: 12,
    replies: 3,
    reposts: 1,
  },
  {
    id: '2',
    author: '@grace',
    content: '### This is testing markdown\n`still testing markdown`\nThis is the realest comment that exists.',
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 45,
    replies: 8,
    reposts: 10,
  },
  {
    id: '3',
    author: '@jgrn',
    content: 'Hi\n I am a real person that is testing links\n[link to my page](/mypage).',
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 69,
    replies: 2,
    reposts: 1,
  },
  {
    id: '4',
    author: '@mzuck',
    content: 'Hi  \nI on the other hand am testing lists:  \n* a\n* b\n* c\n* last item',
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 45,
    replies: 8,
    reposts: 10,
  },
  {
    id: '5',
    author: '@ada',
    content: '# This\n is another post from ada.',
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(),
    likes: 0,
    replies: 0,
    reposts: 0,
  },
];
export const sampleUsers: User[] = [
  {
    id: 1,
    username: '@ada',
    bio: 'Ada Lovelace — the first programmer.',
    followers_count: 120,
    following_count: 80,
    is_admin: false,
    created_at: '2025-10-12T09:15:32.000Z',
    updated_at: '2025-11-01T11:22:10.000Z',
    avatarColor: 'bg-purple-500',
  },
  {
    id: 2,
    username: '@grace',
    bio: 'Grace Hopper — pioneer of compilers.',
    followers_count: 200,
    following_count: 150,
    is_admin: false,
    created_at: '2025-10-11T08:42:18.000Z',
    updated_at: '2025-11-02T13:10:05.000Z',
    avatarColor: 'bg-green-500',
  },
  {
    id: 3,
    username: '@jgrn',
    bio: 'John Greene — likes writing and testing code.',
    followers_count: 75,
    following_count: 50,
    is_admin: false,
    created_at: '2025-10-15T10:55:44.000Z',
    updated_at: '2025-11-03T14:18:21.000Z',
    avatarColor: 'bg-red-500',
  },
  {
    id: 4,
    username: '@mzuck',
    bio: 'Mark Zuck — testing out lists and social stuff.',
    followers_count: 5000,
    following_count: 300,
    is_admin: true,
    created_at: '2025-09-20T12:00:00.000Z',
    updated_at: '2025-11-01T12:00:00.000Z',
    avatarColor: 'bg-blue-500',
  },
];
