<template>
  <div class="app-container">
    <el-form label="代理设置" label-width="120px">
      <el-form-item label="代理平台" prop="platform">
        <el-select value="无忧代理" style="width: 590px;">
          <el-option label="无忧代理" value="wuyouip" />
        </el-select>
      </el-form-item>
      <el-form-item label="代理地址" prop="address">
        <el-input v-model="address" placeholder="请输入代理地址" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="代理密钥" prop="token">
        <el-input v-model="token" placeholder="请输入代理账号" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="使用数量" prop="quantity_used">
        <el-input v-model="quantity_used" placeholder="一个IP可以绑定的账号数量" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="代理开关" prop="switch">
        <el-switch v-model="switch_status" />
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
      switch_status: '',
      platform: 'wuyouip',
      address: 'http://api.user.wuyouip.com',
      token: 'e124006eaf2946f5b854fc94e74bf60c',
      quantity_used: 5 // 假设这是一个默认值，可以根据实际需求调整
    }
  },
  mounted() {
    getConfigs({ key: 'proxy' }).then((res) => {
      this.switch_status = res.data.switch
      this.platform = res.data.platform
      this.address = res.data.address
      this.token = res.data.token
      this.quantity_used = res.data.quantity_used
    })
  },
  methods: {
    handlerSubmit() {
      updateConfigs({
        key: 'proxy',
        config: {
          switch: this.switch_status,
          platform: this.platform,
          address: this.address,
          token: this.token,
          quantity_used: this.quantity_used
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
