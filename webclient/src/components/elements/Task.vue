<template>
  <div class='task'>
    <label class='taskItem'>
      <input type='checkbox' class='checkbox' @click="Update()" :checked="task.Completed">
      <span v-bind:class="{ completed: task.Completed }"> {{ task.Task }} </span>
    </label>
    <span class='removeButton' @click="RemoveTask()"> <i class='fa fa-times'></i> </span>
    <Modal 
      :message="CongratulationsMessage"
      :acceptText="'Yes'"
      :rejectText="'No'"
      v-on:Accept="NavToEvolution"
      v-on:Reject="CloseModal"
      v-if="renderDialog" />
  </div>
</template>

<script>
import router from '@/router';
import { UpdateTask, RemoveTask } from '@/database';
import Modal from '@/components/elements/Modal';

export default {
  props: ['task', 'pid', 'pname'],
  components: {
    'Modal': Modal
  },
  computed: {
    CongratulationsMessage: function () {
      return 'Congratulations! You completed your current evolution of ' + this.pname +
        '.\n\nWould you like to start a new one?';
    }
  },
  data () {
    return {
      renderDialog: false
    };
  },
  methods: {
    RenderDialog: function () {
      this.renderDialog = true;
    },
    NavToEvolution: function () {
      router.push('/evolution/' + this.pid);
    },
    CloseModal: function () {
      this.renderDialog = false;
    },
    Update: function () {
      this.task.Completed = !this.task.Completed;

      UpdateTask(this.task)
        .then((response) => {
          return response.json();
        })
        .then((project) => {
          if (project.Status === 'Completed') {
            this.RenderDialog();
          }
        });
    },
    RemoveTask: function () {
      RemoveTask(this.task)
        .then((response) => {
          this.$emit('GetProject');
        });
    }
  }
};
</script>

<style scoped>
input {
  display: inline-block;
  width: 16px;
  height: 16px;
}
.task {
  display: block;
  font-size: 16px;
  line-height: 16px;
  width: 90%;
  position: relative;
  padding: 2px;
}
.taskItem:hover {
  cursor: pointer;
}
.completed {
  text-decoration: line-through;
  font-style: italic;
  color: slategray;
}
.removeButton {
  display: none;
  width: 16px;
  height: 16px;
  line-height: 16px;
  vertical-align: middle;
  float: right;
  margin: auto;
}
.removeButton:hover {
  cursor: pointer;
}
.removeButton > i {
  height: 16px; 
  width: 16px;
}
.task:hover .removeButton {
  display: inline;
}
</style>
