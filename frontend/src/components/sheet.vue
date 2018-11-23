<template>
  <div class="maximize">
    <el-table
      :data="tableData"
      :summary-method="getAverage"
      :default-sort="{prop: '总分', order: 'descending'}"
      border
      stripe
      show-summary
      class="maximize"
      height="85vh">
      <el-table-column
        v-for="(item, index) in data.Cols"
        :key="index"
        :prop="item"
        :label="item"
        sortable/>
    </el-table>
  </div>
</template>
<script>
import elTable from 'element-ui/lib/table'
import elTableColumn from 'element-ui/lib/table-column'
import 'element-ui/lib/theme-chalk/table.css'
import 'element-ui/lib/theme-chalk/table-column.css'

export default {
  name: 'Sheet',
  components: { elTable, elTableColumn },
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      average: []
    }
  },
  computed: {
    tableData () {
      let res = []
      let template = {}
      this.data.Cols.map(col => {
        template[col] = ''
      })
      this.data.Data.map(row => {
        let obj = Object.assign({}, template)
        row.map((data, i) => {
          obj[this.data.Cols[i]] = data
        })
        res.push(obj)
      })
      return res
    }
  },
  watch: {
    data () {
      this.calcAverage()
    }
  },
  mounted () {
    this.calcAverage()
  },
  methods: {
    calcAverage () {
      let sum = []
      let count = []
      this.data.Cols.map(() => {
        sum.push(0)
        count.push(0)
      })
      this.data.Data.map(row => {
        row.map((num, i) => {
          if (!isNaN(Number(num))) {
            sum[i] += Number(num)
            count[i]++
          }
        })
      })
      this.average = []
      sum.map((sum, i) => {
        if ((i & 1) === 0) {
          this.average[i] = (sum / count[i]).toFixed(2)
        }
      })
      this.average[0] = '平均分'
    },
    getAverage () {
      return this.average
    }
  }
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
.el-table__body-wrapper {
  overscroll-behavior: contain;
}
</style>
