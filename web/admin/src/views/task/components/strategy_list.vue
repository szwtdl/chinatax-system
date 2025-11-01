<template>
  <div class="app-container">
    <el-dialog
      title="选择策略"
      maxwidth="480px"
      :visible.sync="visible"
      :before-close="handleClose"
    >
      <ul class="strategy-list">
        <li v-for="(item,ind) in items" :key="ind" class="strategy-item" @click="handleSelect(item)">
          <h3>{{ item.platform_name }}</h3>
          <p>学习时长: {{ item.hours }} 【<small> {{ item.name }} 触发频率 {{ item.interval }} {{ item.unit }} </small>】 </p>
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
    }
  },
  methods: {
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSelect(item) {
      this.$emit('select', item)
      this.handleClose()
    }
  }
}
</script>
<style scoped lang="scss">
::v-deep {
  .el-dialog {
    width: 480px !important;
    .el-dialog__body {
      padding-top: 10px !important;
      padding-bottom: 10px !important;
    }
  }
}
ul{
  list-style: none;
  padding: 0;
  margin: 0;
  height: 380px;
  overflow-y: auto;
  scrollbar-width: none; /* Firefox 专属属性：隐藏滚动条（none = 完全隐藏，thin = 细滚动条） */
  -ms-overflow-style: none; /* IE/Edge 旧版（可选，现代浏览器可省略） */
  ::-webkit-scrollbar {
    width: 0; /* 隐藏滚动条宽度（纵向滚动条控制 width） */
    height: 0; /* 若有横向滚动，隐藏横向滚动条高度（可选） */
  }
  .strategy-item{
    padding: 3px 0;
    border-bottom: 1px solid #eaeaea;
    cursor: pointer;
    &:hover {
      background-color: #2de611;
      color: #fff;
      h3, p {
        color: #fff;
      }
    }
    &:last-child{
      border-bottom: none;
    }
    h3{
      padding: 5px 10px!important;
      font-size: 14px;
      color: #4A9FF9;
      font-weight: 600;
      margin:0;
    }
    p{
      padding: 0 10px!important;
      margin: 3px 0!important;
      color: #888;
      font-size: 12px;
    }
  }
}
@media only screen and (max-width: 470px) {
  ::v-deep{
    .el-dialog {
      width: 90% !important;
    }
  }
}
</style>
