<template>
  <div class='modal-container'>
    <div class='modal'>
      <div class='content'>
        {{ this.message }}
      </div>
      <div class='button-bar'>
        <button class='accept' @click="Accept"> {{ this.acceptText }} </button>
        <button class='reject' @click="Reject"> {{ this.rejectText }} </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['message', 'acceptText', 'rejectText'],
  methods: {
    Accept: function () {
      this.$emit('Accept');
    },
    Reject: function () {
      this.$emit('Reject');
    }
  }
};
</script>

<style scoped>
.modal-container {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}
.modal {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 50%;
}
.content {
  margin: 25px;  
}
button, a {
  border-radius: 0;
  border: none;
}
button:hover {
  cursor: pointer;
}
.button-bar {
  display: flex;
  margin-top: 25px;
  align-items: center;
  margin-left: 10px;
}
.accept {
  background-color: #529A7F;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  color: #fff;
  margin-right: 10px;
}
.accept:hover {
  background-color: #176548;
}
.reject {
  background: none;
  line-height: 16px;
  font-size: 16px;
  color: #325778;
}
.reject:hover {
  text-decoration: underline;
}
</style>