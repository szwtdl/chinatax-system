<template>
  <div class="app-container">
    <el-dialog
      title="选择照片"
      width="50%"
      :visible.sync="visible"
      :before-close="handleClose"
    >
      <ul class="albums-list">
        <li v-for="(item,ind) in items" :key="ind" class="albums-item" @click="handleSelect(item)">
          <el-image :src="item.url" />
        </li>
      </ul>
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
import { fetchList } from '@/api/photo'

export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    member: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      items: [],
      listQuery: {
        user_id: 0,
        total: 0,
        page: 1,
        limit: 28
      }
    }
  },
  watch: {
    member(val) {
      if (val) {
        this.getList()
      }
    }
  },
  methods: {
    getList() {
      fetchList({
        user_id: this.member,
        page: this.listQuery.page,
        limit: this.listQuery.limit
      }).then(response => {
        this.items = response.data.items
        this.listQuery.total = response.data.total
        this.listQuery.page = response.data.page
        this.listQuery.limit = response.data.page_size
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
.albums-list {
  display: flex;
  flex-wrap: wrap;
  padding: 0;
  margin: 0;
  list-style: none;

  .albums-item {
    width: 156px;
    margin: 10px;
    cursor: pointer;

    :hover {
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    image {
      width: 100%;
      height: 100%;
    }
  }
}
</style>
