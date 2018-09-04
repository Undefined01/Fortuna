<template>
  <el-table
    :data="tableData"
    border
    stripe
    show-summary
    :summary-method="getAverage"
    :default-sort="{prop: '总分', order: 'descending'}"
    :height="this.windowSize.height * 0.95 + 'px'"
    :max-height="this.windowSize.height * 0.95 + 'px'"
    style="width: 100%">
    <el-table-column
      v-for="(item, index) in data.Cols"
      :key="index"
      :prop="item"
      :label="item"
      sortable
      min-width="40px">
    </el-table-column>
  </el-table>
</template>

<script>
import elTable from 'element-ui/lib/table'
import elTableColumn from 'element-ui/lib/table-column'
import 'element-ui/lib/theme-chalk/table.css'
import 'element-ui/lib/theme-chalk/table-column.css'

export default {
  name: 'presenter',
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  computed: {
    tableData () {
      let res = []
      for (let i in this.data.Data) {
        let row = {}
        for (let k in this.data.Data[i]) {
          row[this.data.Cols[k]] = this.data.Data[i][k]
        }
        res.push(row)
      }
      return res
    },
    windowSize () {
      return {
        width: document.documentElement.clientWidth,
        height: document.documentElement.clientHeight
      }
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
  },
  components: { elTable, elTableColumn }
}
</script>

<style>
.el-table__header th, .el-table td {
  padding-top: 0;
  padding-bottom: 0;
}
.el-table_1_column_2 {
  min-width: 60px!important;
}
</style>
