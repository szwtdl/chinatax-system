<template>
  <div class="app-container">
    <el-dialog
      v-loading="loading"
      :visible.sync="visible"
      max-width="680px"
      :before-close="handleClose"
    >
      <template #title>
        <div class="custom-header">
          <p>选择课程</p>
        </div>
      </template>
      <ul class="course-box">
        <li v-for="(user, index) in items" :key="user.user_id + '-' + index">
          <div class="course-item-body">
            <p>机构: {{ user.org_name }} - 【<span style="color: red;">{{ user.nickname }} - {{ user.username }}</span>】</p>
          </div>
          <ul>
            <li
              v-for="item in user.items"
              :key="item.id"
              class="course-item"
              :class="{ active: item.order_status === 1, inactive: item.order_status === 0 }"
              @click="handleChange(item,user)"
            >
              <div class="course-item-photo">
                <img :src="item.cover" :alt="item.title">
              </div>
              <div class="course-item-body">
                <p>{{ item.title }}
                  <el-tag v-if="item.order_status ===1" type="success">正常</el-tag>
                  <el-tag v-else type="danger">到期</el-tag>
                </p>
                <p style="padding: 0;">
                  <el-tag style="margin-right: 5px;">已完成 {{ item.learning_progress }}%</el-tag>
                  <el-tag style="margin-right: 5px;">总章数 {{ item.duration_count }}</el-tag>
                  <el-tag>已学章数 {{ item.total_learning_time }}</el-tag>
                </p>
              </div>
            </li>
          </ul>
        </li>
      </ul>
      <ul v-if="selects.length" class="select-course">
        <li v-for="(item,index) in selects" :key="index" @click="handleDelete(index)">
          <span style="color: red;">{{ item.org_name }}</span> 【{{ item.nickname }} - {{ item.course_name }}】
        </li>
      </ul>
      <div class="el-card__footer">
        <el-button type="primary" @click="handleStart()">添加任务</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getAllCourse } from '@/api/student'
import { batchAddTask } from '@/api/task'
export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    users: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      loading: false,
      dialogOrg: false,
      checked: false,
      items: [],
      token: '',
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
  watch: {
    visible(newVal) {
      if (newVal) {
        this.handlerCourse()
      }
    }
  },
  methods: {
    handlerCourse() {
      this.loading = true
      this.platform_type = this.users[0].platform_type
      this.selects = []
      getAllCourse({
        type: this.platform_type.toString(),
        ids: this.users.map(item => item.id)
      }).then(res => {
        this.items = res.data
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleChange(item, user) {
      if (item.order_status === 0) {
        this.$message.error('课程已到期')
        return
      }
      if (item.learning_progress === 100) {
        return this.$message({
          type: 'warning',
          message: '已经学完不需要'
        })
      }
      const data = {
        user_id: user.user_id,
        account_id: user.account_id,
        nickname: user.nickname,
        username: user.username,
        password: user.password,
        token: user.token,
        org_id: user.org_id,
        region_id: user.region_id,
        org_name: user.org_name,
        course_id: item.id,
        course_name: item.title,
        cover: item.cover,
        order_id: item.order_id,
        platform_type: user.platform_type,
        device_id: user.device_id
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
    handleStart() {
      if (this.selects.length === 0) {
        return this.$message({
          type: 'error',
          message: '请选择课程'
        })
      }
      batchAddTask({
        items: this.selects,
        platform_type: this.platform_type.toString()
      }).then(() => {
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.$emit('update:visible', false)
      }).catch(err => {
        console.log(err)
      })
    },
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
<style lang="scss" scoped>
.select-course{
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
      font-weight: normal;
      font-size: 12px;
      padding: 8px 0;
      margin-right: 10px;
      cursor: pointer;
      display: none;
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
    max-height: 450px;
    overflow-y: auto;
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
      .course-item-checkbox{
        width: 25px;
        padding: 8px;
        line-height: 60px;
      }
      .course-item-photo {
        width: 130px;
        height: 90px;
        padding: 8px;

        img {
          width: 100%;
          height: 100%;
        }
      }

      .course-item-body {
        flex: 1;

        p {
          padding: 8px 0;
          margin: 0;
          font-size: 16px;
          font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
        }
      }
    }
  }

  .el-card__footer{
     height: 40px;
     button{
       float: right;
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
