<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入内容"
              size="small"
              class="input-with-select"
              style="margin-right: 5px;"
            >
              <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
            </el-input>
            <el-button-group>
              <el-button type="primary" size="small" @click="visibleRole = true">添加</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="items" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="ID"
              prop="id"
              width="120"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="角色名称" prop="name" align="left" />
            <el-table-column label="角色备注" prop="Description" align="left" />
            <el-table-column label="创建时间" prop="created_at" align="center" />
            <el-table-column label="更新时间" prop="updated_at" align="center" />
            <el-table-column align="center" label="操作" width="250">
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
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="listQuery.total"
            :current-page.sync="listQuery.page"
            :page-size.sync="listQuery.limit"
            style="margin-top: 20px; text-align: center;"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
          />
          <create-role :visible.sync="visibleRole" :permissions="permissions" @save="handleSave" />
          <edit-role :visible.sync="visibleEdit" :data="form" :permissions="permissions" @save="handleUpdate" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import CreateRole from './components/createRole.vue'
import EditRole from './components/editRole.vue'
import * as ApiRole from '@/api/role'

export default {
  name: 'Role',
  components: { CreateRole, EditRole },
  data() {
    return {
      loading: false,
      visibleRole: false,
      visibleEdit: false,
      items: [],
      permissions: [],
      form: {
        name: '',
        description: '',
        selectedRoles: []
      },
      listQuery: {
        page: 1,
        limit: 20,
        total: 0
      }
    }
  },
  created() {
    this.fetchData()
    this.fetchListPermission()
  },
  methods: {
    fetchData() {
      this.loading = true
      ApiRole.List(this.listQuery).then(response => {
        this.items = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    fetchListPermission() {
      ApiRole.permissions().then(response => {
        this.permissions = response.data
      })
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.fetchData()
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.fetchData()
    },
    handleEdit(data) {
      this.visibleEdit = true
      this.form = {
        ...data,
        selectedRoles: []
      }
    },
    handleSearch() {
      console.log('search', this.listQuery.keywords)
    },
    handleDelete(data) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiRole.Delete(data.id).then(response => {
          this.fetchData()
        })
      }).catch(() => {
      })
    },
    handleSave(item) {
      this.visibleRole = true
      ApiRole.Create(item).then(response => {
        this.visibleRole = false
        this.fetchData()
      })
    },
    handleUpdate(item) {
      ApiRole.Update({
        id: item.id,
        name: item.name,
        description: item.description,
        permissions: item.selectedRoles
      }).then(response => {
        this.visibleEdit = false
        this.fetchData()
      })
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
