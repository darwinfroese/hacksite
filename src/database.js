let id = 0;
let projects = [];

const AddProject = (name) => {
  projects.push({
    id: ++id,
    name: name,
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

export default () => {
  return {
    GetProjects: GetProjects,
    GetProject: GetProject,
    AddProject: AddProject
  };
};
