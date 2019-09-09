import React from 'react';
import Editor from './Editor';
import Search from './Search';
import View from './View';
import {Switch, Route} from 'react-router-dom';

function App() {
  return (
    <Switch>
      <Route path="/new" render={() => {
        return <Editor markdown='# Add Markdown Here' />
      }}/>
      <Route path="/view/:docId" render={({match}) => (
        <View docId={match.params.docId} />
      )} />
      <Route exact path="/" component={Search} />
    </Switch>
  );
}

export default App;
