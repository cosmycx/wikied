import React, {useState, useEffect} from 'react';
import {
  Markdown, 
  VerticalHandler, 
  Preview
} from '../modules';
import { saveMarkdownFiles, getMarkdown } from '../utils/api'; 
import './styles.scss';

export default (props) => {
  const [markdown, setMarkdown] = useState(props.markdown);
  const [viewWidths, setViewWidths] = useState({left: 49, right: 49});

  const publish = function() {
    debugger;
    saveMarkdownFiles(markdown).then(response => {
      props.history.push('/view/' + response.id)
    });
  }

  return (
    <div className="editor">
      <button style={{left: viewWidths.left - 6 + '%'}} onClick={publish}>Publish</button>
      {/* setting inline style. Eventually these will be set by a drag handler on the verticalhandler */}
      <Markdown {...{markdown, setMarkdown}} width={viewWidths.left}/>
      <VerticalHandler {...{setViewWidths}}/>
      <Preview {...{markdown}} width={viewWidths.right}/>
    </div>
  );
}