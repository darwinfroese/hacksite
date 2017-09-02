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

export default {
  AddProject,
  GetProject,
  GetProjects
};
