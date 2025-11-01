<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <!-- 新增的下拉框 -->
            <el-select
              v-model="listQuery.school_id"
              placeholder="请选择代理商"
              size="small"
              style="width: 150px; margin-right: 5px;"
            >
              <el-option
                v-for="item in schoolRows"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
            <!-- 新增的下拉框 -->
            <el-select
              v-model="listQuery.proxy_pool_id"
              placeholder="IP代理"
              class="select-type"
              size="small"
              style="width: 220px; margin-right: 5px;"
            >
              <el-option
                v-for="item in proxyListRows"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
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
            <el-button-group>
              <el-button type="primary" size="small" @click="handleProxy">切换IP代理</el-button>
              <el-button type="primary" size="small" @click="handleStrategy">切换策略</el-button>
              <el-button
                size="small"
                :type="isRunning ? 'danger' : 'primary'"
                @click="isRunning ? handleStop() : handleStart()"
              >
                {{ isRunning ? '停止任务' : '启动任务' }}
              </el-button>
            </el-button-group>
          </div>
          <el-table
            v-loading="loading"
            size="mini"
            :data="list"
            border
            style="width: 100%"
            @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="任务ID"
              width="100px"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="会员姓名" prop="name" min-width="120px" align="left" />
            <el-table-column label="会员账号" prop="username" width="160px" align="left">
              <template slot-scope="scope">
                <span>{{ scope.row.username }} <small v-if="scope.row.proxy_pool_id !== 0" style="color: #2ac06d;">【{{ scope.row.proxy_name }}】</small></span>
              </template>
            </el-table-column>
            <el-table-column label="课程名称" prop="course_name" width="230px" align="left" />
            <el-table-column label="所属学校" prop="school_name" sortable width="150px" align="left" />
            <el-table-column label="机构名称" prop="org_name" width="220px" sortable align="left" />
            <el-table-column label="学习进度" prop="course_name" width="300px" align="left">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.progress !== ''">{{ scope.row.lesson_name }} 【{{ scope.row.progress }}%】</el-tag>
                <el-tag v-else>待更新</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="学习状态" align="left" width="120px">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.status === 0" type="success">正常学习</el-tag>
                <el-tag v-else-if="scope.row.status === 1">考试答题</el-tag>
                <el-tag v-else-if="scope.row.status === 2" type="danger">需要扫脸</el-tag>
                <el-tag v-else-if="scope.row.status === 3">人脸认证</el-tag>
                <el-tag v-else-if="scope.row.status === 4">人脸失败</el-tag>
                <el-tag v-else-if="scope.row.status === 5">没有照片</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="学习策略" prop="strategy_name" sortable width="120px" align="left">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.strategy_name !== ''">{{ scope.row.strategy_name }}</el-tag>
                <el-tag v-else>没有策略</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
            <el-table-column align="left" label="操作" fixed="right" width="180">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handleDelete(scope.$index,scope.row)"
                  >移除
                  </el-button>
                  <el-button
                    v-if="scope.row.platform === 'local' || scope.row.platform === 'school'"
                    size="mini"
                    type="success"
                    @click="handleRefresh(scope.$index,scope.row)"
                  >进度
                  </el-button>
                  <el-button
                    v-if="scope.row.status === 2"
                    size="mini"
                    type="primary"
                    @click="openAlbum(scope.row)"
                  >认证
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
    <!-- 认证照片选择 -->
    <photo-album :visible.sync="showDialog" :member="user_id" @select="handleSelect" @close="handleClose" />
    <strategy-component :visible.sync="strategyDialog" :items="strategyRows" @select="handleStrategyUpdate" @close="handleClose" />
    <proxy-component :visible.sync="proxyDialog" :items="proxyRows" @select="handleProxyUpdate" @close="handleClose" />
  </div>
</template>
<script>

import { upload } from '@/api/photo'
import {
  taskList,
  refreshProgress,
  deleteTask,
  clearTask,
  stopTask,
  startTask,
  running,
  proxySelectList,
  strategyList,
  strategyUpdate,
  proxyUpdate
} from '@/api/task'
import { ProxyList } from '@/api/proxy'
import photoAlbum from './components/photo_album.vue'
import strategyComponent from './components/strategy_list.vue'
import proxyComponent from './components/proxy_list.vue'
import { schoolList } from '@/api/student'

