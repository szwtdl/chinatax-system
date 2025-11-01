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
            size="mini"
            :data="list"
            border
            style="width: 100%"
            @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="ID"
              width="90"
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="会员姓名" prop="nickname" min-width="150px" align="left" />
            <el-table-column label="会员账号" prop="username" min-width="180px" align="left" />
            <el-table-column label="课程名称" prop="course_name" min-width="220px" align="left" />
            <el-table-column label="机构名称" prop="org_name" min-width="220px" align="left" />
            <el-table-column label="创建时间" prop="created_at" min-width="150px" align="left" />
            <el-table-column label="更新时间" prop="updated_at" min-width="150px" align="left" />
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
  </div>
</template>
<script>
import { fetchList } from '@/api/task_record'
export default {
  name: 'Record',
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
        { value: 'nickname', label: '用户姓名' },
        { value: 'course_name', label: '学习课程' },
        { value: 'org_name', label: '报名机构' }
      ]
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSearch() {
      this.loading = true
      this.listQuery.page = 1
      this.getList()
    },
    handleSelectionChange(selection) {
      console.log(selection)
    },
    handleSizeChange(limit) {
      this.loading = true
      this.listQuery.limit = limit
      fetchList({
        page: this.listQuery.page,
        limit: this.listQuery.limit,
        keywords: this.listQuery.keywords,
        filter: this.listQuery.filter
      }).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handlePageChange(page) {
      this.loading = true
      this.listQuery.page = page
      fetchList({
        page: this.listQuery.page,
        limit: this.listQuery.limit,
        keywords: this.listQuery.keywords,
        filter: this.listQuery.filter
      }).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
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
