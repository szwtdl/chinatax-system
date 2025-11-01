<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-button-group>
              <el-button type="primary" size="small" @click="dialogVisible = true">增加地区</el-button>
            </el-button-group>
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
              label="地区ID"
              min-width="100px"
              sortable
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="地区名称" prop="name" min-width="160px" align="left" />
            <el-table-column label="地区域名" prop="domain" min-width="260px" align="left" />
            <el-table-column label="地区状态" prop="status" align="center" min-width="100px" />
            <el-table-column label="地区排序" prop="sort_num" min-width="100px" align="center">
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
            <el-table-column label="创建时间" prop="created_at" min-width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
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
    <create-area-component :visible.sync="dialogVisible" />
    <edit-area-component :visible.sync="editVisible" :form="areaItem" />
  </div>
</template>
<script>
import CreateAreaComponent from '@/views/system/components/create_city.vue'
import EditAreaComponent from '@/views/system/components/edit_city.vue'
import * as ApiCity from '@/api/area'
export default {
  name: 'Area',
  components: { CreateAreaComponent, EditAreaComponent },
  data() {
    return {
      listLoading: false,
      dialogVisible: false,
      editVisible: false,
      listQuery: {
        page: 1,
        limit: 20,
        keywords: '',
        filter: ''
      },
      areaItem: {},
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
      ApiCity.List(this.listQuery).then(response => {
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
        ApiCity.Delete({
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
    handleInputChange(row) {
      this.listLoading = true
      ApiCity.Sort({ id: row.id, sort_num: parseInt(row.sort_num) }).then(() => {
        this.fetchData()
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleSelectChange(data) {
      this.select = data
    }
  }
}
</script>
<style scoped lang="scss">

</style>
