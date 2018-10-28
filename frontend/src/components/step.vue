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
      type: String,
      default: '1'
    }
  }
}
</script>
<style scoped>
.step {
  position: relative;
  margin-bottom: 2rem;
  padding: 1rem;
  background-color: #eee;
  border-left: 5px solid #0af;
}
.step_id {
  position: absolute;
  width: 1.2rem;
  height: 1.2rem;
  top: 1.1rem;
  left: -.8rem;
  color: white;
  background-color: #0af;
  text-align: center;
  line-height: 1.1rem;
  border-radius: 100%;
}
</style>
