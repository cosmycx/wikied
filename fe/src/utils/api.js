//const ES_API = "somewhereongoogle.com";
import projectPage from '../mock/projectpage';

function toJSON(response) {
  // why do this: https://www.tjvantoll.com/2015/09/13/fetch-and-errors/
  if(!response.ok) {
    throw Error(response.statusText);
  }
  return response.json();
}

const ROOT = 'http://localhost:4040';

export function saveMarkdownFiles(body) {
  // alert(user + ' your markdown is all saved!');
  // console.log(body);
  //return fetch(ES_API).then(toJSON);
  const options = {
    method: 'POST',
    body: JSON.stringify({
      content: body
    })
  }
  return fetch(ROOT +'/createinfopost', options).then(toJSON)
}

export function getMarkdown(docId) {
  return fetch(ROOT + '/getinfopostbyid?id=' + docId).then(toJSON);
}

export function searchMarkdown(searchTerm) {
  return fetch(ROOT + '/searchterm?search=' + searchTerm).then(toJSON);
}