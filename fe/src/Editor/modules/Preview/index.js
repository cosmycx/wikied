import React from 'react';
import ReactMarkdown from 'react-markdown';

export default ({markdown, width}) => (
  <div className="editor__preview" style={{width: width + '%'}}>
    <ReactMarkdown  source={markdown} />
  </div>
)