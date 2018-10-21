<template>
  <button
    :class="'btn btn-' + color"
    @click="handler">
    <div>
      <div class="hidden">{{ text }}</div>
      <transition :name="transition">
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
    default: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      text: this.default,
      color: 'primary',
      transition: 'up',

      text1: this.default,
      text2: '',
      odd: true
    }
  },
  methods: {
    changeText (text, color, transition) {
      if (this.odd) this.text2 = text
      else this.text1 = text
      this.text = text
      this.color = color || 'primary'
      this.transition = transition || 'up'

      this.odd = !this.odd
    },
    changeStatus (status) {
      if (status === 0) {
        this.changeText(this.default, 'primary', 'down')
      }
      if (status === 1) {
        this.changeText('正在' + this.default, 'warning', 'up')
      }
      if (status === 2) {
        this.changeText(this.default + '成功', 'success', 'up')
        setTimeout(() => this.changeStatus(0), 1500)
      }
      if (status === 3) {
        this.changeText(this.default + '失败', 'danger', 'up')
        setTimeout(() => this.changeStatus(0), 1500)
      }
    },
    handler () {
      this.$emit('click', {
        changeText: (text, color, transition) => {
          this.changeText(text, color, transition)
        },
        changeStatus: (status) => {
          this.changeStatus(status)
        }
      })
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
