import * as React from 'react';

type result = { title: String, pageid: Number, click: Function };
const Result = ({ title, pageid, click }: result) => {

  return (
    <div id={pageid.toString()} onClick={() => click([pageid, title])}>{title}</div>
  )
}

export default Result;