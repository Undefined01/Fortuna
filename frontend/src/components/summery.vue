<template>
  <el-table
    :data="this.data"
    border
    show-summary
    :summary-method="getAverage"
    style="width: 100%">
    <el-table-column
      prop="name"
      label="姓名">
    </el-table-column>
  </el-table>
</template>
<script>
export default {
  name: 'summery',
  props: {
    data: {
      type: Array,
      required: true
    }
  },
  methods: {
    getAverage (param) {
      const { columns, data } = param
      const result = ['平均分']
      columns.forEach((column, index) => {
        if (index === 0) return
        const values = data.map(item => Number(item[column.property]))
        if (!values.every(value => isNaN(value))) {
          let sum = 0
          values.some(item => { sum += item })
          result[index] = (sum / values.length).toFixed(2)
        } else result[index] = ''
      })

      return result
    }
  }
}
</script>
