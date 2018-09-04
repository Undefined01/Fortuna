<template>
  <div class="container">
    <h2 class="title">成绩查询</h2>
    <step id="1">
      <h4>选择班级：</h4>
      <el-select v-model="query.class" placeholder="请选择">
        <el-option
          v-for="(item, index) in classList" :key="index" :label="item" :value="item">
        </el-option>
      </el-select>
    </step>

    <step id="2">
      <h4>选择考试：</h4>
      <el-select v-model="query.exam" placeholder="请选择">
        <el-option
          v-for="(item, index) in examList" :key="index" :label="item" :value="item">
        </el-option>
      </el-select>
    </step>

    <step id="3">
      <h4>选择操作：</h4>
      <btn @click="btnQueryClick" default="查询成绩"></btn>
      <btn @click="btnExpendClick" default="扩展显示区域"></btn>
    </step>
    <expender :width="expend ? '95vw' : '100%'">
      <el-tabs type="border-card">
        <el-tab-pane v-for="(table, index) in data" :key="index" :label="table.Title">
          <presenter :data="table"></presenter>
        </el-tab-pane>
      </el-tabs>
    </expender>
  </div>
</template>

<script>
import '@/App.css'
import ajax from '@/assets/ajax.js'

import 'element-ui/lib/theme-chalk/icon.css'

import elSelect from 'element-ui/lib/select'
import elOption from 'element-ui/lib/option'
import 'element-ui/lib/theme-chalk/select.css'
import 'element-ui/lib/theme-chalk/option.css'
import elTabs from 'element-ui/lib/tabs'
import elTabPane from 'element-ui/lib/tab-pane'
import 'element-ui/lib/theme-chalk/tabs.css'
import 'element-ui/lib/theme-chalk/tab-pane.css'

import step from '@/components/step'
import btn from '@/components/btn'
import expender from '@/components/expender'
import presenter from '@/components/presenter'

export default {
  name: 'App',
  data () {
    return {
      classList: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12', '13', '14', '15', '16', '17', '18', '19', '20'],
      examList: [],
      query: { class: '', exam: '' },
      data: [],
      expend: false
    }
  },
  mounted () {
    this.UpdateExamList()
  },
  methods: {
    UpdateExamList () {
      ajax({
        url: 'http://localhost:8081/api',
        data: { type: 'examlist' },
        success: (res) => {
          this.examList = JSON.parse(res)
          btn.changeStatus(2)
        },
        error: () => {
          btn.changeStatus(3)
        }
      })
    },
    btnQueryClick (btn) {
      btn.changeStatus(1)
      ajax({
        url: 'http://localhost:8081/api',
        data: {type: 'examdata', exam: this.query.exam, class: this.query.class},
        success: (res) => {
          if (res !== '[]') this.data = JSON.parse(res)
          btn.changeStatus(2)
        },
        error: () => {
          btn.changeStatus(3)
        }
      })
    },
    btnExpendClick (btn) {
      this.expend = !this.expend
      btn.changeStatus(1)
      setTimeout(() => {
        btn.changeStatus(0)
      }, 500)
    }
  },
  components: { elSelect, elOption, elTabs, elTabPane, step, btn, expender, presenter }
}
</script>

<style scoped>
.title {
  display: inline-block;
  padding-right: 20px;
  margin-bottom: 30px;
  color: #000;
  border-bottom: 3px solid #000;
}
</style>
