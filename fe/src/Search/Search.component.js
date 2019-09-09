import React, { useState } from 'react';
import {Form, Container, Row, Col, Button, ListGroup} from 'react-bootstrap';
import {Link} from 'react-router-dom';
import results from '../mock/searchResults';

const linkStyle = {position: 'absolute', top: '20px', right: '20px'}

const Item = ({id, title, description}) => ( 
  <Link to={`/view/${id}`}>
    <ListGroup.Item>
      <h1>{title}</h1>
      <p>{description}</p>
    </ListGroup.Item>
  </Link>
);

function Search() {
  const [searchResults, setResults] = useState([])

  function search(event) {
    event.preventDefault();
    setTimeout(() => {
      setResults(results)
    }, 1000);
  }
 
  return (
    <>
      <Link style={linkStyle} to="/new"><Button variant="outline-info">Add New</Button></Link>
      <Form onSubmit={search}>
        <Form.Group controlId="searchTerm">
          <Container style={{width: '300px', marginTop: '20vh'}} >
            <Row className="justify-content-center text-center">
              <Col>
                <Form.Label><h1>Wikied</h1></Form.Label>
              </Col>
            </Row>
            <Row className="justify-content-sm-center">
              <Col>
                <Form.Control type="input" placeholder="Search"></Form.Control>
              </Col>
            </Row>
          </Container>
        </Form.Group>
      </Form>
      {!!searchResults.length &&
        <ListGroup>
          {searchResults.map(result => <Item key={result.id} {...result} />)}
        </ListGroup>
      }
    </>
  )
}

export default Search;