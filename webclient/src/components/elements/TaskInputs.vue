<template>
  <section id='tasks' class='tasks input'>
    <span class='input-label'>
      Project Tasks
      <i class='fa fa-info-circle info'>
        <span class='info-text'>
          Aigera only lets you select four tasks for your projects to keep them
          small and achievable!
        </span>
      </i>
    </span>
    <input id='taskInput1' type='text' placeholder='Task 1 Description' name='taskInput'>
    <input id='taskInput2' type='text' placeholder='Task 2 Description' name='taskInput'>
    <input id='taskInput3' type='text' placeholder='Task 3 Description' name='taskInput'>
    <input id='taskInput4' type='text' placeholder='Task 4 Description' name='taskInput'>
  </section>
</template>

<style scoped>
input {
  display: block;
  font-size: 16px;
  padding: 5px 10px;
  width: 500px;
  border: none;
  font-weight: 100;
  outline: none;
  color: #325778;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
  margin-top: 10px;
  transition: all 0.5s linear;
}
input:focus {
  border-bottom: 1px solid #325778;
}
.input-label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
}
.tasks {
  margin-top: 10px;
}
.tasks > input {
  margin: 10px 0;
}
.input {
  margin-top: 20px;
}
.info {
  position: relative;
}
.info .info-text {
  visibility: hidden;
  width: 400px;
  background-color: #325778;
  color: #fff;
  text-align: center;
  padding: 10px;
  border-radius: 6px;
  font-size: 12px;
  font-style: italic;

  /* Position the tooltip text - see examples below! */
  position: absolute;
  left: 25px;
  z-index: 1;
}
.info:hover .info-text {
  visibility: visible;
}
</style>
