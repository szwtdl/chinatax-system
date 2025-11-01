<template>
  <div class="app-container">
    <el-dialog
      title="编辑代理"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="账户昵称">
          <el-input v-model="account.nickname" autocomplete="off" style="width: 560px;" />
        </el-form-item>
        <el-row :gutter="20" style="margin-bottom: 10px;">
          <el-col :span="12">
            <el-form-item label="登录账号">
              <el-input v-model="account.username" readonly autocomplete="off" style="width: 220px;" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="登录密码">
              <el-input v-model="account.password" type="password" placeholder="留空不修改密码" autocomplete="off" style="width: 220px;" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20" style="margin-bottom: 10px;">
          <el-col :span="12">
            <el-form-item label="课程金额">
              <el-input v-model="account.price" placeholder="课程金额" autocomplete="off" style="width: 220px;" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="授信额度">
              <el-input v-model="account.credit" placeholder="授信额度数量" autocomplete="off" style="width: 220px;" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="邀请编号">
          <el-input v-model="account.invite_code" placeholder="邀请编号：默认888888" autocomplete="off" style="width: 220px;" />
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
    },
    data: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      account: {
        nickname: '',
        username: '',
        password: '',
        price: 0,
        invite_code: '',
        credit: 0,
        status: true
      }
    }
  },
  watch: {
    visible(newVal) {
      if (newVal) {
        // 兼容 this.data 为空或字段缺失的情况，避免 undefined
        this.account = Object.assign({
          nickname: '',
          username: '',
          password: '',
          invite_code: '',
          price: 0,
          credit: 0,
          status: true
        }, this.data || {})
      }
    }
  },
  methods: {
    handleSubmit() {
      if (!this.account.nickname || !this.account.username || !this.account.invite_code || !this.account.credit) {
        this.$message.error('请填写完整的账户信息')
        return
      }
      this.$emit('save', {
        'id': this.data.id || null,
        'nickname': this.account.nickname,
        'username': this.account.username,
        'password': this.account.password,
        'price': parseFloat(this.account.price) || 0,
        'invite_code': this.account.invite_code || '888888',
        'credit': parseInt(this.account.credit, 10) || 0,
        'status': this.account.status
      })
      this.$emit('update:visible', false)
    },
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
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
