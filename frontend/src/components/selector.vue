<template>
  <div>
    <step id="1">
      <h4>选择班级：</h4>
      <el-select
        v-model="queryParam.class"
        placeholder="请选择">
        <el-option
          v-for="(item, index) in classList"
          :key="index"
          :label="item"
          :value="item"/>
      </el-select>
    </step>
    <step id="2">
      <div @dblclick="inputing = true">
        <h4>选择考试：</h4>
        <input
          v-if="inputing"
          v-model="queryParam.exam"
          autofocus="autofocus"
          @blur="inputing = false"
        >
        <el-select
          v-else
          v-model="queryParam.exam"
          placeholder="请选择">
          <el-option
            v-for="(item, index) in examList"
            :key="index"
            :label="item"
            :value="item"/>
        </el-select>
      </div>
    </step>
    <step id="3">
      <btn
        default="查询成绩"
        @click="selected" />
    </step>
  </div>
</template>
<script>
import step from '@/components/step'
import btn from '@/components/btn'
import ajax from '@/plugins/ajax'

import 'element-ui/lib/theme-chalk/icon.css'
import 'element-ui/lib/theme-chalk/select.css'
import 'element-ui/lib/theme-chalk/option.css'
import elSelect from 'element-ui/lib/select'
import elOption from 'element-ui/lib/option'

export default {
  name: 'Selector',
  components: { elSelect, elOption, step, btn },
  data () {
    return {
      inputing: false,
      classList: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12', '13', '14', '15', '16', '17', '18', '19', '20'],
      examList: [],
      queryParam: { class: '', exam: '' }
    }
  },
  mounted () {
    this.UpdateExamList()
  },
  methods: {
    UpdateExamList () {
      ajax.get('http://localhost:8081/api/examlist')
        .then(res => {
          this.examList = JSON.parse(res)
        })
        .catch(() => {
          console.log('Update ExamList failed!')
        })
    },
    selected (btn) {
      this.$emit('selected', this.queryParam, btn.changeState)
    }
  }
}
</script>
<style scoped>
h1, h2, h3, h4, h5, h6, p {
  margin: 0 0 .5rem 0;
}
</style>
