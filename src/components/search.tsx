import * as React from 'react';
import Result from './result';
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

type SearchProps = {click: Function}

const Search = ( { click }: SearchProps) => {
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
      <div className="results">
        {results.map(result => (
          <Result pageid={result.pageid} click={click} title={result.title} />
        ))}
      </div>
    )
  }

  return (
    <>
      <input type="text" onChange={e => changeText(e.target.value)} value={text} />
      <br />
      <button onClick={submit}>Submit</button>
      {displayResults()}
    </>
  )
}

export default Search;