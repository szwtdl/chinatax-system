<template>
  <div class="app-container">
    <el-form label="App配置" label-width="120px">
      <el-form-item label="app名称" name="name">
        <el-input v-model="name" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="升级描述" name="release_notes">
        <el-input v-model="release_notes" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="安卓连接" name="update_android_url">
        <el-input v-model="update_android_url" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="苹果连接" name="update_ios_url">
        <el-input v-model="update_ios_url" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="升级版本" prop="latest_version">
        <el-input v-model="latest_version" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="最小版本" prop="min_supported_version">
        <el-input v-model="min_supported_version" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="强制升级" name="is_force_update">
        <el-switch v-model="is_force_update" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handlerSubmit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getConfigs, updateConfigs } from '@/api/config'

export default {
  data() {
    return {
      name: '',
      release_notes: '',
      update_android_url: '',
      update_ios_url: '',
      is_force_update: false,
      latest_version: '',
      min_supported_version: ''
    }
  },
  mounted() {
    getConfigs({ key: 'app' }).then((res) => {
      this.name = res.data.name
      this.release_notes = res.data.release_notes
      this.update_android_url = res.data.update_android_url
      this.update_ios_url = res.data.update_ios_url
      this.is_force_update = res.data.is_force_update
      this.latest_version = res.data.latest_version
      this.min_supported_version = res.data.min_supported_version
    })
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handlerSubmit() {
      updateConfigs({
        key: 'app',
        config: {
          name: this.name,
          release_notes: this.release_notes,
          update_android_url: this.update_android_url,
          update_ios_url: this.update_ios_url,
          is_force_update: this.is_force_update,
          latest_version: this.latest_version,
          min_supported_version: this.min_supported_version
        }
      }).then(res => {
        this.$message({
          message: '更新app配置',
          type: 'success'
        })
      })
    }
  }
}
</script>
<style scoped lang="scss">

</style>
