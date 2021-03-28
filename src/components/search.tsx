import * as React from 'react';
const axios = require('axios');

interface SingleResult {
  ns: Number,
  title: String,
  pageid: Number,
  size: Number,
  snippet: String,
  wordcount: Number
}

interface SearchResult {
  data: Array<SingleResult>
}

const Search = () => {
  const [text, changeText] = React.useState('');
  const [results, setResults] = React.useState(Array<SingleResult>());

  const submit = () => {
    axios.get(`/api/search/?t=${text}`)
      .then(( res: SearchResult ) => setResults(res.data))
  }

  const displayResults = () => {
    if (!results.length) {
      return <></>
    }
    return (
      results.map(result => (
        <ul>
          <li>{result.title}</li>
        </ul>
      ))
    )
  }

  return (
    <>
      <input type="text" onChange={e => changeText(e.target.value)} value={text} />
      <button onClick={submit}>Submit</button>
      {displayResults()}
    </>
  )
}

export default Search;