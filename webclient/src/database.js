import 'whatwg-fetch';

const AddProject = (p) => {
  return fetch('http://localhost:8800/api/v1/projects', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(p)
  });
};

const GetProject = (id) => {
  let url = 'http://localhost:8800/api/v1/project?id=' + id;
  return fetch(url);
};

const GetProjects = () => {
  return fetch('http://localhost:8800/api/v1/projects');
};

const UpdateTask = (task) => {
  return fetch('http://localhost:8800/api/v1/tasks', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(task)
  });
};

const RemoveProject = (project) => {
  return fetch('http://localhost:8800/api/v1/projects', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(project)
  });
};

const RemoveTask = (task) => {
  return fetch('http://localhost:8800/api/v1/tasks', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(task)
  });
};

// Temporary Project filling
// AddProject({
//   name: 'Hacksite',
//   description: 'A website for listing weekend hack projects',
//   completed: true,
//   tasks: [
//     {task: 'Add and View projects', id: 1, completed: true},
//     {task: 'Complete Tasks', id: 2, completed: true},
//     {task: 'Remove projects and tasks', id: 3, completed: true},
//     {task: 'Complete Projects', id: 4, completed: true}
//   ]
// });

// AddProject({
//   name: 'Hacksite v0.2',
//   description: 'Second iteration of hacksite, building the backend',
//   completed: false,
//   tasks: [
//     {task: 'Setup API and file serving', id: 1, completed: true},
//     {task: 'Setup database', id: 2, completed: true},
//     {task: 'Setup API endpoints for tasks and projects', id: 3, completed: true},
//     {task: 'Replace javascript storage with API calls', id: 4, completed: false}
//   ]
// });

export default {
  AddProject,
  GetProject,
  GetProjects,
  UpdateTask,
  RemoveProject,
  RemoveTask
};
