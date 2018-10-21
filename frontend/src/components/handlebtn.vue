<template>
  <button
    :class="'btn btn-' + parm.color"
    @click="handler">
    <div>
      <div class="hidden">{{ parm.text }}</div>
      <transition :name="parm.transition">
        <div
          v-if="odd"
          key="text1"
          class="text">
          {{ text1 }}
        </div>
        <div
          v-else
          key="text2"
          class="text">
          {{ text2 }}
        </div>
      </transition>
    </div>
  </button>
</template>

<script>
export default {
  name: 'Btn',
  props: {
    parm: {
      type: Object,
      required: true
    }
  },
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
