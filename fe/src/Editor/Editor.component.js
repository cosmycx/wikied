import React, {useState} from 'react';
import { 
  Markdown, 
  VerticalHandler, 
} from './modules';
import {
  Preview
} from '../modules';
import { saveMarkdownFiles } from '../utils/api';
import './styles.scss';

export default () => {
  const [markdown, setMarkdown] = useState('# Add Markdown Here');
  const [viewWidths, setViewWidths] = useState({left: 49, right: 49});

  const publish = function() {
    saveMarkdownFiles(markdown, 'Sally');
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