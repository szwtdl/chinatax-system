<template>
  <div class="app-container">
    <el-dialog
      :title="member.title"
      width="1366px"
      :visible.sync="member.dialogVisible"
    >
      <el-row>
        <el-col :span="24">
          <div class="filter-container">
            <el-input
              v-model="member.listQuery.keywords"
              placeholder="请输入账号或姓名搜索"
              class="input-with-select"
              style="margin-right: 5px;"
              @keyup.enter.native="handleSearch"
            >
              <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
            </el-input>
            <el-button-group>
              <el-button type="primary" @click="handleSubmit">批量导入</el-button>
            </el-button-group>
          </div>
          <el-table
            v-loading="member.loading"
            :data="member.items"
            size="mini"
            border
            style="width: 100%"
            max-height="450"
            @selection-change="selectionChange"
          >
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              prop="id"
              label="ID"
              width="100"
              element-loading-text="请给我点时间！"
            />
            <el-table-column prop="nickname" label="学员姓名" width="100" align="center" />
            <el-table-column prop="username" label="登录账号" width="180" align="center" />
            <el-table-column prop="sex" label="学员性别" width="80" align="center" />
            <el-table-column prop="age" label="学员年龄" width="80" align="center" />
            <el-table-column label="学员课程" align="left">
              <template slot-scope="scope">
                <el-tag
                  v-for="(item, index) in scope.row.course"
                  :key="index"
                  style="margin-right: 5px;margin-bottom: 5px;"
                >
                  {{ item }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="created_at" label="注册时间" width="150" align="center" />
          </el-table>
          <el-pagination
            background
            layout="total, prev, pager, next, jumper"
            :total="member.listQuery.total"
            :current-page.sync="member.listQuery.page"
            :page-size.sync="member.listQuery.pageSize"
            style="margin-top: 20px; text-align: right;"
            @current-change="handlePageChange"
          />
        </el-col>
      </el-row>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: 'OrgMember',
  props: {
    member: {
      type: Object,
      default: () => ({
        items: [],
        title: '',
        loading: true,
        dialogVisible: true,
        listQuery: {
          keywords: '',
          filter: '',
          total: 0,
          page: 1,
          pageSize: 20
        }
      })
    }
  },
  methods: {
    handleSubmit() {
      this.$emit('submit')
    },
    handleSearch() {
      this.loading = true
      this.$emit('search', this.listQuery)
      setTimeout(function() {
        this.loading = false
      }, 1000)
    },
    selectionChange(selection) {
      this.$emit('change', selection)
    },
    handlePageChange(page) {
      this.$emit('page', page)
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
