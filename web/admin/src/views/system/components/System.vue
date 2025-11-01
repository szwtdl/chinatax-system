<template>
  <div class="app-container">
    <el-form label="系统配置" label-width="120px">
      <el-form-item label="公司名称" name="name">
        <el-input v-model="name" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="公司地址" name="address">
        <el-input v-model="address" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="联系姓名" name="contact">
        <el-input v-model="contact" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="公司电话" name="tel">
        <el-input v-model="tel" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="系统状态" prop="switch">
        <el-switch v-model="switch_status" />
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
      name: '',
      address: '',
      contact: '',
      tel: '',
      switch_status: false
    }
  },
  mounted() {
    getConfigs({ key: 'system' }).then((res) => {
      this.name = res.data.name
      this.address = res.data.address
      this.contact = res.data.contact
      this.tel = res.data.tel
      this.switch_status = res.data.switch
    })
  },
  methods: {
    submit() {
      updateConfigs({
        key: 'system',
        config: {
          name: this.name,
          address: this.address,
          contact: this.contact,
          tel: this.tel,
          switch: this.switch_status
        }
      }).then(res => {
        this.$message({
          message: '保存成功',
          type: 'success'
        })
      })
    }
  }
}
</script>
<style scoped lang="scss">

</style>
