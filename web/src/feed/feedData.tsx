export type Post = {
  id: string;
  author: {
    name: string;
    handle: string;
    avatarColor?: string;
  };
  content: string;
  createdAt: string;
  likes?: number;
  replies?: number;
  reposts?: number;
};


export const samplePosts: Post[] = [
  {
    id: "1",
    author: { name: "Ada Lovelace", handle: "@ada", avatarColor: "bg-purple-500" },
    content: "This is a comment, guys.",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(),
    likes: 12,
    replies: 3,
    reposts: 1,
  },
  {
    id: "2",
    author: { name: "Grace Hopper", handle: "@grace", avatarColor: "bg-green-500" },
    content: "This is the realest comment that exists.",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 45,
    replies: 8,
    reposts: 10,
  },
];
