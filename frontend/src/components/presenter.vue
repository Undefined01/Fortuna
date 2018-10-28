<template>
  <div v-if="data.length === 0">
    <div class="text-center non-data-tip">
      没有数据
    </div>
  </div>
  <div v-else>
    <el-tabs type="border-card">
      <el-tab-pane
        v-for="(table, index) in data"
        :key="index"
        :label="table.Title">
        <sheet :data="table"/>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import sheet from '@/components/sheet'

import elTabs from 'element-ui/lib/tabs'
import elTabPane from 'element-ui/lib/tab-pane'
import 'element-ui/lib/theme-chalk/tabs.css'
import 'element-ui/lib/theme-chalk/tab-pane.css'

export default {
  name: 'Presenter',
  components: { elTabs, elTabPane, sheet },
  props: {
    data: {
      type: Array,
      required: true
    }
  }
}
</script>
<style scoped>
.non-data-tip {
  padding: 5rem;
  box-shadow: inset 0 0 2px 0 black;
  background-color: #f8f8f8;
}
</style>
