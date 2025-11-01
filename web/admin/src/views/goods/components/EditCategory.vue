<template>
  <div class="app-container">
    <el-dialog
      title="编辑分类"
      :visible.sync="visible"
      width="25%"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="父级分类">
          <el-select v-model="data.parent_id" value="0" placeholder="请选择父级分类" style="width: 520px;">
            <el-option label="一级分类" value="0" />
            <el-option label="二级分类" value="1" />
            <el-option label="三级分类" value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类名称">
          <el-input v-model="data.name" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="分类图片">
          <el-input v-model="data.image" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="分类排序">
          <el-input v-model="data.sort_order" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="分类状态">
          <el-switch v-model="data.status" />
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
    data: {
      type: Object,
      default: () => ({})
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
    handleSave() {
      this.$emit('save', this.data)
    }
  }
}
</script>
