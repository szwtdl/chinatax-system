<template>
  <div class="dashboard-editor-container">
    <panel-group :welcome="panel" @handleSetLineChartData="handleSetLineChartData" />
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
      <line-chart :chart-data="lineChartData" />
    </el-row>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import * as DashBoardApi from '@/api/dashboard'
const lineChartData = {
  newVisitis: {
    userData: [100, 120, 161, 134, 105, 160, 165],
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145],
    agentData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    LineChart
  },
  data() {
    return {
      lineChartData: lineChartData.newVisitis,
      panel: {
        agent_count: 0,
        order_paid: 0,
        proxy_count: 0,
        task_count: 0,
        member_count: 0
      }
    }
  },
  mounted() {
    this.getWelcome()
  },
  methods: {
    getWelcome() {
      DashBoardApi.Welcome().then(response => {
        const data = response.data || {}
        this.panel.agent_count = data.agent_count || 0
        this.panel.order_paid = data.order_unpaid || 0
        this.panel.proxy_count = data.proxy_count || 0
        this.panel.task_count = data.task_count || 0
        this.panel.member_count = data.member_count || 0
        this.lineChartData = {
          userData: [10, 2, 2, 12, 2, 3, 8] || this.lineChartData.userData,
          expectedData: [100, 2, 29, 12, -2, 30, 98] || this.lineChartData.expectedData,
          actualData: [1000, 2, 209, 12, 2, 3, 12] || this.lineChartData.actualData,
          agentData: data.agent_list || this.lineChartData.agentData
        }
      }).catch(() => {
        this.$message.error('获取数据失败，请稍后再试')
      })
    },
    handleSetLineChartData(type) {
      this.lineChartData = lineChartData[type]
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
