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
    content: "# This\n is a comment, guys.",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(),
    likes: 12,
    replies: 3,
    reposts: 1,
  },
  {
    id: "2",
    author: { name: "Grace Hopper", handle: "@grace", avatarColor: "bg-green-500" },
    content: "### This is testing markdown\n`still testing markdown`\nThis is the realest comment that exists.",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 45,
    replies: 8,
    reposts: 10,
  },
  {
    id: "3",
    author: { name: "John Greene", handle: "@jgrn", avatarColor: "bg-red-500" },
    content: "Hi\n I am a real person that is testing links\n[link to my page](/mypage).",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 69,
    replies: 2,
    reposts: 1,
  },
  {
    id: "4",
    author: { name: "Mark Zuck", handle: "@mzuck", avatarColor: "bg-blue-500" },
    content: "Hi  \nI on the other hand am testing lists:  \n* a\n* b\n* c\n* last item",
    createdAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
    likes: 45,
    replies: 8,
    reposts: 10,
  },
];

