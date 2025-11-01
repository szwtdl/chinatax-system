<template>
  <div class="app-container">
    <el-dialog
      :visible.sync="visible"
      max-width="650px"
      :before-close="handleClose"
    >
      <template #title>
        <div class="custom-header">
          <p>选择课程</p>
          <ul class="custom-header-org-list">
            <li>【 <span style="color: rgb(31 147 255);">{{ user.nickname }}-{{ user.username }}</span> 】</li>
          </ul>
        </div>
      </template>
      <ul class="course-box">
        <li
          v-for="(item,index) in items"
          :key="index"
          class="course-item"
          :class="{ active: item.order_status === 1, inactive: item.order_status === 0 }"
          @click="handleChange(item)"
        >
          <div v-if="item.org_name === '在线100分'" class="course-item-photo">
            <img :src="proxyCoverUrls[index]" :alt="item.title" :data-src="proxyCoverUrls[index]">
          </div>
          <div v-else class="course-item-photo">
            <img :src="item.cover" :alt="item.title">
            <span v-if="item.is_face">需要人脸</span>
            <span v-else class="green">无需人脸</span>
          </div>
          <div class="course-item-body">
            <p>{{ item.title }} - {{ item.teacher_name }}
              <el-tag v-if="item.order_status === 1" type="success">正常</el-tag>
              <el-tag v-else type="danger">到期</el-tag>
            </p>
            <p style="padding: 0;">
              <el-tag style="margin-right: 5px;">已完成 {{ item.learning_progress }}%</el-tag>
              <el-tag style="margin-right: 5px;">总章数 {{ item.duration_count }}</el-tag>
              <el-tag>已学章数 {{ item.total_learning_time }}</el-tag>
            </p>
            <p style="margin: 0; padding:6px 0; font-size: 12px;color: #909399;">{{ item.org_name }}</p>
          </div>
        </li>
      </ul>
      <ul v-if="selects.length" class="select-course">
        <li v-for="(item,index) in selects" :key="index" @click="handleDelete(index)">
          <span style="color: red;">{{ item.org_name }}</span> - {{ item.course_name }}
        </li>
      </ul>
      <span slot="footer" class="dialog-footer">
        <el-button type="danger" size="small" @click="handleStart('local')">本地学习</el-button>
        <el-button type="primary" size="small" @click="handleStart('live')">线上学习</el-button>
      </span>
      <!-- 切换机构 start -->
      <el-dialog
        title="切换机构"
        :visible.sync="dialogOrg"
        width="25%"
        :before-close="handleOrgClose"
        append-to-body
      >
        <ul class="org-list">
          <li
            v-for="item in org_list"
            :key="item.id"
            class="org-item"
            :class="{ active: item.id === org_id }"
            @click="item.id !== org_id && handleOrgChange(item)"
          >
            {{ item.name }}
          </li>
        </ul>
      </el-dialog>
      <!-- 切换机构 end -->
    </el-dialog>
  </div>
</template>
<script>
import { addTask } from '@/api/task'

