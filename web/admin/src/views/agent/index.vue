<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="show = true">添加代理</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60px" />
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
            <el-table-column label="登录账号" prop="username" min-width="150px" align="left" />
            <el-table-column label="账号昵称" prop="nickname" min-width="180px" align="left" />
            <el-table-column label="学员密码" prop="invite_code" min-width="180px" align="left" />
            <el-table-column label="授信额度" prop="credit" min-width="180px" align="left">
              <template slot-scope="scope">
                <span>{{ scope.row.credit }} 次</span>
              </template>
            </el-table-column>
            <el-table-column label="API令牌" prop="tokens" min-width="180px" align="center">
              <template slot-scope="scope">
                <div v-if="scope.row.tokens && scope.row.tokens.length > 0" style="display: flex; align-items: center; justify-content: center;">
                  <span style="margin-right: 2px;">{{ scope.row.tokens[0].token }}</span>
                  <el-tooltip content="复制令牌" placement="top">
                    <el-button
                      size="mini"
                      type="text"
                      icon="el-icon-document-copy"
                      @click="copyToken(scope.row.tokens[0].token)"
                    />
                  </el-tooltip>
                </div>
                <el-button
                  v-else
                  size="mini"
                  type="primary"
                  @click="GenerateToken(scope.row)"
                >生成令牌</el-button>
              </template>
            </el-table-column>
            <el-table-column label="账号状态" align="center" min-width="150px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.status"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" min-width="180px">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button size="small" type="primary" @click="handlerToken(scope.row)">
                    令牌密钥
                  </el-button>
                  <el-button size="small" type="primary" @click="handlerEdit(scope.row)">
                    编辑
                  </el-button>
                  <el-button size="small" type="danger" @click="handlerDelete(scope.row)">
                    删除
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
    <create-account :visible.sync="show" @save="handleSave" />
    <edit-account :visible.sync="edit" :data="school" @save="handleUpdate" />
    <token-component :visible.sync="token" :item="schoolId" />
  </div>
</template>
<script>
import createAccount from '@/views/agent/components/create.vue'
import editAccount from '@/views/agent/components/edit.vue'
import tokenComponent from '@/views/agent/components/token.vue'
import * as SchoolApi from '@/api/school'

export default {
  name: 'Index',
  components: {
    createAccount,
    editAccount,
    tokenComponent
  },
  data() {
    return {
      show: false,
      edit: false,
      token: false,
      loading: false,
      list: [],
      school: {},
      schoolId: 0,
      listQuery: {
        total: 0,
        page: 1,
        limit: 20
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      SchoolApi.SchoolList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery = {
          page: response.data.page,
          limit: response.data.page_size,
          total: response.data.total
        }
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleSave(data) {
      this.show = false
      SchoolApi.SchoolCreate(data).then(() => {
        this.$message.success('添加成功')
        this.getList()
      })
    },
    handlerToken(row) {
      this.token = true
      this.schoolId = row.id
    },
    copyToken(token) {
      navigator.clipboard.writeText(token).then(() => {
        this.$message.success('令牌已复制')
      }).catch(() => {
        this.$message.error('复制失败，请手动复制')
      })
    },
    handlerEdit(row) {
      this.edit = true
      this.school = row
    },
    handleUpdate(data) {
      SchoolApi.SchoolUpdate(data).then(() => {
        this.$message.success('更新成功')
        this.edit = false
        this.getList()
      })
    },
    handlerDelete(row) {
      this.$confirm('此操作将永久删除该代理, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        SchoolApi.SchoolDelete({ 'id': row.id }).then(() => {
          this.$message.success('删除成功')
          this.getList()
        })
      }).catch(() => {
        this.$message.info('已取消删除')
      })
    },
    handleStatusChange(row) {
      const status = row.status ? 'true' : 'false'
      SchoolApi.status({
        'id': row.id,
        status
      }).then(() => {
        this.$message.success('状态更新成功')
        this.getList()
      })
    },
    GenerateToken(row) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      SchoolApi.GenerateToken({
        'id': row.id
      }).then(response => {
        this.$message({
          type: 'success',
          message: response.data
        })
        this.getList()
      }).finally(() => {
        loading.close()
      })
    },
    handleSizeChange(number) {
    },
    handlePageChange(size) {
    }
  }
}
</script>
<style scoped lang="scss">
::v-deep {
  .el-message-box {
    width: 1200px!important; /* 根据需要设置宽度 */
  }
}
</style>
