<template>
  <div class="app-container">
    <el-dialog
      title="文件管理"
      class="dialog-list"
      width="1340px"
      :visible.sync="upload.outerVisible"
    >
      <el-dialog
        max-width="600px"
        title="文件上传"
        class="dialog-upload"
        :visible.sync="upload.innerVisible"
        append-to-body
        @close="handleClose"
      >
        <el-upload
          ref="upload"
          action="#"
          accept="image/*,video/*"
          list-type="picture-card"
          multiple
          :limit="10"
          :auto-upload="false"
          :file-list="upload.fileList"
          :http-request="customUpload"
          :on-exceed="handleExceed"
          :on-preview="handlePreview"
          :on-change="handleChange"
          :on-remove="handleRemove"
        >
          <i class="el-icon-plus" />
        </el-upload>
        <span slot="footer" class="dialog-footer">
          <el-button @click="upload.innerVisible = false">取 消</el-button>
          <el-button type="primary" @click="submitUpload">确 定</el-button>
        </span>
      </el-dialog>
      <el-dialog
        width="550px"
        title="添加分类"
        :visible.sync="upload.dialogMenuVisible"
        append-to-body
        @close="handleMenuClose"
      >
        <el-form :model="upload">
          <el-form-item label="父级分类" :label-width="formLabelWidth">
            <el-select :value="upload.parent_id" style="width: 100%;" @change="handleMenuChange">
              <el-option
                v-for="item in upload.category_list"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="分类名称" :label-width="formLabelWidth">
            <el-input v-model="upload.category_name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="分类描述" :label-width="formLabelWidth">
            <el-input v-model="upload.category_marker" autocomplete="off" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="handleMenuClose">取 消</el-button>
          <el-button type="primary" @click="saveCategory">确 定</el-button>
        </span>
      </el-dialog>
      <div class="media-box">
        <slot name="menu">
          <div class="media-menu">
            <ul>
              <li
                v-for="item in upload.menus"
                :key="item.id"
                :class="{ active: item.selected === true, inactive: item.selected === false }"
                @click="handleSwitchMenu(item.id)"
              >{{ item.name }}
                <div class="menu-control">
                  <el-tag
                    size="mini"
                    type="primary"
                    style="margin-right: 5px;"
                    @click="handleMenuUpdate(item.id,$event)"
                  >编辑</el-tag>
                  <el-tag size="mini" type="danger" @click="handleMenuDelete(item.id,$event)">删除</el-tag>
                </div>
              </li>
              <li class="add-menu" @click="upload.dialogMenuVisible = true">添加分类</li>
            </ul>
          </div>
        </slot>
        <div class="media-body">
          <slot name="header">
            <div class="media-body-header">
              <div class="media-body-header-search">
                <el-input
                  v-model="upload.keywords"
                  size="small"
                  class="input-with-select"
                  style="margin-right: 5px;width: 340px; margin-left: 15px;"
                >
                  <el-button slot="append" type="primary" @click="handleSearch">搜索</el-button>
                </el-input>
                <el-button-group>
                  <el-button type="primary" size="small" @click="upload.innerVisible = true">上传文件</el-button>
                  <el-button type="primary" size="small" @click="handleSelectAll">全选文件</el-button>
                  <el-button type="danger" size="small" @click="handleSelectImageDelete">批量删除</el-button>
                </el-button-group>
                <el-tag size="medium" type="primary" style="margin-left: 10px;">选择【{{
                  upload.member.name
                }}-{{ upload.member.card }}】
                </el-tag>
              </div>
            </div>
          </slot>
          <slot name="body">
            <ul class="media-list">
              <li
                v-for="item in upload.items"
                :key="item.id"
                :class="{ active: item.selected === true, inactive: item.selected === false }"
                @click="handleSelectImage(item)"
              >
                <div @click.stop="item.selected = !item.selected">
                  <video v-if="item.ext === '.mp4'" :src="item.full_url">-->
                    您的浏览器不支持 video 标签。
                  </video>
                  <el-image v-else :src="item.full_url" />
                  <el-checkbox
                    v-model="item.selected"
                    class="albums-checkbox"
                  />
                </div>
                <el-tag
                  v-if="item.ext === '.mp4'"
                  type="success"
                  class="cut_video"
                  @click="handleCutVideo(item,$event)"
                >
                  切图片
                </el-tag>
                <!--                <p class="media-item-title">{{ item.original_name }}</p>-->
              </li>
            </ul>
            <el-pagination
              background
              layout="total, prev, pager, next, jumper"
              :total="upload.total"
              :current-page.sync="upload.page"
              :page-size.sync="upload.pageSize"
              style="margin-top: 20px; text-align: right;"
              @size-change="handleSizeChange"
              @current-change="handlePageChange"
            />
          </slot>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getFileList, getMediaList } from '@/api/media'

