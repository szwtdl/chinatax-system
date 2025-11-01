<template>
  <div class="app-container">
    <el-dialog
      title="编辑管理员"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form v-model="account">
        <el-form-item label="账号头像">
          <el-avatar :size="40" shape="square" :src="account.avatar" />
        </el-form-item>
        <el-form-item label="账户昵称">
          <el-input v-model="account.nickname" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="登录账号">
          <el-input v-model="account.username" readonly autocomplete="off" style="width: 520px;" />
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
        <el-button type="primary" @click="handleSave">保存</el-button>
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
    account: {
      type: Object, // 类型为对象
      default: () => ({}) // 默认值为空对象
    },
    roles: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    account: {
      handler(val) {
        this.account = val
      },
      deep: true // 监听对象内部的变化
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSave() {
      this.$emit('save', {
        id: this.account.id,
        avatar: this.account.avatar,
        nickname: this.account.nickname,
        username: this.account.username,
        password: this.account.password,
        role_ids: this.account.selectedRoles
      })
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
