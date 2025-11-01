<template>
  <div class="app-container">
    <el-dialog
      title="切换代理"
      max-width="480px"
      :visible.sync="visible"
      :before-close="handleClose"
    >
      <ul class="proxy-list">
        <li v-for="(item, ind) in items" :key="ind" class="proxy-item" @click="handleSelect(item)">
          <h3>{{ item.region_name }} <span>剩余:{{ item.proxy_hours }}</span></h3>
          <p v-if="item.proxy_domain !==''">地址: {{ item.proxy_domain }}:{{ item.proxy_port }}</p>
          <p v-else>地址: {{ item.proxy_ip }}:{{ item.proxy_port }}</p>
          <p>账号: {{ item.proxy_user }}，密码：{{ item.proxy_pass }}</p>
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
  .el-dialog{
    width: 480px !important;
    .el-dialog__body {
      padding-top: 10px !important;
      padding-bottom: 10px !important;
    }
  }
}
ul{
  list-style: none;
  height: 450px;
  overflow-x: auto;
  padding: 0;
  margin: 0;
  scrollbar-width: none; /* Firefox 专属属性：隐藏滚动条（none = 完全隐藏，thin = 细滚动条） */
  -ms-overflow-style: none; /* IE/Edge 旧版（可选，现代浏览器可省略） */
  ::-webkit-scrollbar {
    width: 0; /* 隐藏滚动条宽度（纵向滚动条控制 width） */
    height: 0; /* 若有横向滚动，隐藏横向滚动条高度（可选） */
  }
  .proxy-item {
    padding: 10px;
    border-bottom: 1px solid #eaeaea;
    cursor: pointer;
    &:hover {
      background-color: #f5f5f5;
    }
    &:last-child {
      border-bottom: none;
    }
    h3 {
      width: 100%;
      margin: 0;
      font-size: 14px;
      font-weight: bold;
      color: #333;
      span{
        font-size: 12px;
        float: right;
        color: #29ac8b;
      }
    }
    p {
      margin: 5px 0;
      color: #666;
      font-size: 12px!important;
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
