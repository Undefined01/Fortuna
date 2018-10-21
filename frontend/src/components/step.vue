<template>
  <div class="step">
    <span class="step_id">{{ id }}</span>
    <slot/>
  </div>
</template>
<script>
export default {
  name: 'Step',
  props: {
    id: {
      type: Number,
      default: 1
    }
  }
}
</script>
<style scoped>
.step {
  position: relative;
  margin-bottom: 20px;
  padding: 15px;
  padding-left: 30px;

  background-color: #eee;
  border-left: 5px solid #0af;
}
.step_id {
  position: absolute;
  width: 20px;
  height: 20px;
  top: 10px;
  left: -13px;

  color: white;
  background-color: #0af;
  border-radius: 100%;

  text-align: center;
  line-height: 20px;
}
</style>
