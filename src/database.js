let id = 0;
let projects = [];

const AddProject = (project) => {
  projects.push({
    id: ++id,
    name: project.name,
    description: project.description,
    tasks: project.tasks,
    details: '/details/' + id,
    completed: project.completed
  });
};

const GetProject = (id) => {
  let filtered = projects.filter((project) => {
    return project.id === id;
  });

  // there can be only one, return it
  return filtered[0];
};

const GetProjects = () => {
  return projects;
};

const UpdateTask = (projectId, task) => {
  let project = GetProject(projectId);

  // Since the tasks may not line up 1:1 id to index
  project.tasks.forEach((t, idx) => {
    if (task.idx === t.idx) {
      project.tasks[idx] = task;
      return;
    }
  });

  UpdateProject(project);
};

const RemoveProject = (projectId) => {
  projects = projects.filter((project) => {
    return project.id !== projectId;
  });
};

const RemoveTask = (projectId, task) => {
  let project = GetProject(projectId);
  let tasks = project.tasks;

  tasks = tasks.filter((t) => {
    return t.idx !== task.idx;
  });

  project.tasks = tasks;
};

const UpdateProject = (project) => {
  let unfinishedTasks = project.tasks.filter((task) => {
    return !task.completed;
  });

  if (unfinishedTasks.length === 0) {
    project.completed = true;
  } else {
    project.completed = false;
  }
};

// Temporary Project filling
AddProject({
  name: 'Hacksite',
  description: 'A website for listing weekend hack projects',
  completed: true,
  tasks: [
    {task: 'Add and View projects', idx: 1, completed: true},
    {task: 'Complete Tasks', idx: 2, completed: true},
    {task: 'Remove projects and tasks', idx: 3, completed: true},
    {task: 'Complete Projects', idx: 4, completed: true}
  ]
});

export default {
  AddProject,
  GetProject,
  GetProjects,
  UpdateTask,
  RemoveProject,
  RemoveTask
};
