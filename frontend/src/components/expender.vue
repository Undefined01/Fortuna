<template>
  <div
    :style="outter"
    class="outter">
    <div class="inner">
      <slot/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Expender',
  props: {
    width: {
      type: String,
      default: '100%'
    }
  },
  computed: {
    outter () {
      return {
        left: 'calc(50% - (' + this.width + ') / 2)',
        width: this.width
      }
    }
  }
}
</script>

<style scoped>
.outter {
  position: relative;
  transition: left .5s, width .5s;
}
.inner {
  width: 100%;
}
</style>
