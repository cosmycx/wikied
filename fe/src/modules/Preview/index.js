import React from 'react';
import ReactMarkdown from 'react-markdown';
import './style.scss';

export default ({markdown, width}) => (
  <div className="preview__wrapper" style={{width: width + '%'}}>
    <ReactMarkdown source={markdown} />
  </div>
)