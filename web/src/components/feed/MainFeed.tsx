import { useState } from 'react';
import { type Post, samplePosts } from './feedData';
import NewPostBox from './NewPostBox';
import PostItem from './PostItem';
import type { User } from '../../models/user';

export default function MainFeed() {
  const [posts, setPosts] = useState<Post[]>(samplePosts);

  function addPost(content: string) {
    const newPost: Post = {
      id: String(Date.now()),
      author: { name: 'You', handle: '@you', avatarColor: 'bg-indigo-600' },
      content: content.trim(),
      createdAt: new Date().toISOString(),
      likes: 0,
      replies: 0,
      reposts: 0,
    };
    setPosts((p) => [newPost, ...p]);
  }

  const user: User = {
      id: 0,
      username: 'zbyszek',
      password_hash: 'p@ssword'
  }

  const [val, setVal] = useState("");

  async function doRequest() {
    let payload = JSON.stringify(user);
    console.log(payload)
    fetch("/api/", {
      method: "POST",
      body: payload,
      headers: {
        "Content-Type": "application/json"
      }
    }).then((resp) => {
      console.log(resp)
    });
    setVal("done");
  }

  return (
    <div className="max-w-2xl mx-auto border-x border-gray-300 min-h-screen bg-white">
      <header className="p-4 border-b border-gray-300 text-xl font-bold sticky top-0 bg-white/80 backdrop-blur z-10">
        Home
      </header>
      <NewPostBox onPost={addPost} />
      <p>{val}</p>
      <button className="btn" onClick={doRequest}>Default</button>
      {posts.map((post) => (
        <PostItem key={post.id} post={post} />
      ))}
    </div>
  );
}
