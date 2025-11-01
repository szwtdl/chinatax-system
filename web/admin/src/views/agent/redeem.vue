<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入搜索的内容"
              size="small"
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
              <el-button type="primary" size="small" @click="visible = true">生成激活码</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%">
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
            <el-table-column label="代理账号" prop="school_name" min-width="150px" align="left" />
            <el-table-column label="激活备注" prop="remark" min-width="150px" align="left" />
            <el-table-column label="激活码" prop="code" min-width="150px" align="left" />
            <el-table-column label="激活的账号" prop="is_used" min-width="150px" align="left">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.is_used"
                  active-color="#13ce66"
                  inactive-color="#cccccc"
                />
              </template>
            </el-table-column>
            <el-table-column label="使用账号" prop="used_by" min-width="150px" align="left" />
            <el-table-column label="使用时间" prop="used_at" min-width="150px" align="left" />
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
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
    <redeem-create :visible.sync="visible" @refresh="handleRefresh" />
  </div>
</template>
<script>
import redeemCreate from '@/views/agent/components/redeem_create.vue'
import * as SchoolRedeemApi from '@/api/school_redeem'
export default {
  name: 'Redeem',
  components: {
    redeemCreate
  },
  data() {
    return {
      list: [],
      loading: false,
      visible: false,
      options: [
        { value: 'code', label: '激活码' },
        { value: 'school_id', label: '代理账号' },
        { value: 'used_by', label: '使用账号' }
      ],
      listQuery: {
        page: 1,
        limit: 20,
        total: 0,
        keywords: '',
        filter: 'code'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      SchoolRedeemApi.List(this.listQuery).then(res => {
        this.list = res.data.items
        this.listQuery.total = parseInt(res.data.total)
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSearch() {
      this.listQuery.page = 1
      this.getList()
    },
    handleRefresh() {
      this.getList()
    },
    handlerEdit(row) {
      console.log('edit', row)
    },
    handlerDelete(row) {
      console.log('delete', row)
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handlePageChange(val) {
      this.listQuery.page = val
      this.getList()
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 500px;
}
</style>
