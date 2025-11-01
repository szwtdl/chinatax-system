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
            size="mini"
            border
            style="width: 100%"
          >
            <el-table-column type="selection" align="center" />
            <el-table-column
              align="left"
              label="ID"
              width="100px"
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="会员姓名" prop="nickname" min-width="160px" align="left" />
            <el-table-column label="会员账号" prop="username" min-width="160px" align="left" />
            <el-table-column label="机构名称" prop="org_name" min-width="220px" align="left" />
            <el-table-column label="人脸时间" prop="verify_time" align="left" />
            <el-table-column label="人脸照片" width="100px" align="center">
              <template slot-scope="scope">
                <el-image
                  style="width: 30px;height: 30px;"
                  :src="scope.row.url"
                  :preview-src-list="[scope.row.url]"
                  fit="cover"
                  @click.native="handleImageClick"
                />
              </template>
            </el-table-column>
            <el-table-column label="认证状态" prop="status" min-width="250px" align="left">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.status === '人脸核验成功' || scope.row.status === '认证成功！' || scope.row.status === '认证成功，请在电脑上继续学习！' || scope.row.status === '处理成功' || scope.row.status === '成功' || scope.row.status === '识别成功' || scope.row.status === '对比成功' || scope.row.status === '认证成功，请继续学习！'" type="success">{{ scope.row.status }}</el-tag>
                <el-tag v-else type="danger">{{ scope.row.status }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="160px" align="left" />
            <el-table-column label="更新时间" prop="updated_at" width="160px" align="left" />
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
import { fetchList } from '@/api/cert_record'

export default {
  name: 'CertRecord',
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
        { value: 'org_name', label: '报名机构' }
      ]
    }
  },
  created() {
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
    handleImageClick(event) {
      event.preventDefault()
      event.stopPropagation()
    },
    handleSizeChange(limit) {
      this.listQuery.page = 1
      this.listQuery.limit = limit
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
      }).finally(() => {
        this.loading = false
      })
    },
    handlePageChange(page) {
      this.listQuery.page = page
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
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
