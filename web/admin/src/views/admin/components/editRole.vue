<template>
  <div class="app-container">
    <el-dialog
      title="编辑角色"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="角色名称">
          <el-input v-model="data.name" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="角色备注">
          <el-input v-model="data.Description" autocomplete="off" style="width: 520px;" />
        </el-form-item>
        <el-form-item label="角色权限">
          <el-checkbox-group v-model="selectedPermissions" style="width: 590px; display: block; float: right;">
            <div v-for="parent in permissions" :key="parent.id">
              <!-- 父级 -->
              <el-checkbox
                v-model="parentCheckedMap[parent.id]"
                :label="parent.id"
                :indeterminate="indeterminateMap[parent.id]"
                style="font-weight: bold; display: block;"
                @change="toggleParent(parent)"
              >
                {{ parent.description }}
              </el-checkbox>
              <!-- 子级 -->
              <div style="padding-left: 20px;">
                <el-checkbox
                  v-for="child in parent.children || []"
                  :key="child.id"
                  :label="child.id"
                >
                  {{ child.description }}
                </el-checkbox>
              </div>
            </div>
          </el-checkbox-group>
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
      type: Object, // 类型为对象
      default: () => ({}) // 默认值为空对象
    },
    permissions: {
      type: Array,
      default: () => [] // 默认值为空数组
    }
  },
  data() {
    return {
      selectedPermissions: [],
      parentCheckedMap: {},
      indeterminateMap: {}
    }
  },
  watch: {
    data: {
      handler(val) {
        this.data = val
      },
      deep: true // 监听对象内部的变化
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSave() {
      this.$emit('save', this.data)
      this.$emit('update:visible', false)
    },
    updateParentStates() {
      this.permissions.forEach(parent => {
        const childIds = (parent.children || []).map(c => c.id)
        const checkedCount = childIds.filter(id => this.selectedPermissions.includes(id)).length
        this.parentCheckedMap[parent.id] = checkedCount === childIds.length
        this.indeterminateMap[parent.id] = checkedCount > 0 && checkedCount < childIds.length
      })
    },
    toggleParent(parent) {
      const childIds = (parent.children || []).map(c => c.id)
      if (this.parentCheckedMap[parent.id]) {
        this.selectedPermissions = Array.from(new Set([...this.selectedPermissions, ...childIds]))
      } else {
        this.selectedPermissions = this.selectedPermissions.filter(id => !childIds.includes(id))
      }
      this.updateParentStates()
    }
  }
}
</script>
<style scoped lang="scss">
.el-checkbox-group {
  display: flex;
  flex-wrap: wrap;

  .el-checkbox {
    margin-right: 15px;
  }
}
</style>
