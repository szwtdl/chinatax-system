<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <!-- 新增的下拉框 -->
            <el-select
              v-model="listQuery.school_id"
              placeholder="请选择代理商"
              size="small"
              style="width: 150px; margin-right: 5px;"
            >
              <el-option
                v-for="item in schoolRows"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
            <el-input
              v-model="listQuery.keywords"
              size="small"
              placeholder="输入订单ID/商品名称/商户编号"
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
            <el-button-group>
              <el-button type="primary" size="small" @click="visible = true">批量结算</el-button>
              <el-button type="danger" size="small" @click="handleDeleteAll">批量删除</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%" @selection-change="handleChange">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="订单ID"
              width="80px"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="课程名称" prop="course_name" min-width="220px" align="left" />
            <el-table-column label="订单编号" prop="out_trade_no" min-width="120px" align="left" />
            <el-table-column label="所属学校" prop="school_name" sortable min-width="150px" align="left" />
            <el-table-column label="客户名称" prop="nickname" min-width="150px" align="left" />
            <el-table-column label="客户账号" prop="username" min-width="100px" align="left" />
            <el-table-column label="课程价格" prop="price" min-width="80px" align="center" />
            <el-table-column label="付款状态" align="center" min-width="100px">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.trade_state === 'SUCCESS'" type="success">{{ scope.row.trade_state_desc }}</el-tag>
                <el-tag v-else-if="scope.row.trade_state === 'NOTPAY'" type="danger">{{ scope.row.trade_state_desc }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" min-width="150px" align="center">
              <template slot-scope="scope">
                {{ scope.row.created_at }}
              </template>
            </el-table-column>
            <el-table-column label="更新时间" min-width="150px" align="center">
              <template slot-scope="scope">
                {{ scope.row.updated_at }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="180px" fixed="right" align="center">
              <template slot-scope="scope">
                <el-button v-if="scope.row.trade_state ==='NOTPAY'" type="primary" size="mini" @click="handlerSettlement(scope.row)">结算</el-button>
                <div v-else>
                  <el-button type="primary" size="mini" disabled>已结算</el-button>
                  <el-button type="danger" size="mini" @click="handlerDelete(scope.row)">删除</el-button>
                </div>
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
    <check-password :visible.sync="visible" :password="password" @submit="handlerCheckPass" />
  </div>
</template>
<script>
import checkPassword from '@/views/order/components/pass.vue'
import * as SchoolOrderApi from '@/api/school_order'
import { schoolList } from '@/api/student'

export default {
  name: 'Order',
  components: {
    checkPassword
  },
  data() {
    return {
      visible: false,
      loading: false,
      list: [],
      password: '',
      selected: [],
      schoolRows: [],
      options: [
        { value: 'id', label: '订单ID' },
        { value: 'course_name', label: '课程名称' },
        { value: 'out_trade_no', label: '订单编号' },
        { value: 'nickname', label: '客户名称' },
        { value: 'username', label: '客户账号' },
        { value: 'school_name', label: '所属学校' },
        { value: 'trade_state', label: '付款状态' }
      ],
      listQuery: {
        keywords: '',
        filter: 'username',
        school_id: 0,
        page: 1,
        limit: 20,
        total: 0
      }
    }
  },
  mounted() {
    this.getList()
    this.getSchoolList()
  },
  methods: {
    getList() {
      this.loading = true
      SchoolOrderApi.SchoolOrderList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.listQuery.total = response.data.total
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    getSchoolList() {
      schoolList(this.listQuery).then(response => {
        const items = response.data || []
        this.schoolRows = [
          { value: 0, label: '请选择代理' } // 添加默认选项
        ] // 可选：初始化清空
        items.forEach(item => {
          this.schoolRows.push({
            value: item.id,
            label: item.nickname
          })
        })
      })
    },
    handlerCheckPass(formData) {
      if (this.selected.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择需要结算的订单'
        })
      }
      if (!formData.password) {
        return this.$message({
          type: 'warning',
          message: '请输入密码'
        })
      }
      SchoolOrderApi.SchoolOrderSettlement({
        ids: this.selected.map(item => item.id),
        password: formData.password
      }).then(response => {
        this.$message({
          type: 'success',
          message: '结算成功'
        })
        this.visible = false
        this.getList()
      })
    },
    handleChange(selection) {
      this.selected = selection
    },
    handleSearch() {
      this.listQuery.page = 1
      SchoolOrderApi.SchoolOrderList({
        ...this.listQuery,
        keywords: this.listQuery.keywords
      }).then(response => {
        this.list = response.data.items
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.listQuery.total = response.data.total
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handlerSettlement(row) {
      this.visible = true
      this.password = ''
      this.selected = [row]
      this.$nextTick(() => {
        this.$refs.checkPassword.reset()
      })
    },
    handlerDelete(row) {
      this.$confirm('此操作将永久删除该订单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        SchoolOrderApi.SchoolOrderDelete({ ids: [row.id] }).then(() => {
          this.$message.success('删除成功')
          this.getList()
        })
      })
    },
    handleDeleteAll() {
      if (this.selected.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择需要删除的订单'
        })
      }
      this.$confirm('此操作将永久删除所选订单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        SchoolOrderApi.SchoolOrderDelete({ ids: this.selected.map(item => item.id) }).then(() => {
          this.$message.success('删除成功')
          this.getList()
        })
      })
    },
    handlePageChange(page) {
      this.loading = true
      this.listQuery.page = page
      SchoolOrderApi.SchoolOrderList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSizeChange(val) {
      this.loading = true
      this.listQuery.page = 1
      this.listQuery.limit = val
      SchoolOrderApi.SchoolOrderList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
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
::v-deep {
  .el-dialog__body{
    padding: 0 20px!important;
  }
}
</style>
