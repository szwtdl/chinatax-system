<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入内容"
              size="small"
              class="input-with-select"
              style="margin-right: 5px;"
              @keyup.enter.native="handleSearch"
            >
              <el-select slot="prepend" v-model="listQuery.filter" placeholder="请选择" style="width: 120px;">
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
              <el-button type="primary" size="small" @click="handlerRefresh">同步代理</el-button>
              <el-button type="primary" size="small" @click="handleProxyPing">批量检测</el-button>
            </el-button-group>
          </div>
          <el-table v-loading="loading" :data="list" size="mini" border style="width: 100%" @selection-change="handleProxyChange">
            <el-table-column type="selection" align="center" width="60px" />
            <el-table-column label="代理平台" prop="proxy_name" align="left" />
            <el-table-column label="代理地区" prop="region_name" width="160px" align="left" />
            <el-table-column label="代理IP" prop="proxy_ip" min-width="220px" align="left">
              <template slot-scope="scope">
                <el-tooltip v-if="scope.row.proxy_domain" class="item" effect="dark" :content="`${scope.row.proxy_ip}:${scope.row.proxy_port}`" placement="top">
                  <span v-if="scope.row.proxy_domain!== ''">{{ scope.row.proxy_domain }}:{{ scope.row.proxy_port }}</span>
                  <span v-else>{{ scope.row.proxy_ip }}:{{ scope.row.proxy_port }}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column label="代理类型" prop="proxy_type" />
            <el-table-column label="代理账号" prop="proxy_user" align="left" />
            <el-table-column label="代理密码" prop="proxy_pass" align="left" />
            <el-table-column label="节点切换" prop="use_switch" align="left" />
            <el-table-column label="已使用数" prop="quantity_used" align="left" />
            <el-table-column label="剩余时长" prop="proxy_hours" align="left" />
            <el-table-column align="center" label="操作" fixed="right" width="250">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="handlerRegions(scope.row)"
                  >切换
                  </el-button>
                  <el-button
                    size="mini"
                    type="primary"
                    @click="handlerTestProxy(scope.row)"
                  >测试
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
    <region-list :visible.sync="regionDialog" :items="regions" :groups="region_type" @select="handlerRegion" @close="handleClose" />
  </div>
</template>
<script>

import * as ApiProxy from '@/api/proxy'
import RegionList from './components/region_list.vue'

export default {
  name: 'Strategy',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    RegionList
  },
  data() {
    return {
      show: false,
      edit: false,
      regionDialog: false,
      loading: false,
      regions: [],
      region_type: [
        { 'isSelected': true, 'name': '电信' },
        { 'isSelected': false, 'name': '联通' },
        { 'isSelected': false, 'name': '移动' }
      ],
      options: [
        { value: 'proxy_name', label: '代理平台' },
        { value: 'proxy_ip', label: '代理IP' },
        { value: 'region_name', label: '代理地区' }
      ],
      list: [],
      proxy: {},
      select: [],
      listQuery: {
        total: 0,
        page: 1,
        limit: 20,
        keywords: '',
        filter: 'region_name'
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      ApiProxy.ProxyList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
        this.loading = false
      })
    },
    handlerDelete(row) {
      this.$confirm('此操作将永久删除该代理, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiProxy.ProxyDelete({ id: row.id }).then(response => {
          this.getList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handlerTestProxy(row) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      ApiProxy.ProxyPing({ ids: [row.id], type: 1 }).then(res => {
        loading.close()
        if (res.code === 0 && res.data && res.data.length > 0) {
          const result = res.data[0]
          if (result.result === 'success') {
            this.$message({
              type: 'success',
              message: '测试可用'
            })
          } else {
            this.$message({
              type: 'error',
              message: result.errmsg || '测试失败'
            })
          }
        } else {
          this.$message.error(res.msg || '请求失败')
        }
      }).finally(() => {
        loading.close()
      })
    },
    handlerRegions(row) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      if (row.use_switch < 1) {
        this.$message({
          type: 'warning',
          message: '没有切换次数了'
        })
        return
      }
      this.proxy = row
      ApiProxy.ProxyRegions({ item: row.item_name, stepid: row.price_step_id }).then(res => {
        this.regions = res.data
        this.regionDialog = true
      }).finally(() => {
        loading.close()
      })
    },
    handleSearch() {
      this.listQuery.page = 1
      this.getList()
    },
    handleProxyChange(val) {
      this.select = val
    },
    handleProxyPing() {
      var ids = []
      this.select.forEach(item => {
        ids.push(item.id)
      })
      if (ids.length < 1) {
        this.$message({
          type: 'warning',
          message: '请先选择要检测的代理'
        })
        return
      }
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      ApiProxy.ProxyPing({ ids: ids, type: 1 }).then(res => {
        loading.close()
        if (res.code === 0) {
          const failed = res.data.filter(item => item.result !== 'success')
          if (failed.length > 0) {
            // 有失败的节点
            this.$message({
              type: 'warning',
              message: `部分节点测试失败：${failed.map(f => f.name + '(' + f.errmsg + ')').join('，')}`
            })
          } else {
            // 全部成功
            this.$message({
              type: 'success',
              message: '批量测试成功'
            })
          }

          this.getList()
        } else {
          this.$message.error(res.msg || '操作失败')
        }
      }).finally(() => {
        loading.close()
      })
    },
    handlerRegion(row) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      ApiProxy.ProxySwitch({
        proxy_id: this.proxy.id,
        order_id: this.proxy.order_id,
        region_id: row.number.toString()
      }).then(res => {
        loading.close()
        this.$message({
          type: 'success',
          message: '切换成功'
        })
        this.getList()
      }).finally(() => {
        loading.close()
      })
    },
    handleClose() {
      this.regionDialog = false
    },
    handlerRefresh() {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      ApiProxy.ProxyRefresh({ page: 1, limit: 100 }).then(res => {
        this.getList(this.listQuery)
        this.show = false
      }).finally(() => {
        loading.close()
      })
    },
    handleSizeChange(number) {
      this.listQuery.limit = number
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
.input-with-select {
  width: 450px;
}
</style>
