import 'whatwg-fetch';

// TODO: Handle responses using status codes
// TODO: Display errors to the user
// TODO: Spinners/progress indicators on requests

// TODO: This should be a build time value
const devApiBase = 'http://localhost:8800/api/v1';
// const prodApiBase = '/api/v1';

export const AddProject = (p) => {
  return fetch(devApiBase + '/projects', {
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
  let url = devApiBase + '/project?id=' + id;
  return fetch(url, { credentials: 'include', mode: 'cors' });
};

export const GetProjects = () => {
  return fetch(devApiBase + '/projects', { credentials: 'include', mode: 'cors' });
};

export const UpdateProject = (project) => {
  return fetch(devApiBase + '/projects', {
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
  return fetch(devApiBase + '/projects', {
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
  return fetch(devApiBase + '/tasks', {
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
  return fetch(devApiBase + '/tasks', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(task)
  });
};

export const AddIteration = (iteration) => {
  return fetch(devApiBase + '/iteration', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(iteration)
  });
};

export const ChangeCurrentIteration = (iteration) => {
  return fetch(devApiBase + '/currentiteration', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    mode: 'cors',
    body: JSON.stringify(iteration)
  });
};

export const CreateAccount = (account) => {
  return fetch(devApiBase + '/accounts', {
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
  return fetch(devApiBase + '/login', {
    method: 'GET',
    headers: {
      'Authorization': 'Basic ' + btoa(account.Username + ':' + account.Password)
    },
    mode: 'cors',
    credentials: 'include'
  });
};

export const Authenticate = () => {
  return fetch(devApiBase + '/session', {
    method: 'GET',
    mode: 'cors',
    credentials: 'include'
  });
};

export const Logout = () => {
  return fetch(devApiBase + '/logout', {
    method: 'GET',
    mode: 'cors',
    credentials: 'include'
  });
};
