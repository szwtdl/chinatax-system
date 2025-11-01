<template>
  <div class="app-container">
    <el-dialog
      title="编辑权限"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="父级权限">
          <el-select v-model="data.parent_id" placeholder="请选择父级权限" style="width: 520px;">
            <el-option
              v-for="(item, index) in permissions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="权限名称">
          <el-input v-model="data.name" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="权限方法">
          <el-checkbox-group v-model="data.method">
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
          <el-input v-model="data.path" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="权限备注">
          <el-input v-model="data.description" autocomplete="off" style="width: 520px;" />
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
    data: {
      type: Object,
      default: () => ({})
    },
    permissions: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    data: {
      handler(val) {
        this.data = val
      },
      deep: true
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSubmit() {
      this.$emit('submit', this.data)
    }
  }
}
</script>
<style scoped lang="scss">

</style>
