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
          </div>
          <el-table :data="list" size="mini" border style="width: 100%">
            <el-table-column type="selection" align="center" width="60" />
            <el-table-column
              align="left"
              label="ID"
              width="180"
              element-loading-text="请给我点时间！"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.id }}</span>
              </template>
            </el-table-column>
            <el-table-column label="登录账号">
              <template slot-scope="scope">
                <el-input v-model="scope.row.name" size="mini" />
              </template>
            </el-table-column>
            <el-table-column label="机构名称">
              <template slot-scope="scope">
                <el-input v-model="scope.row.org_name" size="mini" />
              </template>
            </el-table-column>
            <el-table-column label="认证头像">
              <template slot-scope="scope">
                <el-input v-model="scope.row.face" size="mini" />
              </template>
            </el-table-column>
            <el-table-column label="认证状态">
              <template slot-scope="scope">
                <el-input v-model="scope.row.status" size="mini" />
              </template>
            </el-table-column>
            <el-table-column label="操作">
              <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="createFace(scope.row)">Create</el-button>
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
  name: 'Face',
  data() {
    return {
      list: [],
      listQuery: {
        page: 1,
        limit: 20,
        keywords: '',
        filter: ''
      },
      options: [
        {
          value: 'username',
          label: '用户账号'
        },
        {
          value: 'org_name',
          label: '机构名称'
        }
      ]
    }
  },
  methods: {
    createFace() {
      console.log('createFace')
    }
  }
}
</script>
<style scoped lang="scss">
.input-with-select {
  width: 450px;
}
</style>
