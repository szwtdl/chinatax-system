<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="create = true">添加学校</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
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
            <el-table-column label="学校名称" prop="name" min-width="150" align="left" />
            <el-table-column label="学校地址" prop="address" min-width="180" align="left" />
            <el-table-column label="学校类型" align="left" min-width="200">
              <el-tag style="margin-right: 10px;">普通学校</el-tag>
              <el-tag>职业学校</el-tag>
            </el-table-column>
            <el-table-column label="学校状态" align="center" min-width="150">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.status === true" type="success">正常</el-tag>
                <el-tag v-else-if="scope.row.status === false" type="danger">禁用</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" width="250">
              <template slot-scope="scope">
                <el-button
                  size="mini"
                  type="primary"
                  @click="openEditDialog(scope.row)"
                >编辑
                </el-button>
                <el-button
                  size="mini"
                  type="danger"
                  @click="handleDelete(scope.row)"
                >删除</el-button>
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
  name: 'School',
  components: {},
  data() {
    return {
      list: [],
      loading: false,
      listQuery: {
        page: 1,
        limit: 10,
        total: 0,
        page_size: 0
      },
      create: false,
      edit: false,
      roles: [],
      account: {
        nickname: '',
        username: '',
        password: '',
        status: 1,
        selectedRoles: ['1']
      }
    }
  },
  mounted() {
  },
  methods: {
    openEditDialog(row) {
      this.edit = true
      this.account = row
    },
    handleDelete(row) {
      this.$confirm('此操作将永久删除该学校, 是否继续?', '提示', {
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
    }
  }
}
</script>
<style scoped lang="scss">

</style>
