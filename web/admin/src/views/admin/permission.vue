<template>
  <div class="app-container">
    <el-card>
      <el-row>
        <el-col :span="24">
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="create = true">添加权限</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="items" size="mini" border style="width: 100%">
            <el-table-column prop="id" label="ID" width="120px" />
            <el-table-column prop="description" label="权限名称" min-width="150px">
              <template slot-scope="scope">
                {{ scope.row.prefix }}{{ scope.row.description }}
              </template>
            </el-table-column>
            <el-table-column prop="name" label="权限别名" min-width="150px" />
            <el-table-column prop="path" label="权限路径" min-width="150px" />
            <el-table-column prop="created_at" label="创建时间" min-width="150px" />
            <el-table-column prop="updated_at" label="更新时间" min-width="150px" />
            <el-table-column align="center" label="操作" fixed="right" width="180px">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="handleEdit(scope.row)"
                  >编辑
                  </el-button>
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handleDelete(scope.row)"
                  >删除
                  </el-button>
                </el-button-group>
              </template>
            </el-table-column>
          </el-table>
        </el-col>
      </el-row>
      <create-permission :visible.sync="create" :actions="actions" :permissions="permissions" @save="handleSave" />
      <edit-permission :visible.sync="edit" :data="form" :actions="actions" :permissions="permissions" @submit="handleUpdate" />
    </el-card>
  </div>
</template>
<script>
import createPermission from '@/views/admin/components/createPermission.vue'
import editPermission from '@/views/admin/components/editPermission.vue'

import * as ApiPermission from '@/api/permission'

export default {
  name: 'Permission',
  components: {
    createPermission,
    editPermission
  },
  data() {
    return {
      items: [],
      loading: false,
      create: false,
      edit: false,
      listQuery: {
        page: 1,
        limit: 20,
        total: 0
      },
      actions: [
        'GET',
        'POST',
        'PUT',
        'DELETE',
        'PATCH'
      ],
      form: {
        parent_id: 0,
        name: '',
        description: '',
        path: '',
        method: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'],
        sort_num: 99
      },
      permissions: []
    }
  },
  created() {
    this.getItems()
  },
  methods: {
    getItems() {
      this.loading = true
      ApiPermission.List(this.listQuery).then(response => {
        this.items = response.data
        this.permissions = [{ label: '顶级权限', value: 0 }, ...response.data.map(item => ({
          value: item.id,
          label: item.prefix + item.description
        }))]
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleEdit(row) {
      this.edit = true
      this.form = row
    },
    handleSave(data) {
      this.loading = true
      ApiPermission.Create(data).then(() => {
        this.getItems()
        this.$message.success('添加成功')
        this.create = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleUpdate(data) {
      ApiPermission.Update(data).then(() => {
        this.getItems()
        this.edit = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleDelete(row) {
      // 弹出确认框
      this.$confirm('此操作将永久删除该权限, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiPermission.Delete({ id: row.id }).then(() => {
          this.getItems()
        }).finally(() => {
          this.loading = false
        })
      }).catch(() => {
      })
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.getItems()
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getItems()
    }
  }
}
</script>
<style scoped lang="scss">

</style>
