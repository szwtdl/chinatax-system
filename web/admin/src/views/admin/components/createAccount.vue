<template>
  <div class="app-container">
    <el-dialog
      title="新增管理员"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="账号头像">
          <el-avatar :size="40" shape="square" :src="account.avatar" />
        </el-form-item>
        <el-form-item label="账户昵称">
          <el-input v-model="account.nickname" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="登录账号">
          <el-input v-model="account.username" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="登录密码">
          <el-input v-model="account.password" type="password" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="账号角色">
          <div class="el-checkbox-group">
            <el-checkbox-group v-model="account.selectedRoles">
              <el-checkbox
                v-for="(item, index) in roles"
                :key="index"
                :label="item.value"
              >
                {{ item.label }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
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
    roles: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      account: {
        avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
        nickname: '',
        username: '',
        password: '',
        status: true,
        selectedRoles: []
      }
    }
  },
  methods: {
    handleSubmit() {
      const reg = /^[a-zA-Z][a-zA-Z0-9]{3,19}$/
      if (!reg.test(this.account.username)) {
        this.$message.error('登录账号必顫字母开头，后面可以是字母或数字，长度为6-20位')
        return
      }
      if (this.account.password.length < 6) {
        this.$message.error('登录密码长度不能小于6位')
        return
      }
      this.$emit('save', {
        avatar: this.account.avatar,
        nickname: this.account.nickname,
        username: this.account.username,
        password: this.account.password,
        role_ids: this.account.selectedRoles
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
.el-checkbox-group {
  width: 520px;
  float: left;
}
</style>
