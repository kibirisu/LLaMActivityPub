import { UserX } from 'lucide-react';
import { useParams } from 'react-router';
import { type Post, samplePosts, sampleUsers, type User } from '../feed/feedData';
import PostItem from '../feed/PostItem';

export default function UserProfile() {
  const params = useParams();
  const handleParam = params.handle ? `@${params.handle}` : undefined;
  //here we would also get user bio and posts from the backend
  const user = sampleUsers.find((u: User) => u.username === handleParam);

  if (user === undefined) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen text-center p-6">
        <div className="flex flex-col items-center gap-4 max-w-md">
          <div className="bg-red-100 text-red-600 p-4 rounded-full">
            <UserX className="w-10 h-10" />
          </div>
          <h1 className="text-3xl font-bold">Sorry, that user does not exist</h1>
          <p className="text-gray-500">
            The user you’re looking for might have changed their username, deleted their account, or never existed at
            all.
          </p>
          <a
            href="/"
            className="mt-4 inline-block px-4 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 transition-colors"
          >
            Go back home
          </a>
        </div>
      </div>
    );
  }

  const userPosts = samplePosts.filter((p: Post) => p.author === user.username);

  return (
    <div className="max-w-2xl mx-auto border-x border-gray-300 min-h-screen bg-white">
      <header className="p-4 border-b border-gray-300 text-xl font-bold sticky top-0 bg-white/80 backdrop-blur z-10 text-black">
        Profile
      </header>
      <div className="p-6">
        <div className="flex flex-col items-center space-y-4">
          <div className={`w-16 h-16 rounded-full ${user.avatarColor ?? 'bg-gray-300'}`}></div>
          <div className="text-center">
            <div className="text-gray-600">{user.username}</div>
            <div className="text-2xl font-semibold text-gray-900">{user.bio}</div>
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
