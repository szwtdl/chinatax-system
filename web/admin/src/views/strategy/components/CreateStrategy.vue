<template>
  <div class="app-container">
    <el-dialog
      title="新增策略"
      :visible.sync="visible"
      max-width="700px"
      :before-close="handlerClose"
    >
      <el-form>
        <el-form-item label="策略平台">
          <el-select v-model="platform_id" class="select-type" size="small" style="width: 520px;" @change="handlePlatformChange">
            <el-option
              v-for="option in options"
              :key="option.id"
              :label="option.name"
              :value="option.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="策略名称">
          <el-input v-model="name" size="small" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="暂停时段">
          <el-switch v-model="is_pause" />
        </el-form-item>
        <div v-if="is_pause" class="pause-list">
          <div v-for="(item, index) in pauses" :key="index" style="margin-bottom: 10px;">
            <el-row :gutter="20">
              <el-col :span="8">
                <el-time-picker
                  v-model="item.start_time"
                  size="small"
                  format="HH:mm"
                  value-format="HH:mm"
                  placeholder="开始"
                  style="width: 100%;"
                />
              </el-col>
              <el-col :span="8">
                <el-time-picker
                  v-model="item.end_time"
                  size="small"
                  format="HH:mm"
                  value-format="HH:mm"
                  placeholder="结束"
                  style="width: 100%;"
                />
              </el-col>
              <el-col :span="6">
                <el-button-group>
                  <el-button type="primary" size="small" icon="el-icon-plus" @click="addPause" />
                  <el-button type="danger" size="small" icon="el-icon-delete" @click="removePause(index)" />
                </el-button-group>
              </el-col>
            </el-row>
          </div>
        </div>
        <el-form-item label="触发频率">
          <el-input-number v-model="interval" size="small" :min="1" style="width: 150px;" />
          <el-select v-model="unit" size="small" placeholder="请选择单位" style="margin-left: 10px; width: 150px;">
            <el-option label="秒" value="second" />
            <el-option label="分钟" value="minute" />
            <el-option label="小时" value="hour" />
            <el-option label="天" value="day" />
          </el-select>
        </el-form-item>
        <el-form-item label="并发数量">
          <el-input v-model="max_concur" class="input-number" size="small" autocomplete="off" style="width: 350px;" />
        </el-form-item>
        <el-form-item label="策略状态">
          <el-switch v-model="is_active" size="small" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handlerClose">取 消</el-button>
        <el-button type="primary" @click="handlerSubmit">保存</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
export default {
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    options: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      platform_id: null,
      name: '全天策略',
      is_pause: false,
      pauses: [
        { start_time: '23:59', end_time: '07:35' } // 默认一行
      ],
      interval: 10,
      max_concur: 10, // 最大并发数
      unit: 'second',
      is_active: true,
      is_all_day: true // 是否全天执行
    }
  },
  watch: {
    options: {
      immediate: true, // 组件初始化就触发一次
      handler(newVal) {
        if (newVal.length > 0 && !this.platform_id) {
          this.platform_id = newVal[0].id
        }
      }
    }
  },
  methods: {
    handlePlatformChange(value) {
      this.platform_id = value
    },
    handlerClose() {
      this.$emit('update:visible', false)
    },
    addPause() {
      if (this.pauses.length >= 5) {
        this.$message.warning('最多5个时间段')
        return
      }
      this.pauses.push({ start_time: '', end_time: '' })
    },
    removePause(index) {
      if (this.pauses.length <= 1) {
        this.$message.warning('至少保留一个暂停时段')
        return
      }
      this.pauses.splice(index, 1)
    },
    handlerSubmit() {
      if (!this.name) {
        this.$message.error('策略名称不能为空')
        return
      }
      if (this.interval <= 0) {
        this.$message.error('执行频率必须大于0')
        return
      }
      if (!this.unit) {
        this.$message.error('请选择执行频率单位')
        return
      }
      this.$emit('save', {
        name: this.name,
        platform_id: this.platform_id,
        is_pause: this.is_pause,
        pauses: this.pauses,
        interval: this.interval,
        max_concur: parseInt(this.max_concur) || 10,
        unit: this.unit,
        is_active: this.is_active,
        is_all_day: this.is_all_day
      })
      this.$emit('update:visible', false)
    }
  }
}
</script>
<style scoped lang="scss">
::v-deep {
  .el-dialog {
    width: 680px;
  }
}
.pause-list{
  margin-left: 68px;
  margin-bottom: 15px;
}
@media only screen and (max-width: 470px) {
  ::v-deep {
    .el-dialog{
      width: 96%;
    }
    .el-select{
      width: 100% !important;
    }
    .el-input{
      width: 100% !important;
    }
    .el-col-8{
      width: 28% !important;
      margin-left: 0!important;
      margin-right: 5px!important;
      padding: 0!important;
    }
    .el-col-6{
      width: 28% !important;
      margin-left: 0!important;
      padding: 0!important;
    }
    .el-select--small{
      width: 30% !important;
    }
    .el-input-number{
      width: 42% !important;
    }
    .el-button--small{
      padding: 9px 12px!important;
    }
  }
  .el-form-item{
    margin-bottom: 0!important;
  }
  .select-type{
    width: 100% !important;
  }
  .input-number{
    width: 76% !important;
    margin-top: 5px;
  }
  .pause-list{
    margin-left: 10px!important;
  }
}
</style>
