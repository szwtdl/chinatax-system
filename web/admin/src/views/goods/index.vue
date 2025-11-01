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
              <el-button type="primary" @click="visible = true">添加商品</el-button>
            </el-button-group>
          </div>
          <el-table :data="list">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="center"
              label="商品ID"
              width="150"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="商品名称" align="left">
              <template slot-scope="scope">
                {{ scope.row.name }}
              </template>
            </el-table-column>
            <el-table-column label="商品图片" align="left">
              <template slot-scope="scope">
                <img :src="scope.row.image" alt="商品图片" style="width: 100px; height: 100px;">
              </template>
            </el-table-column>
            <el-table-column label="商品分类" align="left">
              <template slot-scope="scope">
                {{ scope.row.category_name }}
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
  </div>
</template>
<script>
export default {
  name: 'Goods',
  data() {
    return {
      visible: false,
      list: [],
      listQuery: {
        page: 1,
        limit: 20,
        total: 0,
        keyword: ''
      }
    }
  },
  methods: {
    handlerEdit(row) {
      console.log(row)
    },
    handleSearch() {
      console.log('search')
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
