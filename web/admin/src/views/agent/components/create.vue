<template>
  <div class="app-container">
    <el-dialog
      title="新增代理"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="账户昵称">
          <el-input v-model="account.nickname" autocomplete="off" style="width: 560px;" />
          <span style="color: red; margin-left: 5px;">*</span>
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="登录账号">
              <el-input v-model="account.username" autocomplete="off" style="width: 220px;" />
              <span style="color: red; margin-left: 5px;">*</span>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="登录密码">
              <el-input v-model="account.password" type="password" autocomplete="off" style="width: 220px;" />
              <span style="color: red; margin-left: 5px;">*</span>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="课程金额">
              <el-input v-model="account.price" autocomplete="off" style="width: 220px;" />
              <span style="color: red; margin-left: 5px;">*</span>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="授信额度">
              <el-input v-model="account.credit" autocomplete="off" style="width: 220px;" />
              <span style="color: red; margin-left: 5px;">*</span>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="邀请编号">
          <el-input v-model="account.invite_code" autocomplete="off" style="width: 220px;" />
          <span style="margin-left: 8px; color: #888; font-size: 13px;">邀请编号：默认12345678</span>
        </el-form-item>
        <el-form-item label="账号状态">
          <el-switch v-model="account.status" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleSubmit">保存</el-button>
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
    }
  },
  data() {
    return {
      account: {
        nickname: '',
        username: '',
        password: '888888',
        invite_code: '12345678',
        price: 50,
        credit: 100,
        status: true
      }
    }
  },
  methods: {
    handleSubmit() {
      // 判断是否填写了必要信息
      if (!this.account.nickname || !this.account.username || !this.account.password || !this.account.invite_code || !this.account.credit) {
        this.$message.error('请填写所有必填项！')
        return
      }
      this.$emit('save', {
        nickname: this.account.nickname,
        username: this.account.username,
        password: this.account.password,
        price: parseFloat(this.account.price) || 0,
        invite_code: this.account.invite_code,
        credit: parseInt(this.account.credit, 10),
        status: this.account.status
      })
      this.$emit('update:visible', false)
    },
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
}

export class createAccount {
}
</script>
<style scoped lang="scss">
::v-deep{
  .el-form-item {
    margin-bottom: 18px!important;
  }
  .el-row{
    margin-bottom: 0!important;
  }
}
</style>
