<template>
  <div class="app-container">
    <div class="filter-container" />
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="addMenu">添加分类</el-button>
            </el-button-group>
          </div>
          <el-table :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              v-loading="loading"
              align="left"
              label="ID"
              width="180"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="分类名称" width="180px" align="left">
              <template slot-scope="scope">
                {{ scope.row.name }}
              </template>
            </el-table-column>
            <el-table-column label="分类描述" align="left">
              <template slot-scope="scope">
                {{ scope.row.description }}
              </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center">
              <template slot-scope="scope">
                {{ scope.row.created_at }}
              </template>
            </el-table-column>
            <el-table-column label="更新时间" align="center">
              <template slot-scope="scope">
                {{ scope.row.updated_at }}
              </template>
            </el-table-column>
            <el-table-column align="center" label="操作" width="250">
              <template #default="scope">
                <el-button
                  size="mini"
                  type="success"
                  style="margin-right: 10px"
                  @click="handleTask(scope.$index,scope.row)"
                >加入学习
                </el-button>
                <el-dropdown @command="(command) => handleCommand(scope.row, scope.$index, command)">
                  <el-button size="mini" type="primary">操作<i class="el-icon-arrow-down el-icon--right" />
                  </el-button>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item command="add">上传视频</el-dropdown-item>
                    <el-dropdown-item command="edit">编辑账号</el-dropdown-item>
                    <el-dropdown-item command="reset">重置密码</el-dropdown-item>
                    <el-dropdown-item command="delete">删除账号</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            :current-page.sync="listQuery.page"
            :page-size.sync="pageSize"
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
export default {
  name: 'Type',
  data() {
    return {
      list: [],
      total: 0,
      pageSize: 10,
      listQuery: {
        page: 1,
        limit: 10
      }
    }
  },
  methods: {
    addMenu() {
      console.log('添加分类')
    },
    handleCommand(row, index, command) {
      console.log(row)
    },
    handlePageChange(page) {
      this.listQuery.page = page
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.listQuery.limit = val
    }
  }
}
</script>
<style scoped lang="scss">

</style>
