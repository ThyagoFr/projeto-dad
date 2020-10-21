import React from 'react';
import Home from './Home';
import Routes from './routes';


function App() {
  return (
    <div className="container">
      <Routes>
        <Home />
      </Routes>
    </div>
  );
}

export default App;
