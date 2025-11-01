<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="create = true">添加账号</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="ID"
              width="100"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="登录账号" prop="username" min-width="150" align="left" />
            <el-table-column label="账号昵称" prop="nickname" min-width="180" align="left" />
            <el-table-column label="账号角色" align="left" min-width="200">
              <el-tag style="margin-right: 10px;">普通管理员</el-tag>
              <el-tag>网站编辑</el-tag>
            </el-table-column>
            <el-table-column label="账号状态" align="center" min-width="150">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.status === true" type="success">正常</el-tag>
                <el-tag v-else-if="scope.row.status === false" type="danger">禁用</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" width="250">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="openEditDialog(scope.row)"
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
        </el-card>
      </el-col>
    </el-row>

    <create-account :visible.sync="create" :roles="roles" @save="handleSave" />
    <edit-account :visible.sync="edit" :account="account" :roles="roles" @save="handleEdit" />
  </div>
</template>
<script>

import createAccount from '@/views/admin/components/createAccount.vue'
import editAccount from '@/views/admin/components/editAccount.vue'
import * as ApiAdmin from '@/api/admin'
import * as ApiRole from '@/api/role'

export default {
  name: 'Admin',
  components: { createAccount, editAccount },
  data() {
    return {
      list: [],
      loading: false,
      listQuery: {
        page: 1,
        limit: 10,
        total: 0,
        page_size: 0
      },
      create: false,
      edit: false,
      roles: [],
      account: {
        nickname: '',
        username: '',
        password: '',
        status: 1,
        selectedRoles: ['1']
      }
    }
  },
  mounted() {
    this.getList()
    this.getRoles()
  },
  methods: {
    getRoles() {
      ApiRole.List().then(response => {
        this.roles = response.data.items.map(item => {
          return {
            label: item.Description,
            value: item.id
          }
        })
      })
    },
    getList() {
      this.loading = true
      ApiAdmin.AdminList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery = {
          page: response.data.page,
          limit: response.data.page_size,
          total: response.data.total
        }
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSave(data) {
      ApiAdmin.addAdmin({
        avatar: data.avatar,
        nickname: data.nickname,
        username: data.username,
        password: data.password,
        role_ids: data.role_ids
      }).then(response => {
        this.showDialog = false
        this.getList()
      })
    },
    openEditDialog(item) {
      this.edit = true
      this.account = {
        id: item.id,
        avatar: item.avatar,
        nickname: item.nickname,
        username: item.username,
        status: item.status,
        super: item.super,
        selectedRoles: [],
        roles: []
      }
    },
    handleEdit(data) {
      ApiAdmin.updateAdmin({
        id: data.id,
        avatar: data.avatar,
        nickname: data.nickname,
        password: data.password,
        role_ids: data.role_ids
      }).then(response => {
        this.edit = false
        this.getList()
      })
    },
    handleDelete(item) {
      this.$confirm('此操作将永久删除该管理员账号, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiAdmin.deleteAdmin({ id: item.id }).then(response => {
          this.getList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.getList()
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.listQuery.limit = val
    }
  }
}
</script>
