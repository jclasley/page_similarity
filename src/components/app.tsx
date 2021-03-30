import axios from 'axios';
import * as React from 'react'
import Search from './search'

const App = () => {
  const [page1, setPage1] = React.useState([0, '']);
  const [page2, setPage2] = React.useState([0, '']);
  const [similarity, setSimilarity] = React.useState(0)

  const handleClick = () => {
    const p1 = axios.get(`/api/extract/?id=${page1[0]}`).then(({ data }) => data);
    const p2 = axios.get(`/api/extract/?id=${page2[0]}`).then(({ data }) => data);
    Promise.all([p1, p2]).then(([t1, t2]) => axios.get(`/api/compare/?t1=${t1}&t2=${t2}`))
      .then(({ data }) => setSimilarity(data * 100));
  }

  return (
    <div className="app">
      <div id="search-one" className="subcontainer">
        {page1[0] ? page1[1] : "Make a selection"}
        <Search click={setPage1} />
      </div>
      <div id="search-two" className="subcontainer">
        {page2[0] ? page2[1] : "Make a selection"}
        <Search click={setPage2} />
      </div>
      <div className="result subcontainer">
        {similarity ? `${similarity.toFixed(2)}% similar` : "Waiting for input..."}
        <button onClick={handleClick}>Submit</button>
      </div>
    </div>
  )
}

export default App;