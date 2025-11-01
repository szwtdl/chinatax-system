<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <el-table :data="list" border style="width: 100%;">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column label="评论ID" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="商品ID" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.goods_id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="用户ID" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.user_id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="评论内容" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.content }}</span>
              </template>
            </el-table-column>
            <el-table-column label="评论时间" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.created_at }}</span>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
              <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handlerEdit(scope.row)">编辑</el-button>
                <el-button type="danger" size="mini" @click="handlerDelete(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  data() {
    return {
      loading: false,
      list: [],
      listQuery: {
        page: 1,
        limit: 20,
        total: 0
      }
    }
  },
  methods: {
    handleSave() {
      this.$emit('save', {
        name: this.name,
        image: this.image,
        parent_id: this.parent_id,
        sort_order: this.sort_order,
        status: this.status
      })
    },
    handlerEdit(row) {
      this.editData = Object.assign({}, row)
      this.dialogEditVisible = true
    },
    handlerDelete(row) {
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
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
