import './App.css';
import MainFeed from './feed/MainFeed';
import TopAppBar from './TopAppBar';

const App = () => {
  const onTopBarSearch = (text) => console.log(text)
  return (
    <div>
      <TopAppBar onSearch={onTopBarSearch}/>
      <MainFeed />
    </div>
  );
};

export default App;
