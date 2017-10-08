<template>
    <div class = "input">
        <span class='input-label'>{{fieldLabel}}</span>
        <input type='text' :placeholder='placeholder' v-model="model" v-validate.initial = "'required'" 
        :class="{'input': true, 'is-danger': errors.has(fieldName) }" :name= "fieldName" />
        <br/>
        <i v-show="errors.has(fieldName)" class="fa fa-warning"></i>
        <span v-show="errors.has(fieldName)" class="error is-danger">{{ errors.first(fieldName) }}</span>
    </div>
</template>

<script>
    export default {
      props: ['fieldLabel', 'placeholder', 'model', 'fieldName']
    };
</script>

<style scoped>

</style>