import React, { useState, useEffect } from "react";
import { Heart, Repeat, MessageCircle, Share2 } from "lucide-react";
import { Post } from "./feedData";
import { timeAgo, initials } from "./utils";
import ReactMarkdown from 'react-markdown'


type Props = {
  post: Post;
};


export default function PostItem({ post }: Props) {

  return (
    <div className="border-b border-gray-200 p-4 hover:bg-gray-50 transition-colors">
      <div className="flex space-x-3">
        <div className="flex-1">
          <div className="flex items-center space-x-1">
            <span className="font-semibold text-gray-900">
              {post.author.name}
            </span>
            <span className="text-gray-500 text-sm">
              {post.author.handle} · {timeAgo(post.createdAt)}
            </span>
          </div>
          <div className="prose max-w-none text-gray-800">
            <ReactMarkdown>{post.content}</ReactMarkdown>
          </div>

          <div className="flex justify-between mt-3 text-gray-500 text-sm max-w-md">
            <button className="flex items-center space-x-1 hover:text-blue-500 transition">
              <MessageCircle size={16} /> <span>{post.replies}</span>
            </button>
            <button className="flex items-center space-x-1 hover:text-green-500 transition">
              <Repeat size={16} /> <span>{post.reposts}</span>
            </button>
            <button className="flex items-center space-x-1 hover:text-pink-500 transition">
              <Heart size={16} /> <span>{post.likes}</span>
            </button>
            <button className="flex items-center space-x-1 hover:text-gray-700 transition">
              <Share2 size={16} />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

