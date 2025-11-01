<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入内容"
              class="input-with-select"
              style="margin-right: 5px;"
            >
              <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
            </el-input>
            <el-button-group>
              <el-button type="primary" @click="visible = true">添加分类</el-button>
            </el-button-group>
          </div>
          <el-table :data="list">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column label="分类ID" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="分类名称" align="left">
              <template slot-scope="scope">
                {{ scope.row.name }}
              </template>
            </el-table-column>
            <el-table-column label="分类图片" align="left">
              <template slot-scope="scope">
                <img :src="scope.row.image" alt="分类图片" style="width: 100px; height: 100px;">
              </template>
            </el-table-column>
            <el-table-column label="父级分类" align="left">
              <template slot-scope="scope">
                {{ scope.row.parent_id }}
              </template>
            </el-table-column>
            <el-table-column label="创建时间" width="180px" align="center">
              <template slot-scope="scope">
                {{ scope.row.created_at }}
              </template>
            </el-table-column>
            <el-table-column label="更新时间" width="180px" align="center">
              <template slot-scope="scope">
                {{ scope.row.updated_at }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="180px" align="center">
              <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handlerEdit(scope.row)">编辑</el-button>
                <el-button type="danger" size="mini" @click="handlerDelete(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <create-category :visible.sync="visible" @save="handlerSave" />
    <edit-category :visible.sync="dialogEditVisible" :data="editData" @save="handlerSave" />
  </div>
</template>
<script>

import CreateCategory from './components/CreateCategory.vue'
import EditCategory from './components/EditCategory.vue'

export default {
  name: 'Category',
  components: { CreateCategory, EditCategory },
  data() {
    return {
      visible: false,
      dialogEditVisible: false,
      list: [],
      editData: {},
      listQuery: {
        page: 1,
        limit: 20,
        total: 0,
        keyword: ''
      }
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      console.log('get list')
    },
    handleSearch() {
      console.log('search')
    },
    handlerSave(item) {
      console.log(item)
    },
    handlerEdit(row) {
      this.editData = row
      this.dialogEditVisible = true
    },
    handlerDelete(row) {
      console.log(row)
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
