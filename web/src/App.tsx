import './App.css';
import { useEffect, useState } from 'react';

const App = () => {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('/api/ping')
      .then(res => res.json())
      .then(data => setMessage(data.message))
      .catch(err => console.error(err));
  }, []);
  
  return (
    <div className="p-8 text-center text-xl">
      <p>Server says: {message || '...'}</p>
    </div>
  );
};

export default App;
