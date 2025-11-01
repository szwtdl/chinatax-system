<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入内容"
              class="input-with-select"
              size="small"
              style="margin-right: 5px;"
              @keyup.enter.native="handleSearch"
            >
              <el-select slot="prepend" v-model="listQuery.filter" size="small" placeholder="请选择" style="width: 120px;">
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
              <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
            </el-input>
          </div>
          <el-table
            v-loading="loading"
            :data="list"
            border
            size="mini"
            style="width: 100%"
          >
            <el-table-column type="selection" align="center" min-width="60px" />
            <el-table-column prop="student_id" label="用户ID" min-width="100px" />
            <el-table-column prop="nickname" label="登录用户" min-width="120px" />
            <el-table-column prop="username" label="登录账号" min-width="160px" />
            <el-table-column prop="login_ip" label="登录IP" align="left" min-width="150px" />
            <el-table-column prop="user_agent" label="浏览器" min-width="150px" align="left">
              <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" :content="scope.row.browser_name" placement="top">
                  <span> {{ scope.row.system_name }} {{ scope.row.browser_name }} {{ scope.row.browser_version }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column label="登录消息" prop="message" align="left" min-width="150px" />
            <el-table-column label="创建时间" prop="created_at" align="left" min-width="180px" />
            <el-table-column label="更新时间" prop="updated_at" align="left" min-width="180px" />
          </el-table>
        </el-card>
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
      </el-col>
    </el-row>
  </div>
</template>
<script>
import * as StudentLogoApi from '@/api/student_logs'

export default {
  name: 'UserLogs',
  data() {
    return {
      loading: false,
      list: [],
      total: 0,
      listQuery: {
        page: 1,
        limit: 20,
        keywords: '',
        filter: 'username'
      },
      options: [
        { value: 'username', label: '会员账号' },
        { value: 'nickname', label: '用户姓名' }
      ]
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      StudentLogoApi.StudentLogoList(this.listQuery).then(res => {
        this.list = res.data.items
        this.total = res.data.total
        this.listQuery = {
          page: res.data.page,
          limit: res.data.page_size
        }
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handlePageChange(val) {
      this.listQuery.page = val
      this.getList()
    },
    handleSearch() {
      this.listQuery.page = 1
      this.getList()
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}

</style>
