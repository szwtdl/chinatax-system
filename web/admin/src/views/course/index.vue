<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.username"
              placeholder="请输入课程名称"
              clearable
              style="width: 200px;margin-right: 10px"
            />
            <el-input
              v-model="listQuery.username"
              placeholder="请输入用户账号"
              clearable
              style="width: 200px;margin-right: 10px"
            />
            <el-select v-model="listQuery.status" placeholder="请选择状态" style="width: 120px;margin-right: 10px;">
              <el-option label="学习中" value="published" />
              <el-option label="待学习" value="published" />
              <el-option label="认证中" value="published" />
              <el-option label="禁用" value="disabled" />
            </el-select>
            <el-select v-model="listQuery.status" placeholder="请选择状态" style="width: 120px;margin-right: 10px;">
              <el-option label="学习中" value="published" />
              <el-option label="待学习" value="published" />
              <el-option label="认证中" value="published" />
              <el-option label="禁用" value="disabled" />
            </el-select>
            <el-button type="primary" icon="el-icon-search">
              搜索
            </el-button>
            <el-button type="primary" icon="el-icon-edit-outline">
              重置
            </el-button>
            <el-button type="primary" icon="el-icon-edit-outline">
              排序
            </el-button>
            <el-button type="primary" icon="el-icon-edit-outline">
              导出
            </el-button>
          </div>
          <el-table :data="list" border fit highlight-current-row style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column label="课程ID" width="150px" align="center">
              <template slot-scope="scope">
                {{ scope.row.id }}
              </template>
            </el-table-column>
            <el-table-column label="课程名称" width="180px" align="center">
              <template slot-scope="scope">
                {{ scope.row.username }}
              </template>
            </el-table-column>
            <el-table-column label="用户账号" align="center">
              <template slot-scope="scope">
                {{ scope.row.username }}
              </template>
            </el-table-column>
            <el-table-column label="学习进度" align="center">
              <template slot-scope="scope">
                {{ scope.row.username }}
              </template>
            </el-table-column>
            <el-table-column label="当前状态" align="center">
              <template slot-scope="scope">
                {{ scope.row.username }}
              </template>
            </el-table-column>
            <el-table-column align="center" label="操作" width="120">
              <template slot-scope="{row}">
                <el-button
                  v-if="row.edit"
                  type="success"
                  size="small"
                  icon="el-icon-circle-check-outline"
                  @click="confirmEdit(row)"
                >
                  Ok
                </el-button>
                <el-button
                  v-else
                  type="primary"
                  size="small"
                  icon="el-icon-edit"
                  @click="row.edit=!row.edit"
                >
                  Edit
                </el-button>
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
  name: 'Course',
  data() {
    return {
      list: null,
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 10
      }
    }
  },
  methods: {
    dialogPvVisible() {
      this.$emit('dialogPvVisible')
    }
  }
}
</script>
<style scoped>
.edit-input {
  padding-right: 100px;
}

.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
