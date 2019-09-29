import React, { useState, Component } from 'react';
import {Form, Container, Row, Col, Button, ListGroup} from 'react-bootstrap';
import {Link} from 'react-router-dom';
import { searchMarkdown } from '../utils/api.js';

const linkStyle = {position: 'absolute', top: '20px', right: '20px'}

const Item = ({id, content, searchTerm}) => {
  const match = content.match(new RegExp(searchTerm, 'i'))
  if (match) {
    const start = match.index - 25;
    const end = match.index + searchTerm.length + 25;
    let title = content.match(new RegExp('(?<=\\n#)(.*?)(?=\\n)'));
    if (title > 0) {
      title = title[0];
    }
    const preview = content.substring(start, end);
    return (
      <li>
      <Link to={`/view/${id}`}>
        <h1>{title}</h1>
        <p>{preview}</p>
      </Link>
    </li>
    )
  }// .if

  return null
};

function Search () {
  const [searchResults, setResults] = useState([]);
  const [searchTerm, setTerm] = useState('');
  const search = (event) => {
    event.preventDefault();
    const searchTerm = document.getElementById('searchTerm').value;
    searchMarkdown(searchTerm).then((results) => {
      setResults(results);
    })
  }

  const handleChange = (event) => {
    const {value} = event.target;
    setTerm(value);
    console.log(searchTerm)
  }

  return (
    <>
      <Link style={linkStyle} to="/new"><Button variant="outline-info">Add New</Button></Link>
      <Form onSubmit={search}>
        <Form.Group controlId="searchTerm">
          <Container style={{width: '300px', marginTop: '20vh'}} >
            <Row className="justify-content-center text-center">
              <Col>
                <Form.Label><h1 className="mb-5"><svg width="300" height="71" fill="none" xmlns="http://www.w3.org/2000/svg"><title>Wikied</title><path d="M67.351 29.2l-4.688.625-9.72 39.826H42.45l-8.645-28.754h-.258l-8.644 28.754H14.45L4.688 29.825 0 29.2v-7.858h21.074V29.2l-4.946.982 4.387 22.1h.258l8.688-30.94h8.43l8.773 31.03h.258l4.344-22.145-4.99-1.027v-7.858h21.075V29.2zM71.265 61.837l6.064-1.34V30.54l-6.71-1.34v-7.857h19.268v39.156l6.022 1.34v7.813H71.264v-7.814zm18.622-52.06H77.33V0h12.558v9.778zM98.704 7.858V0h19.225v40.362h2.795l7.226-10.582-4.086-.58v-7.858h22.88V29.2l-5.505 1.294-8.386 11.966 11.741 18.35 4.774 1.027v7.814h-21.504v-7.814l2.58-.446-7.139-11.966h-5.376v11.073l5.419 1.34v7.813H99.349v-7.814l6.064-1.34v-51.3l-6.709-1.339zM153.411 61.837l6.064-1.34V30.54l-6.71-1.34v-7.857h19.268v39.156l6.021 1.34v7.813h-24.643v-7.814zm18.622-52.06h-12.558V0h12.558v9.778zM204.461 70.588c-6.709 0-12.071-2.232-16.085-6.697-4.014-4.465-6.021-10.135-6.021-17.01v-1.787c0-7.173 1.892-13.081 5.677-17.725 3.814-4.643 8.917-6.95 15.311-6.92 6.279 0 11.154 1.964 14.623 5.893 3.469 3.93 5.204 9.242 5.204 15.94v7.099h-27.783l-.086.267c.229 3.185 1.247 5.805 3.053 7.858 1.835 2.054 4.315 3.081 7.441 3.081 2.781 0 5.089-.283 6.924-.848 1.835-.595 3.842-1.518 6.021-2.768l3.398 8.036c-1.921 1.578-4.416 2.902-7.484 3.974-3.039 1.072-6.437 1.607-10.193 1.607zm-1.118-40.093c-2.322 0-4.157.922-5.505 2.768-1.347 1.845-2.179 4.27-2.494 7.277l.129.224h15.397v-1.161c0-2.769-.631-4.971-1.893-6.608-1.233-1.667-3.111-2.5-5.634-2.5zM259.426 63.757c-1.462 2.233-3.197 3.93-5.204 5.09-1.978 1.16-4.244 1.741-6.795 1.741-5.735 0-10.193-2.173-13.376-6.518-3.154-4.346-4.731-10.09-4.731-17.234v-.938c0-7.62 1.591-13.766 4.774-18.44 3.183-4.673 7.656-7.01 13.419-7.01 2.322 0 4.415.537 6.279 1.608 1.864 1.042 3.498 2.53 4.903 4.465V9.197l-6.709-1.339V0h19.267v60.498l6.021 1.34v7.813h-16.902l-.946-5.894zm-17.547-16.921c0 4.226.673 7.56 2.021 10 1.376 2.412 3.613 3.617 6.709 3.617 1.835 0 3.427-.387 4.774-1.16 1.348-.775 2.452-1.89 3.312-3.35V35.45c-.86-1.547-1.964-2.738-3.312-3.571-1.319-.864-2.881-1.295-4.688-1.295-3.068 0-5.304 1.414-6.709 4.241-1.405 2.828-2.107 6.519-2.107 11.073v.938zM300 60.507c0 5.05-3.944 9.144-8.808 9.144-4.865 0-8.808-4.094-8.808-9.144s3.943-9.144 8.808-9.144c4.864 0 8.808 4.094 8.808 9.144z" fill="#055EA9"/></svg></h1></Form.Label>
              </Col>
            </Row>
            <Row className="justify-content-sm-center">
              <Col>
                <Form.Control onChange={handleChange} value={searchTerm} type="input" placeholder="Search"></Form.Control>
              </Col>
            </Row>
          </Container>
        </Form.Group>
      </Form>
      {!!searchResults &&
        <ul className="list-unstyled pl-0">
          {searchResults.map(result => <Item key={result.id} {...result} searchTerm={searchTerm} />)}
        </ul>
      }
    </>
  )
}

export default Search;