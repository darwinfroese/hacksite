import 'whatwg-fetch';

const devApiBase = 'http://localhost:8800/api/v1';
// const prodApiBase = '/api/v1';

const AddProject = (p) => {
  return fetch(devApiBase + '/projects', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(p)
  });
};

const GetProject = (id) => {
  let url = devApiBase + '/project?id=' + id;
  return fetch(url);
};

const GetProjects = () => {
  return fetch(devApiBase + '/projects');
};

const UpdateTask = (task) => {
  return fetch(devApiBase + '/tasks', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(task)
  });
};

const RemoveProject = (project) => {
  return fetch(devApiBase + '/projects', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(project)
  });
};

const RemoveTask = (task) => {
  return fetch(devApiBase + '/tasks', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(task)
  });
};

export default {
  AddProject,
  GetProject,
  GetProjects,
  UpdateTask,
  RemoveProject,
  RemoveTask
};
