<template>
  <div class="main-stage flex-column scroll-y">
    <div class="box-effect">
      <h2 class="container">成绩查询</h2>
    </div>
    <div class="container box-effect content">
      <selector @selected="query" />
      <a
        ref="anchor"
        href="#anchor"
      />
      <a name="anchor" />
      <expender :width="data.length == 0 ? '100%' : '98vw'">
        <presenter :data="data" />
      </expender>
    </div>
    <div class="secondary text-center">Powered by Fortuna v0.4</div>
  </div>
</template>

<script>
import '@/basic.css'
import ajax from '@/plugins/ajax'
const selector = () => import('@/components/selector')
const expender = () => import('@/components/expender')
const presenter = () => import('@/components/presenter')

export default {
  name: 'App',
  components: { selector, expender, presenter },
  data () {
    return {
      data: []
    }
  },
  methods: {
    query (queryParam, changeState) {
      changeState(1)
      ajax.get('http://localhost:8081/api/examdata', {
        exam: queryParam.exam,
        class: queryParam.class
      }).then(res => {
        this.data = JSON.parse(res)
        changeState(2)
        this.$nextTick(() => {
          this.$refs.anchor.click()
        })
      }).catch(() => {
        changeState(3)
      })
    }
  }
}
</script>
<style scoped>
.content {
  padding: 1rem;
}
</style>
