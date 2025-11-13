import { useParams } from 'react-router';
import { type Post, samplePosts } from '../feed/feedData';
import PostItem from '../feed/PostItem';

export default function UserProfile() {
  const params = useParams();
  const handleParam = params.handle ? `@${params.handle}` : undefined;

  const foundPost = samplePosts.find((p: Post) => p.author.handle === handleParam);
  const user = foundPost?.author ?? {
    name: 'Test User',
    handle: handleParam ?? '@test',
    avatarColor: 'bg-indigo-600',
  };

  const userPosts = samplePosts.filter((p: Post) => p.author.handle === user.handle);

  return (
    <div className="max-w-2xl mx-auto border-x border-gray-300 min-h-screen bg-white">
      <header className="p-4 border-b border-gray-300 text-xl font-bold sticky top-0 bg-white/80 backdrop-blur z-10 text-black">
        Profile
      </header>
      <div className="p-6">
        <div className="flex flex-col items-center space-y-4">
          <div className={`w-16 h-16 rounded-full ${user.avatarColor ?? 'bg-gray-300'}`}></div>
          <div className="text-center">
            <div className="text-2xl font-semibold text-gray-900">{user.name}</div>
            <div className="text-gray-600">{user.handle}</div>
          </div>
        </div>
      </div>
      <div className="border-t border-gray-300" />
      {userPosts.map((post) => (
        <PostItem key={post.id} post={post} />
      ))}
    </div>
  );
}
