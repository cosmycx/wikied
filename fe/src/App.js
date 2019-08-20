import React from 'react';
import Editor from './Editor';
import MarkdownOutput from './MarkdownOutput';
import {Route} from 'react-router-dom';

function App() {
  return (
    <div>
      <Route path="/markdown/edit" component={Editor} />
      <Route path="/markdown" component={MarkdownOutput} /> 
    </div>
  );
}

export default App;
