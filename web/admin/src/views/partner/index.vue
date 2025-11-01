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
              label="ID编号"
              sortable
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="LOGO" prop="avatar" align="left">
              <template slot-scope="scope">
                <el-avatar class="custom-avatar" size="small" :src="scope.row.avatar" />
              </template>
            </el-table-column>
            <el-table-column label="公司名称" prop="nickname" align="left" />
            <el-table-column label="登录账号" prop="username" align="left" />
            <el-table-column label="登录状态" align="left">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.status"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="登录次数" prop="login_count" align="left" />
            <el-table-column label="最后登录" prop="last_time" align="left" />
            <el-table-column label="最后IP" prop="login_ip" align="left" />
            <el-table-column label="注册时间" prop="register_time" align="left" />
            <el-table-column label="更新时间" prop="updated_at" align="left" />
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
                      <el-dropdown-item command="delete">删除地区</el-dropdown-item>
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
  </div>
</template>
<script>
import * as ApiPartner from '@/api/partner'

export default {
  name: 'Partner',
  data() {
    return {
      listQuery: {
        page: 1,
        limit: 20,
        keywords: '',
        filter: 'username'
      },
      options: [
        {
          value: 'username',
          label: '登录账号'
        },
        {
          value: 'nickname',
          label: '公司名称'
        }
      ],
      items: [],
      select: [],
      total: 0,
      total_page: 0,
      pageSize: 0
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      ApiPartner.List(this.listQuery).then(response => {
        const { data } = response
        this.items = data.items
        this.total = data.total
        this.total_page = data.total_page
        this.pageSize = data.limit
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleStatusChange(row) {
      this.listLoading = true
      const status = row.status ? 'true' : 'false'
      ApiPartner.Status({ id: row.id, status }).then(() => {
        this.$message({
          message: '操作成功',
          type: 'success'
        })
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleSearch() {
      this.listLoading = true
      ApiPartner.List(this.listQuery).then(response => {
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
    handleEdit(index, row) {
      this.editVisible = true
      this.areaItem = row
    },
    handleCommand(row, index, command) {
      this.listLoading = true
      if (command === 'delete') {
        ApiPartner.Delete({
          id: row.id
        }).then(response => {
          this.$message({
            message: '操作成功',
            type: 'success'
          })
        }).finally(() => {
          this.listLoading = false
        })
      }
    },
    handleSizeChange(limit) {
      this.listQuery.limit = limit
      this.fetchData()
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.fetchData()
    },
    handleSelectChange(data) {
      this.select = data
    }
  }
}
</script>
<style scoped lang="scss">
.el-avatar {
  border-radius: 2px !important;
  position: relative;
  top: 5px!important;
}
.input-with-select{
  width: 540px!important;
}
</style>
