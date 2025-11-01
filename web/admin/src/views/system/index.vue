<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <el-tabs v-model="activeTab" @tab-click="handleTabClick">
            <el-tab-pane label="系统设置" name="system">
              <SystemComponents />
            </el-tab-pane>
            <el-tab-pane label="上传配置" name="upload">
              <UploadComponents />
            </el-tab-pane>
          </el-tabs>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import SystemComponents from '@/views/system/components/System.vue'
import UploadComponents from '@/views/system/components/Upload.vue'
import { getConfigs, updateConfigs } from '@/api/config'

export default {
  name: 'System',
  components: { SystemComponents, UploadComponents },
  data() {
    return {
      activeTab: 'system'
    }
  },
  methods: {
    uploadHandle(data) {
      updateConfigs({
        key: 'upload',
        config: {
          type: data.type,
          ext: data.ext,
          size: data.size
        }
      }).then(res => {
        this.$message({
          message: '更新上传',
          type: 'success'
        })
      })
    },
    handleTabClick(tab, event) {
      getConfigs({
        key: this.activeTab
      }).then((res) => {
        switch (this.activeTab) {
          case 'system':
            this.system = res.data
            break
          case 'upload':
            this.upload = res.data
            break
        }
      })
    },
    formatTimeToDate(timeString) {
      const date = new Date() // 当前日期
      const [hours, minutes, seconds] = timeString.split(':').map(Number) // 分割时间字符串
      date.setHours(hours, minutes, seconds, 0) // 设置时间部分
      return date.toString() // 返回格式化字符串
    },
    formatTime(dateString) {
      const date = new Date(dateString)
      if (isNaN(date)) {
        throw new Error('Invalid date format')
      }
      // 补零函数
      const padZero = (num) => String(num).padStart(2, '0')
      // 获取时分秒
      const hours = padZero(date.getHours())
      const minutes = padZero(date.getMinutes())
      const seconds = padZero(date.getSeconds())
      // 返回格式化字符串
      return `${hours}:${minutes}:${seconds}`
    }
  }
}
</script>
