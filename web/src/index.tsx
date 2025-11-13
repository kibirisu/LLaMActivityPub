import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router';
import App from './App';
import CommentView from './components/feed/CommentView';
import UserProfile from './components/profile/UserProfile';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    path: '/profile/:handle',
    element: <UserProfile />,
  },
  {
    path: '/foo',
    element: <h1>bar</h1>,
  },
  {
    path: '/post/:postId',
    element: <CommentView />,
  },
]);

const rootEl = document.getElementById('root');
if (rootEl) {
  const root = ReactDOM.createRoot(rootEl);
  root.render(
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>,
  );
}
