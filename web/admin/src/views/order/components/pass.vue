<template>
  <div class="app-container">
    <el-dialog
      title="请输入授权密码"
      :visible.sync="visible"
      width="430px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="授权密码">
          <el-input v-model="password" type="password" autocomplete="off" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleSubmit">验证密码</el-button>
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
      password: ''
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSubmit() {
      this.$emit('submit', {
        password: this.password
      })
      this.handleClose()
    }
  }
}
</script>
<style scoped lang="scss">

</style>
