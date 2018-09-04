<template>
  <el-table
    :data="tableData.data"
    border
    stripe
    show-summary
    :summary-method="getAverage"
    style="width: 100%">
    <el-table-column
      prop="name"
      label="姓名">
    </el-table-column>
    <el-table-column
      v-for="(item, index) in tableData.column"
      :key="index"
      :prop="item"
      :label="item"
      min-width="40px">
    </el-table-column>
  </el-table>
</template>
<script>
export default {
  name: 'detail',
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  computed: {
    tableData () {
      let res = {column: [], data: []}
      for (let k in this.data) {
        for (let id in this.data[k]) {
          res.column.push(id)
        }
        break
      }
      for (let k in this.data) {
        let data = {}
        Object.assign(data, this.data[k])
        data['name'] = k
        res.data.push(data)
      }
      return res
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
