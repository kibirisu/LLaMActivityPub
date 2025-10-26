import './App.css';
import MainFeed from './components/feed/MainFeed';
import TopAppBar from './components/TopAppBar';

const App = () => {
  const onTopBarSearch = (text) => console.log(text);
  return (
    <div>
      <TopAppBar onSearch={onTopBarSearch} />
      <MainFeed />
    </div>
  );
};

export default App;
