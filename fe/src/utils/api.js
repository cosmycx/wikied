//const ES_API = "somewhereongoogle.com";

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