export default {
  name: 'Task',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    strategyComponent,
    proxyComponent,
    photoAlbum
  },
  data() {
    return {
      list: [],
      showDialog: false,
      user_id: 0,
      isRunning: false,
      selectedRows: [],
      strategyRows: [],
      schoolRows: [
        { value: 0, label: '全部代理商' } // 添加默认选项
      ],
      proxyRows: [],
      proxyListRows: [
        { value: 0, label: '全部代理IP' } // 添加默认选项
      ],
      strategyDialog: false,
      proxyDialog: false,
      loading: false,
      listQuery: {
        page: 1,
        limit: 20,
        total: 0,
        school_id: 0,
        proxy_pool_id: 0,
        filter: 'username',
        keywords: ''
      },
      options: [
        {
          value: 'username',
          label: '会员账号'
        },
        {
          value: 'name',
          label: '会员姓名'
        },
        {
          value: 'course_name',
          label: '课程名称'
        }
      ]
    }
  },
  mounted() {
    this.getList()
    this.getProxyList()
    this.getSchoolList()
    running().then(res => {
      this.isRunning = res.data.is_running
    })
  },
  methods: {
    getSchoolList() {
      schoolList(this.listQuery).then(response => {
        const items = response.data || []
        items.forEach(item => {
          this.schoolRows.push({
            value: item.id,
            label: item.nickname
          })
        })
      })
    },
    getProxyList() {
      proxySelectList({
        page: 1,
        limit: 100
      }).then(response => {
        const items = response.data || []
        items.forEach(item => {
          this.proxyListRows.push({
            value: item.id,
            label: item.region_name
          })
        })
      })
    },
    getList() {
      this.loading = true
      taskList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
        this.loading = false
      })
    },
    handleSearch() {
      this.loading = true
      this.listQuery.page = 1
      taskList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleStart() {
      startTask().then(response => {
        this.isRunning = true
        this.$message({
          type: 'success',
          message: response.data
        })
      })
    },
    handleStop() {
      this.$confirm('您确定要停止任务吗, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        stopTask().then(response => {
          this.isRunning = false
          this.$message({
            type: 'success',
            message: response.data
          })
        })
      })
    },
    handleDelete(index, row) {
      this.$confirm('此操作将永久删除该, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTask({
          ids: [row.id]
        }).then(response => {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.getList()
        })
      })
    },
    handleRefresh(index, row) {
      this.$confirm('此操作将刷新进度, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        refreshProgress({
          task_id: row.id,
          user_id: row.user_id,
          username: row.username,
          password: row.password,
          course_id: row.course_id,
          region_id: row.region_id,
          org_id: row.org_id
        }).then(response => {
          this.$message({
            type: 'success',
            message: '刷新成功'
          })
          this.getList()
        })
      })
    },
    handleSelectionChange(selection) {
      this.selectedRows = selection
    },
    openAlbum(row) {
      this.user_id = row.user_id
      this.showDialog = true
    },
    handleSelect(item) {
      var data = {
        user_id: this.user_id,
        url: item.url
      }
      this.list.forEach((element, index) => {
        if (element.user_id === this.user_id) {
          data.task_id = element.id
          data.token = element.token
          data.org_id = element.org_id
        }
      })
      upload(data).then(response => {
        this.$message({
          type: 'success',
          message: '上传成功'
        })
        this.showDialog = false
        this.getList()
      })
    },
    handleClose() {
      this.show = false
    },
    handleStrategy() {
      var items = []
      for (let i = 0; i < this.selectedRows.length; i++) {
        items.push(this.selectedRows[i].id)
      }
      if (this.selectedRows.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择需要更新的任务'
        })
        return
      }
      strategyList({
        page: 1,
        limit: 100
      }).then(response => {
        this.strategyDialog = true
        this.strategyRows = response.data
      })
    },
    handleStrategyUpdate(row) {
      strategyUpdate({
        ids: this.selectedRows.map(item => item.id),
        strategy_id: row.id
      }).then(res => {
        this.$message({
          type: 'success',
          message: '切换策略成功'
        })
        this.strategyDialog = false
        this.getList()
      })
    },
    handleProxy() {
      var items = []
      for (let i = 0; i < this.selectedRows.length; i++) {
        items.push(this.selectedRows[i].id)
      }
      if (this.selectedRows.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择需要更新的任务'
        })
        return
      }
      ProxyList({
        page: 1,
        limit: 100
      }).then(response => {
        this.proxyRows = response.data.items.map(item => {
          // 截取 proxy_ip 最后的数字
          const ipParts = item.proxy_ip.split('.')
          const ipSuffix = ipParts[ipParts.length - 1]
          return { ...item, ipSuffix }
        })
        if (this.proxyRows.length === 0) {
          this.$message({
            type: 'warning',
            message: '没有可用的代理，请先添加代理'
          })
          this.proxyDialog = false
          return
        }
        this.proxyDialog = true
      })
    },
    handleProxyUpdate(row) {
      var ids = this.selectedRows.map(item => item.id)
      if (ids.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择需要更新的任务'
        })
        return
      }
      this.$confirm('此操作将切换代理, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        proxyUpdate({
          ids: ids,
          proxy_id: row.id
        }).then(response => {
          this.$message({
            type: 'success',
            message: '切换代理成功'
          })
          this.proxyDialog = false
          this.getList()
        })
      })
    },
    handleClear() {
      var items = []
      for (var i = 0; i < this.selectedRows.length; i++) {
        items.push(this.selectedRows[i].id)
      }
      if (this.selectedRows.length > 0) {
        this.$confirm('此操作将永久删除, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          clearTask({ ids: items }).then(response => {
            this.$message({
              type: 'success',
              message: '删除成功'
            })
            this.getList()
          })
        })
      }
    },
    handlePageChange(page) {
      this.loading = true
      this.listQuery.page = page
      taskList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSizeChange(val) {
      this.loading = true
      this.listQuery.page = 1
      this.listQuery.limit = val
      taskList(this.listQuery).then(response => {
        this.list = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    formatter(row, column, cellValue) {
      return (Number(cellValue) || 0).toFixed(2) + '%'
    }
  }
}
</script>

<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
@media only screen and (max-width: 470px) {
  .input-with-select{
    width: 98% !important;
    margin-top: 5px!important;
    margin-bottom: 5px!important;
  }
  .select-type{
    margin-top: 5px!important;
  }
}
</style>
