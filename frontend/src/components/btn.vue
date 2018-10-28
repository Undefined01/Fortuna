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
    changeState (status) {
      if (status === 0) {
        this.changeText(this.default, 'primary', 'down')
      }
      if (status === 1) {
        this.changeText('正在' + this.default, 'warning', 'up')
      }
      if (status === 2) {
        this.changeText(this.default + '成功', 'success', 'up')
        setTimeout(() => this.changeState(0), 1500)
      }
      if (status === 3) {
        this.changeText(this.default + '失败', 'danger', 'up')
        setTimeout(() => this.changeState(0), 1500)
      }
    },
    handler () {
      this.$emit('click', {
        changeText: (text, color, transition) => {
          this.changeText(text, color, transition)
        },
        changeState: (status) => {
          this.changeState(status)
        }
      })
    }
  }
}
</script>

<style scoped>
.hidden { opacity: 0; }

.btn-primary {
  color: #fff;
  background-color: #007bff;
  border-color: #007bff;
}
.btn-primary:hover, .btn-primary:focus {
  background-color: #0069d9;
  border-color: #0062cc;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.5);
}

.btn-success {
  color: #fff;
  background-color: #28a745;
  border-color: #28a745;
}
.btn-success:hover, .btn-success:focus {
  background-color: #218838;
  border-color: #1e7e34;
  box-shadow: 0 0 0 0.2rem rgba(40, 167, 69, 0.5);
}

.btn-warning {
  color: #212529;
  background-color: #ffc107;
  border-color: #ffc107;
}
.btn-warning:hover, .btn-warning:focus {
  background-color: #e0a800;
  border-color: #d39e00;
  box-shadow: 0 0 0 0.2rem rgba(255, 193, 7, 0.5);
}

.btn-danger {
  color: #fff;
  background-color: #dc3545;
  border-color: #dc3545;
}
.btn-danger:hover, .btn-danger:focus {
  background-color: #c82333;
  border-color: #bd2130;
  box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.5);
}

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
