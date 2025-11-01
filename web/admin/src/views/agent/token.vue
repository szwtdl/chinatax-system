<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="danger" size="mini" @click="handleBatch">批量删除</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%" @selection-change="handleTokenChange">
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
            <el-table-column label="代理账号" prop="username" min-width="180px" align="left" />
            <el-table-column label="代理昵称" prop="Nickname" min-width="150px" align="left">
              <template slot-scope="scope">
                <span>{{ scope.row.nickname }}</span>
              </template>
            </el-table-column>
            <el-table-column label="API令牌" prop="token" min-width="180px" align="left">
              <template slot-scope="scope">
                <span>{{ scope.row.token }}</span>
              </template>
            </el-table-column>
            <el-table-column label="过期时间" prop="ExpiredAt" min-width="180px" align="left">
              <template slot-scope="scope">
                <span>{{ scope.row.expired_at }}</span>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" min-width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" min-width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" min-width="160px">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="handlerUpdate(scope.row)"
                  >刷新令牌
                  </el-button>
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handlerDelete(scope.row)"
                  >删除
                  </el-button>
                </el-button-group>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import * as ApiSchoolToken from '@/api/school_token'
export default {
  name: 'Token',
  data() {
    return {
      loading: false,
      showDialog: false,
      list: [],
      listQuery: {
        total: 0,
        page: 1,
        limit: 20
      },
      selected: []
    }
  },
  mounted() {
    this.getItems()
  },
  methods: {
    getItems() {
      this.loading = true
      ApiSchoolToken.List(this.listQuery).then(response => {
        const { items, page, page_size, total } = response.data
        this.list = items || []
        this.listQuery.page = page || 1
        this.listQuery.limit = page_size || 20
        this.listQuery.page = page || 1
        this.total = total || 0
      }).finally(() => {
        this.loading = false
      })
    },
    handleTokenChange(selection) {
      this.selected = selection.map(item => item.id)
    },
    handlerUpdate(row) {
      this.$confirm('是否确认重新生成该代理的API令牌？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiSchoolToken.update({ id: row.id }).then(() => {
          this.$message.success('API令牌生成成功')
          this.getItems()
        })
      }).finally(() => {
      })
    },
    handleBatch() {
      if (this.selected.length === 0) {
        this.$message.warning('请先选择要删除的项')
        return
      }
      const ids = this.selected.map(item => item.id)
      this.deleteItems(ids)
    },
    handlerDelete(row) {
      const data = this.selected.length > 0 ? this.selected : [row.id]
      this.$confirm('是否确认删除选中项？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteItems(data)
      }).finally(() => {
      })
    },
    deleteItems(ids) {
      ApiSchoolToken.del({ ids: ids }).then(() => {
        this.$message.success('删除成功')
        this.getItems()
      })
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.getItems()
    },
    handleSizeChange(size) {
      this.listQuery.limit = size
      this.getItems()
    }
  }
}
</script>

<style scoped lang="scss">

</style>