export default {
  props: {
    upload: {
      type: Object,
      default: () => {
        return {
          keywords: '',
          page: 0,
          pageSize: 0,
          total: 0,
          outerVisible: false,
          innerVisible: false,
          dialogImageUrl: '',
          dialogVisible: false,
          dialogMenuVisible: false,
          fileList: [],
          menus: [],
          member: {
            name: '',
            card: ''
          },
          items: [],
          category_list: [],
          category_id: 1,
          parent_id: 0,
          action: 'add',
          category_name: '',
          category_marker: ''
        }
      }
    },
    accept: {
      type: String,
      default: 'image/*,video/*'
    },
    visible: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      formLabelWidth: '120px',
      listQuery: {
        keywords: '',
        page: 1,
        limit: 28,
        total: 0
      }
    }
  },
  watch: {
    accept: {
      handler(val) {
        this.accept = val
      },
      deep: true
    }
  },
  // updated() {
  //   this.init()
  // },
  methods: {
    init() {
      // 文件列表
      getFileList({
        category_id: 1,
        keywords: this.listQuery.keywords,
        page: this.listQuery.page,
        limit: this.listQuery.limit
      }).then(response => {
        this.items = response.data.items
        this.listQuery = {
          page: response.data.page,
          limit: response.data.page_size,
          total: response.data.total
        }
      })
      // 分类列表
      getMediaList().then(response => {
        this.menus = response.data
      })
    },
    handleSwitchMenu(item) {
      this.$emit('switch', item)
    },
    handleMenuChange(val) {
      this.$emit('change', val)
    },
    handleMenuDelete(val, event) {
      if (event) {
        event.preventDefault() // 阻止默认事件
        event.stopPropagation() // 阻止事件冒泡（如果需要）
      }
      this.$emit('menu_delete', val)
    },
    handleMenuUpdate(val, event) {
      if (event) {
        event.preventDefault() // 阻止默认事件
        event.stopPropagation() // 阻止事件冒泡（如果需要）
      }
      this.upload.dialogMenuVisible = true
      this.upload.category_list.map((item, index) => {
        if (item.value === val) {
          this.upload.category_id = item.value
          this.upload.category_name = item.label
          this.upload.action = 'edit'
        }
      })
      this.$emit('menu_update', val)
    },
    saveCategory() {
      const data = {
        category_id: this.upload.category_id,
        parent_id: this.upload.parent_id,
        category_name: this.upload.category_name,
        category_marker: this.upload.category_marker,
        action: this.upload.action
      }
      this.$emit('save', data)
    },
    handleCutVideo(item, event) {
      if (event) {
        event.preventDefault() // 阻止默认事件
        event.stopPropagation() // 阻止事件冒泡（如果需要）
      }
      this.$emit('cutVideo', item)
    },
    handleSelectImage(item) {
      this.$emit('select', item)
    },
    handleSizeChange(size) {
      this.$emit('changeSize', size)
    },
    handlePageChange(page) {
      this.$emit('changePage', page)
    },
    customUpload({ file }) {
      this.$emit('upload', file)
    },
    handleSelectImageSave() {
      this.$emit('selectImageSave')
    },
    handleSelectAll() {
      this.$emit('selectAll')
    },
    handleSelectImageDelete() {
      this.$emit('selectImageDelete')
    },
    handleMenuClose() {
      if (this.upload.action === 'edit') {
        this.upload.category_id = 1
        this.upload.parent_id = 0
        this.upload.category_name = ''
        this.upload.category_marker = ''
        this.upload.action = 'add'
      }
      this.$emit('menuClose')
    },
    handleClose() {
      this.$emit('close')
    },
    handleSearch() {
      this.$emit('search')
    },
    dialogFormVisible() {
      this.$emit('submit')
    },
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handleChange(file, fileList) {
      if (file.raw.type.indexOf('video') !== -1) {
        const video = document.createElement('video')
        video.src = URL.createObjectURL(file.raw)
        video.muted = true
        video.autoplay = true
        video.onloadeddata = () => {
          const canvas = document.createElement('canvas')
          canvas.width = video.videoWidth
          canvas.height = video.videoHeight
          const ctx = canvas.getContext('2d')
          ctx.drawImage(video, 0, 0, canvas.width, canvas.height)
          const index = fileList.findIndex(f => f.uid === file.uid)
          fileList[index].url = canvas.toDataURL('image/png')
        }
      }
    },
    handlePreview(file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    handleExceed(files, fileList) {
      this.$message.warning(`当前限制选择 10 个文件，本次选择了 ${files.length} 个文件`)
    },
    submitUpload() {
      this.$refs.upload.submit()
    }
  }
}
</script>
<style scoped lang="scss">
.media-box {
  display: flex;
  height: 580px;

  .media-menu {
    width: 200px;
    height: 100%;
    overflow-y: auto;
    border-right: 1px solid #f1f1f1;

    &::-webkit-scrollbar {
      width: 5px;
    }

    &::-webkit-scrollbar-thumb {
      background: #99c9f6; /* 滑块颜色 */
      border-radius: 10px; /* 圆角 */
    }

    &::-webkit-scrollbar-thumb:hover {
      background: #99c9f6; /* 滑块悬停颜色 */
    }

    /* 滚动条轨道 */
    &::-webkit-scrollbar-track {
      background: #f1f1f1; /* 轨道颜色 */
    }

    /* 滚动条轨道（当用户有阴影需求时） */
    &::-webkit-scrollbar-track:hover {
      background: #ddd; /* 悬停时轨道颜色 */
    }

    &::-webkit-scrollbar-button {
      display: none; /* 隐藏滚动条按钮 */
    }

    ul {
      list-style: none;
      padding: 0 10px;

      li {
        padding: 10px;
        cursor: pointer;
        border-bottom: 1px solid #eee;
        position: relative;

        .menu-control {
          float: right;
          position: absolute;
          top: 8px;
          display: none;
          right: 5px;
        }

        &:hover {
          .menu-control {
            display: block;
          }
        }
      }

      .add-menu {
        text-align: center;
        color: #1890ff;
        padding: 8px 20px !important;
        border: 1px solid #1890ff;
        margin-top: 10px;
        border-radius: 4px;
      }

      .active {
        background-color: #1890ff;
        color: white;
        font-weight: bold;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
      }
    }
  }

  .media-body {
    flex: 3;
    margin-left: 10px;

    ul.media-list {
      list-style: none;
      padding: 0 10px;
      height: 460px;
      overflow-y: auto;

      &::-webkit-scrollbar {
        width: 5px;
      }

      &::-webkit-scrollbar-thumb {
        background: #99c9f6; /* 滑块颜色 */
        border-radius: 10px; /* 圆角 */
      }

      &::-webkit-scrollbar-thumb:hover {
        background: #99c9f6; /* 滑块悬停颜色 */
      }

      /* 滚动条轨道 */
      &::-webkit-scrollbar-track {
        background: #f1f1f1; /* 轨道颜色 */
      }

      /* 滚动条轨道（当用户有阴影需求时） */
      &::-webkit-scrollbar-track:hover {
        background: #ddd; /* 悬停时轨道颜色 */
      }

      &::-webkit-scrollbar-button {
        display: none; /* 隐藏滚动条按钮 */
      }

      li {
        width: 140px;
        height: 140px;
        float: left;
        border: 1px solid #eee;
        cursor: pointer;
        margin: 0 5px 10px;
        position: relative;

        .media-item-title {
          position: absolute;
          bottom: 0;
          padding: 5px 10px;
          background-color: rgba(0, 0, 0, 0.5); /* 黑色，50%透明 */
          display: none;
          color: #fff;
          margin: 0;
        }
        div{
          width: 100%;
          height: 100%;
          position: relative;
          .albums-checkbox{
            position: absolute;
            top: 5px;
            right: 5px;
            z-index: 10;
          }
        }
        .cut_video {
          position: absolute;
          left: calc(50% - 25px);
          top: calc(50% - 10px);
          user-select: none;
          display: none;
        }

        .active {
          background-color: #1890ff;
          border: 1px solid #1890ff;
        }

        &:hover {
          .cut_video {
            display: block;
          }

          .media-item-title {
            display: block;
          }

        }

        video {
          width: 100%;
          height: 100%; /* 填充高度 */
          object-fit: cover;
        }

        img {
          width: 100%; /* 填充宽度 */
          height: 100%; /* 填充高度 */
          user-select: none;
          object-fit: contain; /* 保持图片比例 */
        }
      }
    }
  }
}
@media only screen and (max-width: 470px) {
  ::v-deep {
    .el-dialog{
      width: 90%;
    }
    .el-pagination span:not([class*=suffix]), .el-pagination button{
      margin-left:5px!important;
    }
    .el-button-group{margin-left: 10px}
    .input-with-select{margin-left: 10px!important;}
  }
  .input-with-select{
    width: 290px!important;
  }
  .media-box{
    .media-body{
      width: 100%;
      margin-left: 0;
      .media-list{
        height: 400px!important;
        display: flex;
        padding: 0!important;
        flex-wrap: wrap;
        li{
          flex: 0 0 calc(33.333% - 10px);  /* 一行3个 */
          text-align: center;
          box-sizing: border-box;
        }
      }
    }
  }
}
</style>
