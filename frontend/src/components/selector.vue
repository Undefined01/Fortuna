<template>
  <div>
    <el-select :value="value" @input="updateValue($event)" placeholder="请选择">
      <el-option
        v-for="(item, index) in ExamList" :key="index" :label="item" :value="item">
      </el-option>
    </el-select>
  </div>
</template>
<script>
import ajax from '@/assets/ajax.js'

export default {
  name: 'summery',
  props: ['value'],
  data () {
    return { ExamList: ['正在获取考试列表'] }
  },
  mounted () {
    ajax({
      url: 'http://localhost:8081/api',
      data: {type: 'exam'},
      success: (res) => {
        res = JSON.parse(res)
        this.ExamList = res
      }
    })
  },
  methods: {
    updateValue (value) {
      this.$emit('input', value)
    }
  }
}
</script>
