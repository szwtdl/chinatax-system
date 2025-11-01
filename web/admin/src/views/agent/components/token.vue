<template>
  <div class="app-container">
    <el-dialog
      title="API接口令牌"
      :visible.sync="visible"
      width="50%"
      :before-close="handleClose"
    >
      <div class="container">
        <div class="filter-container">
          <el-button-group>
            <el-button type="primary" size="small" @click="handleGenerate">增加令牌</el-button>
            <el-button type="danger" size="small" @click="handleBatch">批量删除</el-button>
          </el-button-group>
        </div>
        <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%" @selection-change="handleTokenChange">
          <el-table-column type="selection" align="center" width="60px" />
          <el-table-column
            align="left"
            label="ID"
            min-width="100px"
            element-loading-text="请给我点时间！"
          >
            <template slot-scope="scope">
              <span>{{ scope.row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column label="API令牌" prop="token" min-width="180px" align="left">
            <template slot-scope="scope">
              <span>{{ scope.row.token }}</span>
              <el-tooltip content="复制令牌" placement="top">
                <el-button
                  size="mini"
                  type="text"
                  icon="el-icon-document-copy"
                  @click="copyToken(scope.row.token)"
                />
              </el-tooltip>
            </template>
          </el-table-column>
          <el-table-column label="过期时间" prop="ExpiredAt" min-width="180px" align="left">
            <template slot-scope="scope">
              <span>{{ scope.row.expired_at }}</span>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" prop="created_at" min-width="180px" align="center" />
          <el-table-column align="center" label="操作" fixed="right" min-width="150px">
            <template slot-scope="scope">
              <el-button-group>
                <el-button
                  size="small"
                  type="primary"
                  @click="handlerUpdate(scope.row)"
                >刷新令牌
                </el-button>
                <el-button
                  size="small"
                  type="danger"
                  @click="handlerDelete(scope.row)"
                >删除
                </el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import * as ApiSchoolToken from '@/api/school_token'
export default {
  name: 'Token',
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    item: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      loading: false,
      list: [],
      total: 0,
      listQuery: {
        page: 1,
        limit: 20,
        school_id: 0
      },
      multipleSelection: []
    }
  },
  watch: {
    item(val) {
      if (val) {
        this.listQuery.school_id = val
        this.getList()
      }
    }
  },
  methods: {
    getList() {
      this.loading = true
      ApiSchoolToken.List(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.loading = false
      })
    },
    deleteItems(ids) {
      ApiSchoolToken.del({ ids: ids }).then(() => {
        this.$message.success('删除成功')
        this.getList()
      })
    },
    copyToken(token) {
      navigator.clipboard.writeText(token).then(() => {
        this.$message.success('令牌已复制')
      }).catch(() => {
        this.$message.error('复制失败，请手动复制')
      })
    },
    handleGenerate() {
      ApiSchoolToken.create({ school_id: this.item }).then(() => {
        this.$message.success('生成成功')
        this.getList()
      })
    },
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleTokenChange(val) {
      this.multipleSelection = val
    },
    handleBatch() {
      if (this.multipleSelection.length === 0) {
        this.$message.warning('请选择要删除的项')
        return
      }
      this.$confirm('是否确认删除选中项？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const ids = []
        this.multipleSelection.forEach(item => {
          ids.push(item.id)
        })
        // 调用删除接口
        this.deleteItems(ids)
      }).catch(() => {
        this.$message.info('已取消删除')
      })
    },
    handlerUpdate(row) {
      this.$confirm('是否确认重新生成该代理的API令牌？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiSchoolToken.update({ id: row.id }).then(() => {
          this.$message.success('API令牌生成成功')
          this.getList()
        })
      }).finally(() => {
      })
    },
    handlerDelete(row) {
      this.$confirm('是否确认删除该代理？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 调用删除接口
        this.deleteItems([row.id])
      })
    }
  }
}
</script>
