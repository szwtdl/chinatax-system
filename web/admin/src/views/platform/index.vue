<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="dialogVisible = true">增加平台</el-button>
              <el-button type="danger" size="small" @click="handleRefresh">刷新授权</el-button>
            </el-button-group>
          </div>
          <el-table
            :data="items"
            size="mini"
            border
            style="width: 100%"
            @selection-change="handleSelectChange"
          >
            <el-table-column type="selection" align="center" width="60px" />
            <el-table-column
              align="left"
              label="平台ID"
              min-width="100px"
              sortable
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="平台名称" prop="name" min-width="160px" align="left" />
            <el-table-column label="平台AppId" prop="app_key" min-width="120px" align="left" />
            <el-table-column label="平台AppSecret" prop="app_secret" min-width="120px" align="left" />
            <el-table-column label="平台域名" prop="base_url" min-width="260px" align="left" />
            <el-table-column label="平台状态" align="center" min-width="100px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.status"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="是否调试" align="center" min-width="100px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.debug"
                  active-color="#13ce66"
                  inactive-color="#cccccc"
                  @change="handleDebugChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="授权到期" prop="ExpireTime" min-width="160px" align="center">
              <template slot-scope="scope">
                <el-tag
                  :type="getExpireTagType(scope.row)"
                >
                  {{ scope.row.ExpireTime }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="平台排序" prop="sort_num" min-width="100px" align="center">
              <template slot-scope="scope">
                <el-input
                  v-model="scope.row.sort_num"
                  placeholder="请输入库存"
                  size="small"
                  :min="0"
                  type="text"
                  style="width: 50px; text-align: center!important;"
                  @input="handleInputChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" min-width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" width="200">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="success"
                    @click="handleEdit(scope.$index,scope.row)"
                  >编辑
                  </el-button>
                  <el-dropdown @command="(command) => handleCommand(scope.row, scope.$index, command)">
                    <el-button size="mini" type="primary">操作<i class="el-icon-arrow-down el-icon--right" />
                    </el-button>
                    <el-dropdown-menu>
                      <el-dropdown-item command="auth">授权验证</el-dropdown-item>
                      <el-dropdown-item command="delete">删除账号</el-dropdown-item>
                    </el-dropdown-menu>
                  </el-dropdown>
                </el-button-group>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            :current-page.sync="listQuery.page"
            :page-size.sync="listQuery.limit"
            style="margin-top: 20px; text-align: center;"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
          />
        </el-card>
      </el-col>
    </el-row>
    <el-dialog
      title="增加平台"
      :visible.sync="dialogVisible"
      max-width="550px"
      :before-close="handleClose"
    >
      <el-form ref="createForm" :model="form" :rules="formRules">
        <el-form-item label="平台名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="form.name" placeholder="测试平台" autocomplete="off" />
        </el-form-item>
        <el-form-item label="平台编码" :label-width="formLabelWidth" prop="code">
          <el-input v-model="form.code" placeholder="test" autocomplete="off" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="16">
            <el-form-item label="平台域名" :label-width="formLabelWidth" prop="domain">
              <el-input v-model="form.domain" placeholder="https://www.example.com" autocomplete="off" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item prop="sort_num">
              <el-input v-model="form.sort_num" placeholder="平台排序" autocomplete="off" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="编辑平台"
      :visible.sync="editVisible"
      max-width="550px"
      :before-close="handleClose"
    >
      <el-form
        ref="editForm"
        :model="edit"
        :rules="formRules"
        label-position="left"
      >
        <el-form-item label="平台名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="edit.name" placeholder="测试平台" autocomplete="off" />
        </el-form-item>
        <el-form-item label="平台编码" :label-width="formLabelWidth" prop="code">
          <el-input v-model="edit.code" placeholder="test" disabled autocomplete="off" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="16">
            <el-form-item label="平台域名" :label-width="formLabelWidth" prop="domain">
              <el-input v-model="edit.domain" placeholder="https://www.example.com" autocomplete="off" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item prop="sort_num">
              <el-input v-model="edit.sort_num" placeholder="平台排序" autocomplete="off" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="editForm">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="输入授权密码"
      :visible.sync="passwordDialogVisible"
      width="450px"
      :close-on-click-modal="false"
      :before-close="handleDialogClose"
    >
      <el-form ref="passwordForm" :model="passwordForm" label-width="100px">
        <el-form-item
          label="授权密码"
          prop="password"
          :rules="[
            { required: true, message: '请输入授权密码', trigger: 'blur' },
            { min: 8, message: '密码长度不能少于8位', trigger: 'blur' },
            { max: 20, message: '密码长度不能超过20位', trigger: 'blur' },
            { pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d@$!%*#?&]{8,}$/, message: '密码必须包含字母和数字', trigger: 'blur' }
          ]"
        >
          <el-input
            v-model="passwordForm.password"
            type="password"
            placeholder="请输入授权密码以继续操作"
            @keyup.enter.native="confirmPassword"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="handleDialogClose">取消</el-button>
        <el-button type="primary" @click="confirmPassword">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import * as ApiPlatForm from '@/api/platform'

export default {
  name: 'Platform',
  data() {
    const validateDomain = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('请输入平台域名'))
      }
      // 验证必须包含 http:// 或 https:// 前缀
      if (!/^https?:\/\//.test(value)) {
        return callback(new Error('域名必须包含 http:// 或 https:// 前缀'))
      }
      // 验证末尾不能有 /
      if (value.endsWith('/')) {
        return callback(new Error('域名末尾不能包含斜杠 /'))
      }
      // 提取域名部分（去除前缀）
      const pureDomain = value.replace(/^https?:\/\//, '')
      // 域名主体验证正则（支持多级域名、端口号）
      const domainReg = /^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](\.[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])*(:\d{1,5})?$/
      if (domainReg.test(pureDomain)) {
        callback()
      } else {
        callback(new Error('请输入有效的域名（如 https://example.com 或 http://api.example.com:8080）'))
      }
    }
    return {
      loading: false,
      listLoading: true,
      dialogVisible: false,
      editVisible: false,
      listQuery: {
        page: 1,
        limit: 20,
        school_id: 0,
        keywords: '',
        filter: 'username'
      },
      passwordDialogVisible: false,
      passwordForm: {
        password: ''
      },
      currentRow: null,
      originalDebugState: null,
      total: 0,
      total_page: 0,
      pageSize: 0,
      items: [],
      select: [],
      form: {
        name: '',
        code: '',
        domain: '',
        sort_num: '99'
      },
      edit: {
        id: 0,
        name: '',
        code: '',
        domain: '',
        sort_num: '99'
      },
      formLabelWidth: '80px',
      formRules: {
        name: [
          { required: true, message: '请输入平台名称', trigger: 'blur' },
          { min: 2, max: 50, message: '平台名称长度在 2-50 个字符之间', trigger: 'blur' },
          { pattern: /^[\u4e00-\u9fa5a-zA-Z0-9_【】]+$/, message: '平台名称只能包含中文、字母、数字、下划线和【】', trigger: 'blur' }
        ],
        code: [
          { required: true, message: '平台编码不能为空', trigger: 'blur' },
          { pattern: /^[a-zA-Z0-9_]+$/, message: '平台编码只能包含字母、数字和下划线', trigger: 'blur' },
          { min: 2, max: 30, message: '平台编码长度在 2-30 个字符之间', trigger: 'blur' }
        ],
        domain: [
          { required: true, message: '请输入平台域名', trigger: 'blur' },
          { validator: validateDomain, trigger: 'blur' }
        ],
        sort_num: [
          { required: true, message: '请输入排序号', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    getExpireTagType(row) {
      if (row.ExpireTime === '永久授权') {
        return 'success' // 绿色
      }
      // 判断是否过期
      const expireTime = new Date(row.ExpireTime)
      const now = new Date()
      return expireTime < now ? 'danger' : 'success'
    },
    handleStatusChange(row) {
      this.listLoading = true
      const status = row.status ? 'true' : 'false'
      ApiPlatForm.Status({ id: row.id, status }).then(() => {
        this.$message({
          message: '操作成功',
          type: 'success'
        })
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleDebugChange(row) {
      this.currentRow = row
      this.originalDebugState = !row.debug
      this.passwordDialogVisible = true
    },
    handleInputChange(row) {
      this.listLoading = true
      ApiPlatForm.Sort({ id: row.id, sort_num: parseInt(row.sort_num) }).then(() => {
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleRefresh() {
      if (this.select.length === 0) {
        this.$message({
          message: '请选择要刷新的平台',
          type: 'warning'
        })
        return
      }
      this.listLoading = true
      ApiPlatForm.Refresh({
        id: this.select.length > 0 ? this.select.map(it => it.id) : []
      }).then(() => {
        this.$message({
          message: '刷新授权成功',
          type: 'success'
        })
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleDialogClose() {
      if (this.currentRow) {
        this.currentRow.debug = this.originalDebugState
      }
      this.passwordDialogVisible = false
      this.passwordForm.password = ''
      this.currentRow = null
      this.originalDebugState = null
    },
    confirmPassword() {
      this.$refs.passwordForm.validate((valid) => {
        if (valid) {
          // 密码验证通过，执行调试状态切换
          this.listLoading = true
          const debug = this.currentRow.debug ? 'true' : 'false'
          ApiPlatForm.Debug({
            id: this.currentRow.id,
            debug,
            password: this.passwordForm.password
          }).then(() => {
            this.$message({
              message: '调试成功',
              type: 'success'
            })
          }).finally(() => {
            this.listLoading = false
            this.passwordDialogVisible = false
            this.passwordForm.password = '' // 清空密码
            this.fetchData() // 重新加载数据
          })
        }
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.fetchData()
    },
    handlePageChange(val) {
      this.listQuery.page = val
      this.fetchData()
    },
    handleSearch() {
      this.listQuery.page = 1
      this.fetchData()
    },
    handleSelectChange(val) {
      this.select = val
    },
    handleDelete() {
      if (this.select.length === 0) {
        this.$message({
          message: '请选择要删除的平台',
          type: 'warning'
        })
        return
      }
      this.$confirm('是否确认删除选中平台？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiPlatForm.Delete({ id: this.select.map(it => it.id) }).then(() => {
          this.$message({
            message: '删除成功',
            type: 'success'
          })
          this.fetchData()
        }).finally(() => {
          this.select = []
          this.listLoading = false
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    fetchData() {
      ApiPlatForm.List(this.listQuery).then(response => {
        const { data } = response
        this.items = data.items
        this.total = data.total
        this.total_page = data.total_page
        this.pageSize = data.limit
        this.listLoading = false
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleCommand(row, index, command) {
      this.listLoading = true
      if (command === 'auth') {
        ApiPlatForm.Auth({ id: row.id }).then(response => {
          this.listLoading = false
          const { data } = response
          this.$message({
            message: `授权成功，有效期至 ${data.ExpireTime}`,
            type: 'success'
          })
          this.fetchData()
        }).finally(() => {
          this.listLoading = false
          this.fetchData()
        })
      } else if (command === 'delete') {
        this.$confirm('是否确认删除平台？', '提示', {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          ApiPlatForm.Delete({ id: [row.id] }).then(() => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.fetchData()
          }).finally(() => {
            this.listLoading = false
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      }
    },
    resetForm() {
      this.form = {
        name: '',
        code: '',
        domain: '',
        sort_num: '99',
        status: true,
        debug: false
      }
    },
    handleClose() {
      this.dialogVisible = false
      this.editVisible = false
      this.resetForm()
    },
    handleEdit(index, row) {
      this.edit.id = row.id
      this.edit.name = row.name
      this.edit.code = row.code
      this.edit.domain = row.base_url
      this.edit.sort_num = row.sort_num.toString()
      this.editVisible = true
      this.listLoading = false
    },
    cancel() {
      this.dialogVisible = false
      this.editVisible = false
      this.resetForm()
    },
    editForm() {
      this.$refs.editForm.validate((valid) => {
        if (valid) {
          ApiPlatForm.Update(this.edit).then(() => {
            this.$message({
              message: '操作成功',
              type: 'success'
            })
          }).finally(() => {
            this.listLoading = false
            this.editVisible = false
            this.resetForm()
            this.fetchData()
          })
        } else {
          this.$message.warning('请完善表单信息')
          return false
        }
      })
    },
    submit() {
      this.$refs.createForm.validate((valid) => {
        if (valid) {
          this.listLoading = true
          ApiPlatForm.Create(this.form).then(response => {
            this.$message({
              message: '操作成功',
              type: 'success'
            })
          }).finally(() => {
            this.dialogVisible = false
            this.listLoading = false
            this.resetForm()
            this.fetchData()
          })
        } else {
          this.$message.warning('请完善表单信息')
          return false
        }
      })
    }
  }
}
</script>
<style lang="scss" scoped>
  .input-with-select {
    width: 450px!important;
  }
  ::v-deep {
    .el-dialog {
      width: 550px;
    }
    .el-dialog__body{
      padding-top: 8px!important;
      padding-bottom: 8px!important;
    }
    .el-form{
      .el-input__inner{
        text-align: left!important;
      }
    }
    .el-input__inner {
      text-align: center !important;
    }
  }
  @media only screen and (max-width: 470px) {
    ::v-deep {
      .el-dialog {
        width: 96%;
      }
    }
  }
</style>
