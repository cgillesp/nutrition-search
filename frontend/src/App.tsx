import React from "react";
import "./App.css";
import SearchBox from "./App/SearchBox";

function App() {
  return (
    <div className="App">
      <div className="header">
        <div className="headerTitle">
          Nutrition{" "}
          <span role="img" aria-label="Magnifying Glass (Search Icon)">
            🔍
          </span>
        </div>
      </div>
      <div className="mainView">
        <div className="centerColumn">
          <SearchBox />
          <div className="disclaimer">
            © Charlie Gillespie. This is not medical advice; errors may be
            present in data; use at your own risk. Icons by{" "}
            <a href="https://openmoji.org">OpenMoji</a> under{" "}
            <a href="https://creativecommons.org/licenses/by-sa/4.0/">
              CC BY-SA 4.0
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
