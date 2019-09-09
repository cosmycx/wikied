//const ES_API = "somewhereongoogle.com";
import projectPage from '../mock/projectpage';

function toJSON(response) {
  // why do this: https://www.tjvantoll.com/2015/09/13/fetch-and-errors/
  if(!response.ok) {
    throw Error(response.statusText);
  }
  return response.json();
}

export function saveMarkdownFiles(body, user) {
  alert(user + ' your markdown is all saved!');
  console.log(body);
  //return fetch(ES_API).then(toJSON);
}

export function getMarkdown(docId) {
  return fetch('https://jsonplaceholder.typicode.com/todos/1').then(response => {
    return projectPage
  })
}