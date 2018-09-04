<template>
  <button :class="'btn btn-' + this.parm.color"
    @click="handler">
    <div>
      <div class="hidden">{{ this.parm.text }}</div>
      <transition :name="this.parm.transition">
        <div class="text" v-if="this.odd" key="text1">
          {{ this.text1 }}
        </div>
        <div class="text" v-else key="text2">
          {{ this.text2 }}
        </div>
      </transition>
    </div>
  </button>
</template>

<script>
export default {
  name: 'btn',
  props: ['parm'],
  data () {
    return {
      text1: this.parm.text,
      text2: '',
      odd: true
    }
  },
  watch: {
    parm: {
      deep: true,
      handler (newValue, oldValue) {
        if (this.odd) this.text2 = this.parm.text
        else this.text1 = this.parm.text
        this.odd = !this.odd
      }
    }
  },
  methods: {
    handler () {
      this.$emit('click')
    }
  }
}
</script>

<style scoped>
.hidden { opacity: 0; }
.btn > div {
  position: relative;
  overflow: hidden;
}
.text {
  position: absolute;
  top: 0;
}

.up-enter {
  top: 100%;
}
.up-enter-active {
  transition: top .15s;
}
.up-enter-to {
  top: 0%;
}
.up-leave {
  opacity: 1;
}
.up-leave-active {
  transition: opacity .15s;
}
.up-leave-to {
  opacity: 0;
}

.down-enter {
  top: -100%;
}
.down-enter-active {
  transition: top .15s;
}
.down-enter-to {
  top: 0%;
}
.down-leave {
  opacity: 1;
}
.down-leave-active {
  transition: opacity .15s;
}
.down-leave-to {
  opacity: 0;
}
</style>
