<template>
  <div class="app-container">
    <el-form label="授权令牌" label-width="120px">
      <el-form-item label="授权令牌" name="token">
        <el-input v-model="token" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="任务设置" prop="delivery">
        <el-switch v-model="period.status" />
      </el-form-item>
      <el-form-item v-if="period.status" label="停止任务" style="width: 590px;">
        <el-time-picker
          v-model="period.value"
          is-range
          :clearable="false"
          range-separator="至"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          placeholder="选择时间范围"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getConfigs, updateConfigs } from '@/api/config'

export default {
  data() {
    return {
      token: '',
      period: {
        status: true,
        value: [new Date(2016, 9, 10, 1, 30), new Date(2016, 9, 10, 7, 0)]
      }
    }
  },
  mounted() {
    getConfigs({ key: 'authorize' }).then((res) => {
      this.token = res.data.token
      this.period.status = res.data.period.status
      this.period.value = [
        new Date(this.formatTimeToDate(res.data.period.start)),
        new Date(this.formatTimeToDate(res.data.period.end))
      ]
    })
  },
  methods: {
    submit() {
      updateConfigs({
        key: 'authorize',
        config: {
          token: this.token,
          period: {
            status: this.period.status,
            start: this.formatTime(this.period.value[0]),
            end: this.formatTime(this.period.value[1])
          }
        }
      }).then(res => {
        this.$message({
          message: '更新授权',
          type: 'success'
        })
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
<style scoped lang="scss">

</style>
