import React from 'react';

export default ({markdown, setMarkdown, width}) => {
  const handleChange = function(e) {
    const { value } = e.target;
    setMarkdown(value);
  };
  return (
    <textarea className="editor__markdown"  
      style={{width: width + '%'}} 
      onChange={handleChange}
      value={markdown}
    >
    </textarea>
  )
};