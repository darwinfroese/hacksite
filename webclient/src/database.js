import 'whatwg-fetch';

// TODO: Handle responses using status codes
// TODO: Display errors to the user
// TODO: Spinners/progress indicators on requests

// TODO: This should be a build time value
// const apiBaseUrl = 'http://localhost:8800/api/v1';
const apiBaseUrl = '/api/v1';

export const AddProject = (p) => {
  return fetch(apiBaseUrl + '/projects', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    mode: 'cors',
    credentials: 'include',
    body: JSON.stringify(p)
  });
};

export const GetProject = (id) => {
  let url = apiBaseUrl + '/project?id=' + id;
  return fetch(url, { credentials: 'include', mode: 'cors' });
};

export const GetProjects = () => {
  return fetch(apiBaseUrl + '/projects', { credentials: 'include', mode: 'cors' });
};

export const UpdateProject = (project) => {
  return fetch(apiBaseUrl + '/projects', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(project)
  });
};

export const RemoveProject = (project) => {
  return fetch(apiBaseUrl + '/projects', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(project)
  });
};

export const UpdateTask = (task) => {
  return fetch(apiBaseUrl + '/tasks', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(task)
  });
};

export const RemoveTask = (task) => {
  return fetch(apiBaseUrl + '/tasks', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(task)
  });
};

export const AddEvolution = (evolution) => {
  return fetch(apiBaseUrl + '/evolution', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(evolution)
  });
};

export const ChangeCurrentEvolution = (evolution) => {
  return fetch(apiBaseUrl + '/currentevolution', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(evolution)
  });
};

export const CreateAccount = (account) => {
  return fetch(apiBaseUrl + '/accounts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(account)
  });
};

export const Login = (account) => {
  return fetch(apiBaseUrl + '/login', {
    method: 'GET',
    headers: {
      'Authorization': 'Basic ' + btoa(account.Username + ':' + account.Password)
    },
    mode: 'cors',
    credentials: 'include'
  });
};

export const Authenticate = () => {
  return fetch(apiBaseUrl + '/session', {
    method: 'GET',
    mode: 'cors',
    credentials: 'include'
  });
};

export const Logout = () => {
  return fetch(apiBaseUrl + '/logout', {
    method: 'GET',
    mode: 'cors',
    credentials: 'include'
  });
};
