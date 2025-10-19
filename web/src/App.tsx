import './App.css';
import { useEffect, useState } from 'react';

const App = () => {
  const [message, setMessage] = useState('');
  const [dbmessage, setDBMessage] = useState('');

  useEffect(() => {
    fetch('/api/ping')
      .then(res => res.json())
      .then(data => setMessage(data.message))
      .catch(err => console.error(err));
    fetch('/api/users')
      .then(res => res.json())
      .then(data => setDBMessage(data.users))
      .catch(err => console.error(err));
  }, []);
  
  return (
    <div className="p-8 text-center text-xl">
      <p>Server says: {message || '...'}</p>
      <p>Database says: {JSON.stringify(dbmessage) || '...'}</p>
    </div>
  );
};

export default App;