export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    items: {
      type: Array, // 类型为数组
      default: () => [] // 默认值为空数组
    },
    user: {
      type: Object, // 类型为对象
      default: () => ({}) // 默认值为空对象
    }
  },
  data() {
    return {
      dialogOrg: false,
      checked: false,
      token: '',
      nickname: '',
      account_id: 0,
      platform_type: 0,
      device_id: 0,
      org_id: 0,
      org_name: '',
      region_id: 0,
      selects: [],
      org_list: []
    }
  },
  computed: {
    proxyCoverUrls() {
      const prefix = 'https://cdn.zaixian100f.com/'
      let proxyBase = ''
      // 判断是否在H5中
      if (typeof window !== 'undefined') {
        proxyBase = window.location.origin + '/proxy/'
      }
      return this.items.map(item => {
        if (item.cover && item.cover.startsWith(prefix)) {
          return proxyBase + item.cover.slice(prefix.length)
        }
        return item.cover
      })
    }
  },
  watch: {
    items: {
      handler(value) {
        this.selects = []
      },
      deep: false // 监听对象内部的变化
    }
  },
  methods: {
    handleOrgClose() {
      this.dialogOrg = false
    },
    handleChange(item) {
      if (item.order_status === 0) {
        this.$message.error('课程已到期')
        return
      }
      if (item.duration_count === item.total_learning_time && item.learning_progress === 100) {
        return this.$message({
          type: 'warning',
          message: '已经学完不需要'
        })
      }
      const data = {
        user_id: this.user.id,
        account_id: item.student_id,
        name: this.user.nickname,
        username: this.user.username,
        password: this.user.password,
        region_id: item.region_id,
        cover: item.cover,
        org_id: item.org_id,
        org_name: item.org_name,
        course_id: item.id,
        course_name: item.title,
        token: this.token,
        mobile: item.mobile,
        order_id: item.order_id,
        platform_type: this.user.platform_type,
        device_id: this.user.device_id
      }
      for (let i = 0; i < this.selects.length; i++) {
        if (this.selects[i].course_id === data.course_id && this.selects[i].org_id === data.org_id && this.selects[i].user_id === data.user_id) {
          return this.$message({
            type: 'warning',
            message: '已经选择了'
          })
        }
      }
      this.selects.push(data)
    },
    handleDelete(index) {
      this.selects.splice(index, 1)
    },
    handleStart(type) {
      if (this.selects.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择课程'
        })
      }
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      addTask({
        type: type,
        tasks: this.selects
      }).then(response => {
        loading.close()
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.$emit('update:visible', false)
      }).finally(() => {
        loading.close()
      })
    },
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
<style lang="scss" scoped>
::v-deep {
  .el-dialog {
    width: 600px;
  }
}
.select-course {
  list-style: none;
  padding: 0;
  margin: 0;

  li {
    padding: 3px 0;
    color: #409EFF;
  }
}

.custom-header {
  font-size: 16px;
  font-weight: bold;

  p {
    text-align: left;
    float: left;
    margin: 5px 20px 5px 0;
  }

  .custom-header-org-list {
    text-align: left;
    float: left;
    list-style: none;
    padding: 0;
    margin: 0;
    li {
      font-size: 12px;
      padding: 8px 0;
      margin-right: 10px;
      cursor: pointer;
      font-weight: bold;
      display: block;
    }

    .active {
      display: block;
      font-weight: bold;
      color: #409EFF;
    }
  }
}

.org-list {
  list-style: none;
  padding: 0;
  margin: 0;

  li {
    padding: 8px 0;
    cursor: pointer;
  }

  .active {
    color: #409EFF;
  }
}

.el-dialog__body {
  padding: 10px 20px;
  display: block;
  clear: both;

  .course-box {
    list-style: none;
    padding: 0;
    display: block;
    clear: both;

    .inactive {
      user-select: none;
      pointer-events: none;
      cursor: none;
    }

    .course-item {
      display: flex;
      user-select: none;
      text-align: left;
      margin-bottom: 8px;
      cursor: pointer;
      overflow: hidden;
      background-color: #f2f2f2;

      .course-item-checkbox {
        width: 25px;
        padding: 8px;
        line-height: 60px;
      }

      .course-item-photo {
        width: 150px;
        height: 100px;
        padding: 8px;
        position: relative;
        box-sizing: border-box;
        flex-shrink: 0;

        img {
          width: 100%;
          height: 100%;
        }
        span{
          position: absolute;
          bottom: 8px;
          left: 8px;
          right: 8px;
          background: rgb(230, 17, 35,0.6);
          color: #fff;
          padding: 4px 4px;
          text-align: center;
          font-size: 12px;
          border-radius: 2px;
        }
        .green{
          background: rgb(45, 230, 17,0.6);
          color: white;
        }
      }

      .course-item-body {
        flex: 1;
        p {
          padding: 5px 0;
          margin: 0;
          font-size: 14px;
          font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
          span{
            font-size: 12px!important;
          }
        }
      }
    }
  }
}

@media only screen and (max-width: 470px) {
  ::v-deep {
    .el-dialog {
      width: 96% !important;
      .el-dialog__body{
        padding: 10px;
        .course-box {
          .course-item {
            .course-item-photo {
              width: 110px;
              height: 90px;
              font-size: 12px!important;
            }
            .course-item-body{
              p {
                font-size: 14px;
              }
              span{
                font-size: 12px;
              }
              .el-tag--medium{
                padding: 0 6px;
                font-size: 10px!important;
              }
            }
          }
        }
      }
    }
  }
}
</style>
