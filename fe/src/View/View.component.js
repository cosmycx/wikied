import React, {useState, useEffect} from 'react';
import Preview from '../modules/Preview';
import Editor from '../Editor';
import {Switch, Route, Link} from 'react-router-dom';
import {Button} from 'react-bootstrap';
import { getMarkdown } from '../utils/api';

const bannerStyle = {height: '80px'};
const linkStyle = {position: 'absolute', left: '20px', top: '20px'};
const editStyle = {position: 'absolute', right: '20px', top: '20px'};

function View(props) {
  const [markdown, setMarkdown] = useState('');
  const { docId } = props;
  useEffect(() => {
    const fetchMarkdown = async () => {
      const {content} = await getMarkdown(docId);
      setMarkdown(content);
    }
    if(!markdown) {
      if(docId) {
        fetchMarkdown();
      }
    }
  }, [markdown, docId])

  return (
    <>
      <div style={bannerStyle}>
        <Link style={linkStyle} to="/"><Button variant="outline-info">{`< Back to Search`}</Button></Link>
      </div>
      <Switch>
        <Route path="/view/:docId/edit" render={({match}) => (
          <Editor markdown={markdown} />
        )} />
        <Route path="/view/:docId" render={({match}) => (
          <>
            <Link style={editStyle} to={`/view/${match.params.docId}/edit`}><Button variant="outline-primary">{`Edit`}</Button></Link>
            <Preview markdown={markdown} />
          </>
          
        )} />
      </Switch>
    </>
  )
}

export default View;
