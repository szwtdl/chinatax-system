<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="show">添加机构</el-button>
            </el-button-group>
          </div>
          <el-table :data="list" size="mini" show-header border style="width: 100%">
            <el-table-column type="selection" align="center" min-width="80px" />
            <el-table-column prop="org_id" align="left" label="机构ID" min-width="120px" element-loading-text="请给我点时间！" />
            <el-table-column prop="username" label="机构账号" min-width="150px" align="left" />
            <el-table-column prop="name" label="机构名称" min-width="180px" align="left" />
            <el-table-column label="机构状态" align="center" min-width="100px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.status"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" min-width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" min-width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" min-width="150px">
              <template slot-scope="scope">
                <el-dropdown @command="(command) => handleCommand(scope.row, scope.$index, command)">
                  <el-button size="mini" type="primary">操作 <i class="el-icon-arrow-down el-icon--right" />
                  </el-button>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item command="import">导入学员</el-dropdown-item>
                    <el-dropdown-item command="edit">编辑账号</el-dropdown-item>
                    <el-dropdown-item command="delete">删除账号</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog
      title="新增机构"
      :visible.sync="dialogVisible"
      width="650px"
      :before-close="handleClose"
    >
      <el-form :model="form">
        <el-form-item label="机构名称" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录账号" :label-width="formLabelWidth">
          <el-input v-model="form.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" autocomplete="off" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submitHandler">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="编辑机构"
      :visible.sync="dialogEditVisible"
      width="650px"
      :before-close="handleClose"
    >
      <el-form :model="edit">
        <el-form-item label="机构名称" :label-width="formLabelWidth">
          <el-input v-model="edit.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录密码" :label-width="formLabelWidth">
          <el-input v-model="edit.password" autocomplete="off" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="editHandler">确 定</el-button>
      </span>
    </el-dialog>
    <org-member
      :member="member"
      :visible.sync="member.dialogVisible"
      @change="handleChange"
      @page="handlePage"
      @search="handleSearch"
      @submit="handleUsersSubmit"
    />
  </div>
</template>
<script>
import OrgMember from '@/views/organization/components/Member.vue'
import { addOrg, getMemberList, addMembers, fetchOrgList, updateOrg, deleteOrg } from '@/api/org'

export default {
  name: 'Organization',
  components: { OrgMember },
  data() {
    return {
      loading: false,
      dialogVisible: false,
      dialogEditVisible: false,
      listQuery: {
        page: 1,
        limit: 20
      },
      member: {
        items: [],
        title: '',
        loading: true,
        dialogVisible: false,
        listQuery: {
          keywords: '',
          filter: '',
          page: 1,
          pageSize: 20
        }
      },
      org: {},
      formLabelWidth: '120px',
      list: [],
      selection: [],
      form: {
        name: '',
        username: '',
        password: ''
      },
      edit: {
        id: null,
        name: '',
        password: '',
        status: true
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      fetchOrgList(this.listQuery).then(response => {
        this.list = response.data.items
        this.loading = false
      })
    },
    cancel() {
      this.dialogVisible = false
      this.dialogEditVisible = false
    },
    editHandler() {
      updateOrg(this.edit).then(response => {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.dialogEditVisible = false
        this.getList()
      })
    },
    handleCommand(row, index, command) {
      switch (command) {
        case 'import':
          this.org = row
          this.member.title = row.name
          this.member.dialogVisible = true
          getMemberList({
            org_id: row.org_id,
            page: 1
          }).then(response => {
            this.member.listQuery.page = response.data.page
            this.member.listQuery.pageSize = response.data.page_size
            this.member.listQuery.total = parseInt(response.data.total) * parseInt(response.data.page_size)
            this.member.items = response.data.items
            this.member.loading = false
          })
          break
        case 'edit':
          this.dialogEditVisible = true
          var { id, name, password, status } = row
          this.edit = { id, name, password, status }
          break
        case 'delete':
          this.$confirm('此操作将永久删除该, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            deleteOrg({
              id: row.id
            }).then(response => {
              this.$message({
                type: 'success',
                message: '删除成功'
              })
              this.getList()
            })
          })
      }
    },
    submitHandler() {
      if (this.form.username === '' && this.form.password === '') {
        return this.$message({
          type: 'warning',
          message: '账号或密码不为空'
        })
      }
      if (this.form.name === '') {
        return this.$message({
          type: 'warning',
          message: '机构名称不为空'
        })
      }
      addOrg(this.form).then(response => {
        this.dialogVisible = false
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        // 刷新列表
        this.getList()
      })
    },
    handleClose() {
      this.dialogEditVisible = false
      this.dialogVisible = false
    },
    show() {
      this.dialogVisible = true
    },
    handlePage(val) {
      getMemberList({
        org_id: this.org.org_id,
        keywords: this.member.listQuery.keywords,
        page: val
      }).then(response => {
        this.member.listQuery.page = response.data.page
        this.member.listQuery.pageSize = response.data.page_size
        this.member.listQuery.total = parseInt(response.data.total) * parseInt(response.data.page_size)
        this.member.items = response.data.items
        this.member.loading = false
      })
    },
    handleChange(data) {
      this.selection = []
      data.forEach(item => {
        this.selection.push({
          nickname: item.nickname,
          username: item.username,
          password: 'Dp111111'
        })
      })
    },
    handleSearch() {
      getMemberList({
        org_id: this.org.org_id,
        keywords: this.member.listQuery.keywords,
        page: 1
      }).then(response => {
        this.member.listQuery.page = response.data.page
        this.member.listQuery.pageSize = response.data.page_size
        this.member.listQuery.total = parseInt(response.data.total) * parseInt(response.data.page_size)
        this.member.items = response.data.items
        this.member.loading = false
      })
    },
    handleUsersSubmit() {
      if (this.selection.length === 0) {
        this.$message({
          type: 'warning',
          message: '必须选择用户'
        })
        return false
      }
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      addMembers({
        org_id: this.org.org_id,
        users: this.selection
      }).then(response => {
        loading.close()
        this.$message({
          type: 'success',
          message: '导入成功'
        })
      }).finally(() => {
        loading.close()
      })
    }
  }
}
</script>
