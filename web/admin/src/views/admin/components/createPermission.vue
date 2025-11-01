<template>
  <div class="app-container">
    <el-dialog
      title="新增权限"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="父级权限">
          <el-select v-model="parent_id" placeholder="请选择父级权限" style="width: 520px;">
            <el-option
              v-for="(item, index) in permissions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="权限名称">
          <el-input v-model="name" autocomplete="off" placeholder="login" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="权限方法">
          <el-checkbox-group v-model="method">
            <el-checkbox
              v-for="(item, index) in actions"
              :key="index"
              :label="item"
            >
              {{ item }}
            </el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="权限路由">
          <el-input v-model="path" autocomplete="off" placeholder="admin/login" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="权限备注">
          <el-input v-model="description" autocomplete="off" placeholder="管理员登录" style="width: 520px;" />
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
    actions: {
      type: Array,
      default: () => []
    },
    permissions: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      parent_id: 0,
      name: '',
      description: '',
      path: '',
      method: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'],
      sort_num: 99
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSubmit() {
      this.$emit('save', {
        parent_id: this.parent_id,
        name: this.name,
        desc: this.description,
        path: this.path,
        method: this.method,
        sort_num: this.sort_num
      })
    }
  }
}
</script>
<style scoped lang="scss">

</style>
