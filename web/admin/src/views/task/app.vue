<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <div class="filter-container">
          <el-input
            v-model="listQuery.keywords"
            placeholder="请输入内容"
            class="input-with-select"
            style="margin-right: 5px;"
            @keyup.enter.native="handleSearch"
          >
            <el-select slot="prepend" v-model="listQuery.filter" placeholder="请选择" style="width: 120px;">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
            <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
          </el-input>
          <el-button-group>
            <el-button type="danger" @click="handleClear">批量删除</el-button>
          </el-button-group>
        </div>
        <el-table
          v-loading="loading"
          size="mini"
          :data="items"
          border
          style="width: 100%"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" align="center" width="60" />
          <el-table-column prop="id" label="任务ID" width="150" />
          <el-table-column prop="nickname" label="客户姓名" />
          <el-table-column prop="username" label="客户账号" />
          <el-table-column prop="course_name" label="课程名称" />
          <el-table-column prop="org_name" label="机构名称" />
          <el-table-column prop="isLive" label="是否在线" />
          <el-table-column prop="created_at" label="创建时间" align="center" />
          <el-table-column prop="updated_at" label="更新时间" align="center" />
          <el-table-column align="center" label="操作" width="200px">
            <template #default="scope">
              <el-button
                size="mini"
                type="danger"
                @click="handleDelete(scope.$index,scope.row)"
              >移除
              </el-button>
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
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  name: 'App',
  data() {
    return {
      loading: false,
      items: [],
      options: [
        { value: 'username', label: '登录账号' },
        { value: 'nickname', label: '客户姓名' },
        { value: 'org_name', label: '学习机构' },
        { value: 'course_name', label: '课程名称' }
      ],
      multipleSelection: [],
      listQuery: {
        page: 1,
        limit: 20,
        total: 0,
        keywords: '',
        filter: 'username'
      }
    }
  },
  mounted() {
  },
  methods: {
    getData() {
    },
    handleClear() {
      this.$confirm('此操作将永久删除选中数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleDelete(index, row) {
      this.$confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleSearch() {
      this.loading = true
      setTimeout(() => {
        this.loading = false
      }, 1000)
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    handleSizeChange(limit) {
      this.listQuery.limit = limit
      this.getData()
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.getData()
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
