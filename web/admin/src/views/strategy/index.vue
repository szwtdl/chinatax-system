<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="show = true">新增策略</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="策略ID"
              width="100"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="平台名称" prop="platform_name" min-width="150px" align="left" />
            <el-table-column label="策略名称" prop="name" min-width="150px" align="left" />
            <el-table-column label="学习时长" prop="is_all_day" align="left" min-width="120px">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.is_pause" type="info" size="mini">{{ scope.row.hours }} 小时</el-tag>
                <el-tag v-else type="success" size="mini">全天24小时</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="并发数量" prop="max_concur" align="left" min-width="80px" />
            <el-table-column label="触发频率" prop="interval" align="left" min-width="80px">
              <template slot-scope="scope">
                <span>{{ scope.row.interval }} {{ unitToChinese(scope.row.unit) }}</span>
              </template>
            </el-table-column>
            <el-table-column label="策略状态" align="center" min-width="80px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.is_active"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="策略排序" prop="sort_num" min-width="80px" align="center">
              <template slot-scope="scope">
                <el-input
                  v-model="scope.row.sort_num"
                  placeholder="请输入排序"
                  size="small"
                  :min="0"
                  type="text"
                  style="width: 50px; text-align: center!important;"
                  @input="handleInputChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column align="center" label="操作" fixed="right" width="160px">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="handlerEdit(scope.row)"
                  >编辑
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
    <create-strategy :visible.sync="show" :options="platforms" @save="handlerSave" />
    <edit-strategy :visible.sync="edit" :strategy="strategy" :options="platforms" @save="handlerUpdate" />
  </div>
</template>
<script>

import CreateStrategy from '@/views/strategy/components/CreateStrategy.vue'
import EditStrategy from '@/views/strategy/components/EditStrategy.vue'
import * as ApiStrategy from '@/api/strategy'

export default {
  name: 'Strategy',
  components: {
    CreateStrategy,
    EditStrategy
  },
  data() {
    return {
      show: false,
      edit: false,
      loading: false,
      list: [],
      strategy: {},
      platforms: [],
      listQuery: {
        total: 0,
        page: 1,
        limit: 20
      }
    }
  },
  mounted() {
    this.getList()
  },
  created() {
    this.getOptions()
  },
  methods: {
    unitToChinese(unit) {
      switch (unit) {
        case 'second':
          return '秒'
        case 'minute':
          return '分钟'
        case 'hour':
          return '小时'
        case 'day':
          return '天'
        default:
          return unit
      }
    },
    handleInputChange(row) {
      this.listLoading = true
      ApiStrategy.Sort({ id: row.id, sort_num: parseInt(row.sort_num) }).then(() => {
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleStatusChange(row) {
      const status = row.is_active ? 'true' : 'false'
      ApiStrategy.Status({ id: row.id, status }).then(res => {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.getList(this.listQuery)
      }).finally(() => {
        this.loading = false
      })
    },
    getOptions() {
      ApiStrategy.PlatformList().then(response => {
        const { data } = response
        this.platforms = data
      })
    },
    getList() {
      ApiStrategy.StrategyList(this.listQuery).then(res => {
        this.list = res.data.items
        this.listQuery = {
          page: res.data.page,
          limit: res.data.page_size,
          total: res.data.total
        }
        this.loading = false
      })
    },
    handlerSave(data) {
      ApiStrategy.StrategyCreate(data).then(res => {
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.getList(this.listQuery)
        this.show = false
      })
    },
    handlerEdit(row) {
      this.edit = true
      this.strategy = row
    },
    handlerUpdate(data) {
      ApiStrategy.StrategyUpdate(data).then(res => {
        this.$message({
          type: 'success',
          message: '修改成功'
        })
        this.getList(this.listQuery)
      })
    },
    handlerDelete(row) {
      this.$confirm('此操作将永久删除该策略, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiStrategy.StrategyDelete({ id: row.id }).then(response => {
          this.getList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleSizeChange(size) {
      this.listQuery.limit = size
      this.getList()
    },
    handlePageChange(size) {
      this.listQuery.page = size
      this.getList()
    }
  }
}
</script>
<style scoped lang="scss">
::v-deep {
  .el-form{
    .el-input__inner{
      text-align: left!important;
    }
    .el-input-number--medium{
      .el-input__inner{
        text-align: center!important;
      }
    }
  }
  .el-input__inner {
    text-align: center !important;
  }
}
</style>
