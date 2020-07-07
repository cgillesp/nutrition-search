import React from "react";
import "./App.css";
import SearchBox from "./App/SearchBox";

function App() {
  return (
    <div className="App">
      <div className="header">
        <div className="headerTitle">Nutrition 🔍</div>
      </div>
      <div className="mainView">
        <div className="centerColumn">
          <SearchBox />
          <div className="disclaimer">
            © Charlie Gillespie. This is not medical advice; errors may be
            present in data; use at your own risk.
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
