import './App.css';
import {useState} from "react";
import List from "./Pages/List"
import Stats from "./Pages/Stats"
import Find from "./Pages/Find"

function App() {
    const [page, setPage] = useState("stats");

  return (
    <div className="App">
      <header>
          <ul className="nav justify-content-center">
              <li className="nav-item">
                  <a className="nav-link active" href="#" onClick={e=>setPage("stats")}>Stats</a>
              </li>
              <li className="nav-item">
                  <a className="nav-link" href="#" onClick={e=>setPage("list")}>List</a>
              </li>
              <li className="nav-item">
                  <a className="nav-link" href="#" onClick={e=>setPage("find")}>Find</a>
              </li>
          </ul>
      </header>
      <main>
          {page == "stats" && <Stats/>}
          {page == "list" && <List/>}
          {page == "find" && <Find/>}
      </main>
    </div>
  );
}

export default App;
