<template>
  <div class="app-container">
    <el-dialog
      title="切换代理"
      width="35%"
      :visible.sync="visible"
      :before-close="handleClose"
    >
      <div class="proxy-header">
        <el-row>
          <el-col :span="16">
            <div class="group_list">
              <div
                v-for="(group,index) in groups"
                :key="index"
                class="group_item"
                :class="{ 'active': activeIndex === index }"
                @click="selectGroup(index)"
              >{{ group.name }}</div>
            </div>
          </el-col>
          <el-col :span="8">
            <el-input
              v-model="keywords"
              placeholder="搜索地区名称"
              @clear="handleClear"
              @keyup.enter.native="handleSearch"
            >
              <el-button slot="append" type="primary" icon="el-icon-search" @click="handleSearch" />
            </el-input>
          </el-col>
        </el-row>
      </div>
      <ul class="proxy-list">
        <li v-if="currentList.length === 0" class="no-data">未找到数据</li>
        <li v-for="(item, index) in currentList" :key="index" class="proxy-item" @click="handleSelect(item)">
          <p>{{ item.name }}</p>
          <small>可用【{{ item.usercount }}】</small>
        </li>
      </ul>
    </el-dialog>
  </div>
</template>
<script>
export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    items: {
      type: Array,
      default: () => []
    },
    groups: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      activeIndex: 0,
      keywords: '',
      originalItems: []
    }
  },
  computed: {
    currentList() {
      // 优先使用已经保存的 originalItems（若存在），否则回退到 props items
      const source = (this.originalItems && this.originalItems.length) ? this.originalItems : (this.items || [])
      if (!source || source.length === 0 || this.activeIndex < 0 || this.activeIndex >= source.length) return []
      const list = source[this.activeIndex].List || []
      const kw = (this.keywords || '').trim().toLowerCase()
      if (!kw) return list
      // 不要修改原始 list，直接返回过滤结果
      return list.filter(item => (item && item.name && item.name.toString().toLowerCase().includes(kw)))
    }
  },
  watch: {
    items: {
      handler(newVal) {
        this.originalItems = (newVal || []).map(g => ({
          // 只做一层拷贝：group 对象浅拷贝 + List 数组浅拷贝
          ...g,
          List: Array.isArray(g.List) ? g.List.slice() : []
        }))
      },
      immediate: true,
      deep: true
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
      this.keywords = ''
    },
    selectGroup(index) {
      this.activeIndex = index
      this.keywords = ''
    },
    handleSearch() {
      this.keywords = (this.keywords || '').trim()
    },
    handleClear() {
      this.keywords = ''
    },
    handleSelect(item) {
      this.$emit('select', item)
      this.handleClose()
    }
  }
}
</script>

<style scoped lang="scss">
.proxy-header{
  margin-bottom: 10px;
  .group_list{
    display: block;
    margin-bottom: 10px;
    .group_item{
      display: inline-block;
      padding: 10px 20px;
      border-radius: 5px;
      margin-right: 10px;
      overflow: hidden;
      cursor: pointer;
      background-color: grey;
      color: white;
    }
    .active{
      background-color: #1890ff;
      color: white;
    }
  }
}
::v-deep {
  .el-dialog__body {
    padding-top: 8px !important;
    padding-bottom: 10px !important;
  }
}
ul{
  list-style: none;
  max-height: 480px!important;
  overflow-x: hidden;
  padding: 0;
  margin: 0;
  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none; /* 完全隐藏滚动条组件 */
  }
  li{
    width: calc(22% - 20px);
    padding: 10px 6px;
    display: inline-block;
    cursor: pointer;
    &:hover {
      background-color: #f5f5f5;
    }
    p {
      margin: 0;
      font-size: 12px;
      color: #000000;
    }
    small{
      color: #888888;
    }
  }
}
</style>
