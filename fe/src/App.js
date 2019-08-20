import React from 'react';
import Editor from './Editor';
import MarkdownOutput from './MarkdownOutput';
import {Route} from 'react-router-dom';

function App() {
  return (
    <>
      <Route path="/markdown/edit" component={Editor} />
      <Route path="/user/:userId/markdown/:markdownId" render={(props) => {
        return <MarkdownOutput markdown="# Here is some markdown" {...props}/>
      }} />
      <Route exact path="/" render={()=> <div>Home</div>} />
    </>
  );
}

export default App;
