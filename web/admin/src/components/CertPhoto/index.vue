<template>
  <div class="app-container">
    <el-dialog
      title="查看相册"
      max-width="980px"
      :visible.sync="visible"
      :before-close="handleClose"
    >
      <div class="albums">
        <div class="albums-header">
          <el-button-group>
            <el-button type="primary" size="small" @click="handleAll">全选</el-button>
            <el-button type="danger" size="small" @click="handleDelete">删除</el-button>
          </el-button-group>
        </div>
        <ul class="albums-list">
          <li v-for="(item,ind) in items" :key="ind" class="albums-item" @click="handleSelect(item)">
            <div @click.stop="item.checked = !item.checked">
              <el-image :src="item.url" />
              <el-checkbox
                v-model="item.checked"
                class="albums-checkbox"
              />
            </div>
          </li>
        </ul>
      </div>
      <el-pagination
        background
        layout="total, prev, pager, next, jumper"
        :total="listQuery.total"
        :current-page.sync="listQuery.page"
        :page-size.sync="listQuery.limit"
        style="margin-top: 20px; text-align: right;"
        @current-change="handlePageChange"
      />
    </el-dialog>
  </div>
</template>
<script>
import { fetchList, remove } from '@/api/photo'

export default {
  props: {
    member: {
      type: Number,
      default: 0
    },
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      items: [],
      listQuery: {
        user_id: 0,
        total: 0,
        page: 1,
        limit: 24
      }
    }
  },
  watch: {
    member(val) {
      if (val) {
        this.listQuery.page = 1
        this.items = []
        this.getList()
      }
    }
  },
  methods: {
    getList() {
      if (!this.member) {
        this.$message.error('请选择用户')
        return
      }
      fetchList({
        user_id: this.member,
        page: this.listQuery.page,
        limit: this.listQuery.limit
      }).then(response => {
        this.items = response.data.items
        // 循环遍历，添加选中状态
        this.items.forEach(item => {
          this.$set(item, 'checked', false)
        })
        this.listQuery.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
      })
    },
    handleDelete() {
      const ids = this.items.filter(item => item.checked).map(item => item.id)
      if (ids.length === 0) {
        this.$message.error('请选择要删除的图片')
        return
      }
      remove({ items: ids }).then(() => {
        this.$message.success('删除成功')
        this.items = this.items.filter(item => !item.checked)
        if (this.items.length === 0) {
          this.listQuery.page = this.listQuery.page > 1 ? this.listQuery.page - 1 : 1
          this.getList()
        }
      })
    },
    handleAll() {
      this.items.forEach(item => {
        item.checked = !item.checked
      })
    },
    handleSelect(item) {
      this.$emit('select', item)
    },
    handleClose() {
      this.$emit('update:visible', false)
    },
    handlePageChange(page) {
      this.listQuery.page = page
      this.getList()
    }
  }
}
</script>
<style scoped lang="scss">
::v-deep {
  .el-dialog {
    width: 900px!important;
    min-height: 80%!important;
  }
}
.albums {
  max-height: 540px;
  .albums-header {
    text-align: right;
    border-bottom: 1px solid #ebeef5;
    padding-bottom: 10px;
  }

  .albums-list {
    max-height: 480px;
    display: grid;
    padding: 0;
    margin: 5px 0;
    grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
    justify-content: start;
    overflow-y: auto;
    list-style: none;
    .albums-item {
      height: 150px;
      max-width: 150px;
      margin:5px!important;
      float: left;
      cursor: pointer;
      position: relative;
      .albums-checkbox {
        position: absolute;
        top: 5px;
        right: 5px;
      }

      :hover {
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
      }

      div {
        width: 100%;
        height: 100%;
        display: block;
        img {
          width: 100%;
          height: 100%;
        }
      }
    }
  }
}
@media only screen and (max-width: 470px) {
  .albums-list {
    display: flex;
    flex-wrap: wrap;
    height: 400px;
    overflow-y: auto;
    .albums-item{
      flex: 0 0 calc(33.333% - 10px);
      text-align: center;
      margin:5px!important;
      box-sizing: border-box;
    }
  }
}
</style>
