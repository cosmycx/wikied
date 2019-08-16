import React, {Fragment, useState} from 'react';
import { 
  Markdown, 
  VerticalHandler, 
  Preview 
} from './modules';
import './styles.scss';

export default () => {
  const [markdown, setMarkdown] = useState('# Add Markdown Here');
  const [viewWidths, setViewWidths] = useState({left: 49, right: 49});

  return (
    <div className="editor">
      {/* setting inline style. Eventually these will be set by a drag handler on the verticalhandler */}
      <Markdown {...{markdown, setMarkdown}} width={viewWidths.left}/>
      <VerticalHandler {...{viewWidths, setViewWidths}}/>
      <Preview {...{markdown}} width={viewWidths.right}/>
    </div>
  );
}