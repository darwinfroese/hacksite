let id = 0;
let projects = [];

const AddProject = (project) => {
  projects.push({
    id: ++id,
    name: project.name,
    description: project.description,
    tasks: project.tasks,
    details: '/details/' + id
  });
};

const GetProject = (id) => {
  let filtered = projects.filter((project) => {
    return project.id === id;
  });
  return filtered[0];
};

const GetProjects = () => {
  return projects;
};

const UpdateTask = (projectId, task) => {
  let project = projects[projectId - 1];
  project.tasks[task.idx - 1] = task;
};

const RemoveProject = (projectId) => {
  projects = projects.filter((project) => {
    return project.id !== projectId;
  });
};

AddProject({
  name: 'Hacksite',
  description: 'A website for listing weekend hack projects',
  tasks: [
    {task: 'Add and View projects', idx: 1, completed: true},
    {task: 'Complete Tasks', idx: 2, completed: true},
    {task: 'Remove projects and tasks', idx: 3, completed: false},
    {task: 'Complete Projects', idx: 4, completed: false}
  ]
});

export default {
  AddProject,
  GetProject,
  GetProjects,
  UpdateTask,
  RemoveProject
};
