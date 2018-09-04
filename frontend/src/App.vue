<template>
  <el-container>
    <el-header>
      <selector v-model="exam" />
    </el-header>
    <el-main>
      <el-tabs type="border-card">
        <!-- <el-tab-pane label="总成绩" :key="-1"><summery :data="summeryData" /></el-tab-pane> -->
        <el-tab-pane v-for="(subject, index) in summeryData" :key="index" :label="subject.subject">
          <detail :data="subject.data" />
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<script>
import ajax from '@/assets/ajax.js'

import selector from '@/components/selector'
import summery from '@/components/summery'
import detail from '@/components/detail'

export default {
  name: 'App',
  data () {
    return {
      exam: '',
      summeryData: []
    }
  },
  watch: {
    exam () {
      this.updateData()
    }
  },
  methods: {
    updateData () {
      const subjectList = ['语文', '数学', '数学文', '数学理', '英语', '物理', '化学', '生物', '历史', '地理', '政治']
      this.summeryData = []

      for (let i in subjectList) {
        ajax({
          url: 'http://localhost:8081/api',
          data: {type: 'subscore', exam: this.exam, class: '09', subject: subjectList[i]},
          success: (res) => {
            if (res === '{}') return
            res = JSON.parse(res)
            this.summeryData.push({subject: subjectList[i], data: res})
            this.summeryData.sort((x, y) => subjectList.indexOf(x.subject) > subjectList.indexOf(y.subject))
          }
        })
      }
    }
  },
  components: { selector, summery, detail }
}
</script>
<style>
.el-table td {
  padding-top: 0;
  padding-bottom: 0;
}
</style>
