<template>
  <div class="app-container">
    <el-row>
      <el-col :span="24">
        <el-card>
          <div class="filter-container">
            <!-- 新增的下拉框 -->
            <el-select
              v-model="listQuery.school_id"
              placeholder="请选择代理商"
              size="small"
              style="width: 150px; margin-right: 5px;"
            >
              <el-option
                v-for="item in schoolRows"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
            <el-input
              v-model="listQuery.keywords"
              placeholder="请输入内容"
              size="small"
              class="input-with-select"
              style="margin-right: 5px;"
              @keyup.enter.native="handleSearch"
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
            <el-button-group>
              <el-button type="primary" size="small" @click="dialogVisible = true">添加用户</el-button>
              <el-upload
                class="upload-btn"
                style="float: left;border: none;"
                accept=".xls,.xlsx"
                action="#"
                :before-upload="handleBeforeUpload"
              >批量导入
              </el-upload>
              <el-button type="primary" size="small" @click="openSwitchDialog">切换平台</el-button>
              <el-button type="danger" size="small" @click="handleDelete">批量删除</el-button>
            </el-button-group>
          </div>
          <el-table
            :data="list"
            size="mini"
            border
            style="width: 100%"
            @selection-change="handleUserChange"
          >
            <el-table-column type="selection" align="center" width="60px" />
            <el-table-column
              align="left"
              label="用户ID"
              width="100px"
              prop="id"
              element-loading-text="请给我点时间！"
            />
            <el-table-column label="会员头像" min-width="90px" prop="avatar" align="center">
              <template slot-scope="scope">
                <el-image
                  v-if="scope.row.avatar"
                  style="width: 30px;height: 30px;"
                  :lazy="true"
                  :src="scope.row.avatar"
                  :preview-src-list="[scope.row.avatar]"
                  fit="cover"
                />
                <el-image
                  v-else
                  style="width: 30px;height: 30px;"
                  :src="defaultImageUrl"
                  fit="cover"
                />
              </template>
            </el-table-column>
            <el-table-column label="会员姓名" prop="learn" sortable min-width="160px">
              <template slot-scope="scope">
                <div v-if="scope.row.learn">{{ scope.row.nickname }}-<span style="color: #4A9FF9;">【学习中】</span></div>
                <div v-else>{{ scope.row.nickname }}</div>
              </template>
            </el-table-column>
            <el-table-column label="会员账号" prop="username" min-width="160px" align="left">
              <template slot-scope="scope">
                <el-tooltip v-if="scope.row.invite_code" class="item" effect="dark" :content="scope.row.invite_code" placement="top">
                  <div v-if="scope.row.invite_code">{{ scope.row.username }}-<span style="color: #4A9FF9;">【激活码】</span></div>
                </el-tooltip>
                <div v-else>{{ scope.row.username }}</div>
              </template>
            </el-table-column>
            <el-table-column label="所属学校" prop="school_name" min-width="120px" align="left" />
            <el-table-column label="认证照片" prop="is_photo" min-width="130px" sortable align="center">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.is_photo === true" style="cursor: pointer;" @click="album(scope.row)">人脸照片</el-tag>
                <el-tag v-else type="danger">没有照片</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="账号平台" prop="platform_type" min-width="120px" sortable align="left">
              <template slot-scope="scope">
                <el-tag style="cursor: pointer;">
                  {{ getPlatformName(scope.row.platform_type) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="当前状态" align="center" min-width="80px">
              <template slot-scope="scope">
                <el-switch
                  v-model="scope.row.status"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="handleStatusChange(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="created_at" width="180px" align="center" />
            <el-table-column label="更新时间" prop="updated_at" width="180px" align="center" />
            <el-table-column align="center" label="操作" fixed="right" width="200">
              <template slot-scope="scope">
                <el-button-group>
                  <el-button
                    size="mini"
                    type="success"
                    @click="handleTask(scope.$index,scope.row)"
                  >选择课程
                  </el-button>
                  <el-dropdown @command="(command) => handleCommand(scope.row, scope.$index, command)">
                    <el-button size="mini" type="primary">操作<i class="el-icon-arrow-down el-icon--right" />
                    </el-button>
                    <el-dropdown-menu>
                      <el-dropdown-item command="upload">上传视频</el-dropdown-item>
                      <el-dropdown-item command="edit">编辑账号</el-dropdown-item>
                      <el-dropdown-item command="delete">删除账号</el-dropdown-item>
                    </el-dropdown-menu>
                  </el-dropdown>
                </el-button-group>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            :current-page.sync="listQuery.page"
            :page-size.sync="listQuery.limit"
            style="margin-top: 20px; text-align: center;"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
          />
        </el-card>
      </el-col>
    </el-row>
    <el-dialog
      title="新增用户"
      :visible.sync="dialogVisible"
      width="550px"
      :before-close="handleClose"
    >
      <el-form :model="form">
        <el-form-item label="登录账号" :label-width="formLabelWidth">
          <el-input v-model="form.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="账号平台" :label-width="formLabelWidth">
          <el-select v-model="form.platform_type" style="width: 100%;" @change="handlePlatformChange">
            <el-option
              v-for="platform in form.platform_list"
              :key="platform.id"
              :label="platform.name"
              :value="platform.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="form.sign !==''" label="验证码" :label-width="formLabelWidth">
          <el-input v-model="form.code" autocomplete="off" style="width: 180px;margin-right: 10px;" />
          <el-image
            style="width: 100px;height: 35px; position: absolute;top: 0;"
            :src="form.image"
            @click="getCodeImage"
          />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submitForm">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="编辑用户"
      :visible.sync="dialogEditVisible"
      width="550px"
      :before-close="handleClose"
    >
      <el-form :model="edit">
        <el-form-item label="用户姓名" :label-width="formLabelWidth">
          <el-input v-model="edit.nickname" autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录账号" :label-width="formLabelWidth">
          <el-input v-model="edit.username" readonly autocomplete="off" />
        </el-form-item>
        <el-form-item label="登录密码" :label-width="formLabelWidth">
          <el-input v-model="edit.password" autocomplete="off" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="editForm">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="切换平台"
      :visible.sync="switchVisible"
      width="550px"
      :before-close="handleCloseSwitch"
    >
      <ul class="platform-box">
        <li v-for="(item,index) in form.platform_list" :key="index" :class="['platform-item', activePlatformId === item.id ? 'active' : '']" @click="switchPlatform(item)">{{ item.name }}</li>
      </ul>
    </el-dialog>
    <media-upload
      :upload="upload"
      :visible.sync="upload.outerVisible"
      @upload="handleFile"
      @change="handleChangeMenu"
      @selectAll="handleSelectAll"
      @select="handleSelectionImage"
      @save="handleMenuSave"
      @search="handleSearchImages"
      @close="handleUploadClose"
      @switch="handleMenuSwitch"
      @menu_delete="handleMenuDelete"
      @menu_update="handleMenuUpdate"
      @menuColse="handleMenuClose"
      @cutVideo="handleCutVideo"
      @selectImageSave="handleSelectionImageSave"
      @selectImageDelete="handleSelectionImageDelete"
      @changeSize="handleUploadSizeChange"
      @changePage="handleUploadPageChange"
      @submit="handleUpload"
    />
    <course-switch :visible.sync="dialogCourseVisible" :user="user" :items="course_list" />
    <cert-photo :visible.sync="isPhoto" :member="photoId" />
    <SelectCourse :visible.sync="dialogSchoolVisible" :users="select" />
  </div>
</template>
<script>
import MediaUpload from '@/views/users/components/MediaUpload.vue'
import CertPhoto from '@/components/CertPhoto/index.vue'
import CourseSwitch from '@/views/users/components/CourseSwitch.vue'
import SelectCourse from '@/views/users/components/select_course.vue'

import * as ApiStudent from '@/api/student'
import * as ApiMedia from '@/api/media'
import * as ApiOrg from '@/api/org'
import * as ApiTask from '@/api/task'
import XLSX from 'xlsx'

export default {
  name: 'Users',
  components: { SelectCourse, MediaUpload, CertPhoto, CourseSwitch },
  data() {
    return {
      dialogVisible: false,
      dialogSchoolVisible: false,
      dialogEditVisible: false,
      switchVisible: false,
      dialogCourseVisible: false,
      loading: false,
      listLoading: true,
      isPhoto: false,
      photoId: 0,
      activePlatformId: null,
      defaultImageUrl: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80',
      formLabelWidth: '100px',
      user: {},
      list: [],
      course_list: [],
      options: [
        {
          value: 'username',
          label: '会员账号'
        },
        {
          value: 'nickname',
          label: '会员姓名'
        },
        {
          value: 'id',
          label: '用户ID'
        }
      ],
      schoolRows: [],
      listQuery: {
        page: 1,
        limit: 20,
        school_id: 0,
        keywords: '',
        filter: 'username'
      },
      total: 0,
      total_page: 0,
      pageSize: 0,
      region: '',
      regions: [],
      headers: [],
      excelData: [],
      select: [],
      upload: {
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
      },
      selectUser: {
        username: '',
        password: '',
        org_id: 0,
        region_id: 0,
        token: ''
      },
      form: {
        nickname: '',
        username: '',
        password: 'Dp111111',
        platform_type: 0,
        platform_list: [],
        org_id: 0,
        code: '',
        sign: '',
        image: '',
        is_select_org: true
      },
      edit: {
        id: null,
        nickname: '',
        username: '',
        password: '',
        org_id: 0
      }
    }
  },
  created() {
    this.getList()
    this.getOptions()
    this.getOrgList()
    this.getSchoolList()
  },
  methods: {
    openSwitchDialog() {
      if (this.select.length > 0) {
        this.switchVisible = true
      } else {
        this.$message.warning('请先选择平台')
      }
    },
    getPlatformName(id) {
      const item = this.form.platform_list.find(p => p.id === id)
      return item ? item.name : '待开放'
    },
    handleStatusChange(row) {
      const status = row.status ? 'true' : 'false'
      ApiStudent.status({ id: row.id, status }).then(() => {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.getList()
      }).finally(() => {
        this.loading = false
      })
    },
    getOptions() {
      ApiStudent.PlatformList().then(response => {
        const { data } = response
        this.form.platform_list = data
        this.form.platform_type = data.length > 0 ? data[0].id : 0
      })
    },
    getList() {
      this.loading = true
      ApiStudent.fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.pageSize = response.data.page_size
        this.total_page = response.data.total_page
        this.listQuery.page = response.data.page
        this.loading = false
      })
    },
    album(row) {
      this.isPhoto = true
      this.photoId = row.id
    },
    getOrgList() {
      this.loading = true
      ApiOrg.fetchOrgList(this.listQuery).then(response => {
        var items = []
        for (let i = 0; i < response.data.items.length; i++) {
          items.push({
            label: response.data.items[i].name,
            value: response.data.items[i].org_id
          })
        }
        this.regions = items
        this.region = response.data.items[0].name
        this.form.org_id = response.data.items[0].org_id
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    getSchoolList() {
      ApiStudent.schoolList(this.listQuery).then(response => {
        const items = response.data || []
        this.schoolRows = [
          { value: 0, label: '请选择代理' } // 添加默认选项
        ] // 可选：初始化清空
        items.forEach(item => {
          this.schoolRows.push({
            value: item.id,
            label: item.nickname
          })
        })
      })
    },
    getMediaType() {
      this.upload.menus = []
      this.upload.items = []
      this.upload.category_list = []
      ApiMedia.getMediaList({ page: 1, limit: 100 }).then(response => {
        var length = response.data.items.length
        this.upload.category_list.push({
          label: '顶级分类',
          value: 0
        })
        for (var i = 0; i < length; i++) {
          this.upload.category_list.push({
            label: response.data.items[i].name,
            value: response.data.items[i].id
          })
          this.upload.menus.push({
            id: response.data.items[i].id,
            name: response.data.items[i].name,
            selected: response.data.items[i].id === 1
          })
        }
      })
      ApiMedia.getFileList({
        page: 1,
        limit: 21,
        category_id: this.upload.category_id
      }).then(res => {
        this.upload.items = res.data.items.map(item => {
          item.selected = false // 修改拼写为 "selected"
          return item // 确保 map 返回修改后的对象
        })
        this.upload.total = res.data.total
        this.upload.pageSize = res.data.page_size
        this.upload.total_page = res.data.total_page
        this.upload.page = res.data.page
      })
    },
    handleBeforeUpload(file) {
      this.parseExcel(file)
      return false // 阻止默认上传行为
    },
    parseExcel(file) {
      const reader = new FileReader()
      reader.onload = (e) => {
        const data = new Uint8Array(e.target.result)
        const workbook = XLSX.read(data, { type: 'array' })
        const sheetName = workbook.SheetNames[0]
        const sheet = workbook.Sheets[sheetName]
        const jsonData = XLSX.utils.sheet_to_json(sheet, { header: 1 })
        this.headers = jsonData[0] || []
        this.excelData = jsonData.slice(1) // 排除表头
        const response = this.excelData.map((item) => ({
          nickname: String(item[1]),
          username: String(item[2]),
          password: String(item[3]),
          platform: String(item[4])
        }))
        if (response.length > 0) {
          ApiStudent.ImportExcel({ users: response }).then(res => {
            const { success, fail, failLog } = res.data
            if (fail > 0 && failLog.length > 0) {
              this.$alert(failLog.join('<br>'), '失败的账号', {
                confirmButtonText: '知道了',
                dangerouslyUseHTMLString: true,
                width: '550px'
              })
            } else {
              this.$message({
                type: 'success',
                message: `导入完成，成功：${success}，失败：${fail}`
              })
              this.getList()
            }
          })
        }
        this.excelData = []
        this.headers = []
      }
      reader.readAsArrayBuffer(file)
    },
    handleUserChange(val) {
      this.select = val
    },
    handleCloseSwitch() {
      this.switchVisible = false
    },
    switchPlatform(row) {
      this.activePlatformId = row.id
      if (this.select.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择需要切换平台的用户'
        })
      }
      var platformType = row.id
      var ids = this.select.filter(item => item.platform_type !== platformType).map(item => item.id)
      if (ids.length === 0) {
        return this.$message({
          type: 'warning',
          message: '没有需要切换的平台'
        })
      }
      this.$prompt('请输入授权密码', '授权验证', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputType: 'password'
      }).then(({ value }) => {
        ApiStudent.switchPlatform({
          ids: ids,
          platform_type: platformType,
          password: value
        }).then(res => {
          this.$message({
            type: 'success',
            message: '切换成功'
          })
          this.getList()
          this.activePlatformId = null
          this.switchVisible = false
        })
      }).catch(() => {
        this.$message.info('已取消输入')
        this.activePlatformId = null
      })
    },
    handleDelete() {
      if (this.select.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择需要删除的用户'
        })
      }
      this.$confirm('此操作将永久删除该, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiStudent.deleteUser({
          ids: this.select.map(item => item.id)
        }).then(response => {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.getList()
        })
      })
    },
    handleSchoolDialog() {
      if (this.select.length === 0) {
        return this.$message({
          type: 'warning',
          message: '请选择需要添加到团队的用户'
        })
      }
      var items = []
      var platformTypeSet = new Set(this.select.map(item => item.platform_type))
      var platformType = 0
      if (platformTypeSet.size > 1) {
        return this.$message({
          type: 'warning',
          message: '多个平台不支持一起学习'
        })
      } else {
        this.select.forEach(item => {
          items.push({
            user_id: item.id,
            platform: item.platform_type
          })
          platformType = item.platform_type
        })
        if (items.length > 5) {
          return this.$message({
            type: 'warning',
            message: '一次最多添加5条数据'
          })
        }
        this.dialogSchoolVisible = true
        console.log(platformType)
        // addSchoolTask({
        //   users: items,
        //   platform: platformType.toString()
        // }).then(res => {
        //   this.$message({
        //     type: 'success',
        //     message: '添加成功'
        //   })
        // })
      }
    },
    handleChangeMenu(val) {
      this.upload.parent_id = val
    },
    handleMenuSwitch(id) {
      this.upload.menus = this.upload.menus.map(item => {
        if (item.id === id) {
          item.selected = true
          this.upload.category_id = item.id
        } else {
          item.selected = false
        }
        return item // 确保 map 返回修改后的对象
      })
      ApiMedia.getFileList({
        page: 1,
        limit: 24,
        category_id: this.upload.category_id
      }).then(res => {
        this.upload.items = res.data.items.map(item => {
          item.selected = false // 修改拼写为 "selected"
          return item // 确保 map 返回修改后的对象
        })
        this.upload.total = res.data.total
        this.upload.pageSize = res.data.page_size
        this.upload.total_page = res.data.total_page
        this.upload.page = res.data.page
      })
    },
    handleMenuUpdate(id) {
      // this.upload.menus.map(item => {
      //   if (item.id === id) {
      //     console.log(item)
      //     this.upload.action = 'update'
      //     this.upload.dialogMenuVisible = true
      //     this.upload.category_marker = item.name
      //     this.upload.category_name = item.name
      //   }
      // })
    },
    handleMenuDelete(id) {
      ApiMedia.deleteMediaType({
        id: id
      }).then(response => {
        ApiMedia.getFileList({
          page: 1,
          limit: this.upload.pageSize,
          category_id: 1
        }).then(response => {
          this.upload.menus = this.upload.menus.filter(menu => menu.id !== id)
          this.upload.category_id = 1
          this.upload.items = response.data.items.map(item => {
            item.selected = false // 修改拼写为 "selected"
            return item // 确保 map 返回修改后的对象
          })
          this.upload.total = response.data.total
          this.upload.pageSize = response.data.page_size
          this.upload.total_page = response.data.total_page
          this.upload.page = response.data.page
          this.$message({
            type: 'success',
            message: '删除分类'
          })
        })
      })
    },
    handleMenuSave(data) {
      if (data.category_name === '') {
        this.$message({
          type: 'warning',
          message: '分类名称不能为空'
        })
        return false
      }
      ApiMedia.addMediaType(data).then(res => {
        ApiMedia.getMediaList({ page: 1, limit: this.upload.pageSize }).then(response => {
          var length = response.data.items.length
          this.upload.category_list = []
          this.upload.menus = []
          for (var i = 0; i < length; i++) {
            this.upload.category_list.push({
              label: response.data.items[i].name,
              value: response.data.items[i].id
            })
            this.upload.menus.push({
              id: response.data.items[i].id,
              name: response.data.items[i].name,
              selected: response.data.items[i].id === 1
            })
          }
        })
        this.upload.dialogMenuVisible = false
        this.upload.category_marker = ''
        this.upload.category_name = ''
        this.$message({
          type: 'success',
          message: '添加成功'
        })
      })
    },
    handleCutVideo(item) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      ApiMedia.cutVideo({
        id: item.id,
        user_id: this.selectUser.id
      }).then(res => {
        loading.close()
        this.$message({
          type: 'success',
          message: res.data
        })
      }).finally(() => {
        loading.close()
      })
    },
    handleSelectionImage(value) {
      this.upload.items = this.upload.items.map(item => {
        if (value.id === item.id) {
          item.selected = !item.selected // 直接取反，无需判断当前状态
        }
        return item // 返回修改后的 item
      })
    },
    handleSelectionImageSave() {
      var items = []
      this.upload.items.map(item => {
        if (item.selected === true) {
          items.push(item.id)
        }
      })
      if (items.length > 0) {
        this.upload.outerVisible = false
        console.log(items, this.selectUser.id)
      } else {
        this.$message({
          type: 'warning',
          message: '没有选择图片'
        })
      }
    },
    handleSelectAll() {
      this.upload.items = this.upload.items.map(item => {
        item.selected = !item.selected // 直接取反，无需判断当前状态
        return item // 返回修改后的 item
      })
    },
    handleSelectionImageDelete() {
      var items = []
      this.upload.items.map(item => {
        if (item.selected === true) {
          items.push(item.id)
        }
      })
      ApiMedia.deleteFile({ ids: items }).then(res => {
        ApiMedia.getFileList({
          page: 1,
          limit: this.upload.pageSize,
          category_id: this.upload.category_id
        }).then(res => {
          this.upload.items = res.data.items.map(item => {
            item.selected = false // 修改拼写为 "selected"
            return item // 确保 map 返回修改后的对象
          })
          this.upload.total = res.data.total
          this.upload.pageSize = res.data.page_size
          this.upload.total_page = res.data.total_page
          this.upload.page = res.data.page
        })
      })
    },
    handleSearchImages() {
      ApiMedia.getFileList({
        page: 1,
        limit: 24,
        category_id: this.upload.category_id,
        keywords: this.upload.keywords
      }).then(res => {
        this.upload.items = res.data.items.map(item => {
          item.selected = false // 修改拼写为 "selected"
          return item // 确保 map 返回修改后的对象
        })
        this.upload.total = res.data.total
        this.upload.pageSize = res.data.page_size
        this.upload.total_page = res.data.total_page
        this.upload.page = res.data.page
        this.upload.keywords = ''
      })
    },
    handleSearch() {
      this.loading = true
      ApiStudent.fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.pageSize = response.data.page_size
        this.total_page = response.data.total_page
        this.listQuery.page = response.data.page
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleRegionChange(value) {
      this.form.org_id = value
    },
    handlePlatformChange(value) {
      console.log(value)
      this.form.platform_type = value
    },
    cancel() {
      this.dialogVisible = false
      this.dialogEditVisible = false
    },
    handleTask(index, row) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      const start = Date.now()
      ApiStudent.getCourseList({
        user_id: row.id,
        username: row.username,
        password: row.password,
        platform_type: row.platform_type.toString(),
        device_id: row.device_id.toString(),
        code: '',
        sign: ''
      }).then(res => {
        const duration = Date.now() - start
        const delay = Math.max(300 - duration, 0) // 保证 loading 至少展示 300ms
        setTimeout(() => {
          this.course_list = res.data
          this.user = row
          // 使用 requestAnimationFrame 保证 DOM 渲染稳定后再显示弹窗
          requestAnimationFrame(() => {
            this.dialogCourseVisible = true
            loading.close()
          })
        }, delay)
      }).catch(() => {
        loading.close()
      })
    },
    handleLocalTask(index, row) {
      console.log('待开发')
    },
    handleSelectionChange(val) {
      if (val.learning_progress === 100) {
        return this.$message({
          type: 'warning',
          message: '已经学完不需要'
        })
      }
      var data = {
        user_id: this.selectUser.user_id,
        name: this.selectUser.name,
        username: this.selectUser.username,
        password: this.selectUser.password,
        region_id: this.selectUser.region_id,
        org_id: this.selectUser.org_id,
        course_id: val.id,
        course_name: val.title,
        order_id: val.order_id,
        token: this.selectUser.token
      }
      ApiTask.addTask(data).then(response => {
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.dialogCourseVisible = false
      })
    },
    ImportDialog() {
      console.log('待开发')
    },
    DownTemplate() {
      console.log('待开发')
    },
    submitForm() {
      if (this.form.org_id === 0) {
        return this.$message({
          type: 'warning',
          message: '机构不能为空'
        })
      }
      if (this.form.username === '') {
        return this.$message({
          type: 'warning',
          message: '登录账号不能为空'
        })
      }
      if (this.form.password === '') {
        return this.$message({
          type: 'warning',
          message: '账号密码不能为空'
        })
      }
      ApiStudent.createUser(this.form).then(response => {
        if (response.data.sign) {
          this.form.sign = response.data.sign
          this.form.image = response.data.image
          this.form.platform_type = response.data.platform_type
          this.form.code = ''
          this.$message({
            type: 'warning',
            message: '请填写验证码'
          })
        } else {
          this.form.sign = ''
          this.form.image = ''
          this.$message({
            type: 'success',
            message: '添加成功'
          })
          this.dialogVisible = false
          this.getList()
        }
      })
    },
    getCodeImage() {
      ApiStudent.VerifyImage().then(response => {
        this.form.sign = response.data.sign
        this.form.image = response.data.image
      })
    },
    editForm() {
      ApiStudent.updateUser(this.edit).then(response => {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.dialogEditVisible = false
        this.getList()
      })
    },
    handleClose() {
      this.dialogVisible = false
      this.dialogEditVisible = false
      this.dialogCourseVisible = false
    },
    handleMenuClose() {
      this.upload.dialogMenuVisible = false
      this.upload.innerVisible = false
    },
    handleUploadClose() {
      this.upload.innerVisible = false
    },
    handleUploadSizeChange(val) {
      this.upload.size = val
      ApiMedia.getFileList({ page: this.upload.page, limit: val }).then(response => {
        this.upload.items = response.data.items.map(item => {
          item.selected = false // 修改拼写为 "selected"
          return item // 确保 map 返回修改后的对象
        })
        this.upload.total = response.data.total
        this.upload.pageSize = response.data.page_size
        this.upload.total_page = response.data.total_page
        this.upload.page = response.data.page
      })
    },
    handleUploadPageChange(page) {
      this.upload.page = page
      ApiMedia.getFileList({ page: page, limit: this.upload.pageSize }).then(response => {
        this.upload.items = response.data.items.map(item => {
          item.selected = false // 修改拼写为 "selected"
          return item // 确保 map 返回修改后的对象
        })
        this.upload.total = response.data.total
        this.upload.pageSize = response.data.page_size
        this.upload.total_page = response.data.total_page
        this.upload.page = response.data.page
      })
    },
    handleUpload() {
      this.upload.outerVisible = false
    },
    handleFile(file) {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      const formData = new FormData()
      formData.append('file', file)
      formData.append('category_id', this.upload.category_id) // 添加额外参数
      ApiMedia.uploadFile(formData).then(response => {
        this.upload.innerVisible = false
        loading.close()
        this.upload.fileList = []
        ApiMedia.getFileList({ page: 1, limit: this.upload.pageSize, category_id: this.upload.category_id }).then(response => {
          this.upload.items = response.data.items.map(item => {
            item.selected = false // 修改拼写为 "selected"
            return item // 确保 map 返回修改后的对象
          })
          this.upload.total = response.data.total
          this.upload.pageSize = response.data.page_size
          this.upload.total_page = response.data.total_page
          this.upload.page = response.data.page
        })
      }).finally(() => {
        loading.close()
      })
    },
    handlePageChange(page) {
      this.loading = true
      this.listQuery.page = page
      ApiStudent.fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleSizeChange(val) {
      this.loading = true
      this.listQuery.page = 1
      this.listQuery.limit = val
      ApiStudent.fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listQuery.limit = response.data.page_size
        this.listQuery.page = response.data.page
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    },
    handleCommand(row, index, command) {
      switch (command) {
        case 'upload':
          this.selectUser = row
          this.upload.outerVisible = true
          this.upload.fileList = []
          this.upload.member = {
            name: row.nickname,
            card: row.username
          }
          this.getMediaType()
          break
        case 'edit':
          this.dialogEditVisible = true
          this.edit.id = row.id
          this.edit.nickname = row.nickname
          this.edit.username = row.username
          this.edit.password = row.password
          this.edit.org_id = row.org_id
          this.regions.map((item) => {
            if (item.value === row.org_id) {
              this.region = item.label
            }
          })
          break
        case 'update':
          console.log(this.command)
          break
        case 'delete':
          this.$confirm('此操作将永久删除该, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            ApiStudent.deleteUser({
              ids: [row.id]
            }).then(response => {
              this.$message({
                type: 'success',
                message: '删除成功'
              })
              this.getList()
            })
          })
          break
      }
    }
  }
}
</script>
<style lang="scss" scoped>
::v-deep {
  .el-upload {
    font-size: 12px !important;
  }
}
.upload-btn {
  border: none;
  width: 98px;
  height: 32px;
  line-height: 32px;
  display: inline-block;
  text-align: center;
  font-size: 14px;
  color: white;
  background: #1890ff;

  :active {
    color: white;
    background: #1890ff;
  }

  :hover {
    color: white;
    background: #1890ff;
  }

  :focus {
    color: white;
    background: #1890ff;
  }

  :visited {
    color: white;
    background: #1890ff;
  }
}

.input-with-select {
  width: 450px;
}
.platform-box{
  padding: 0;
  margin: 0;
  .platform-item{
    list-style: none;
    padding: 10px;
    margin: 5px 0;
    background: #f2f2f2;
    cursor: pointer;
    font-size: 16px;
    :hover {
      background: #409EFF!important;
      color: white;
    }
  }
  .active {
    background: #409EFF!important;
    color: #fff;
    border-radius: 4px;
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

.el-dialog__body {
  padding: 10px 15px;
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

      .course-item-photo {
        width: 130px;
        height: 80px;
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
}
@media only screen and (max-width: 470px) {
  ::v-deep {
    .el-dialog {
      width: 96% !important;
      .el-dialog__body{
        padding: 10px;
      }
    }
    .el-card__body{
      padding: 15px;
    }
    .el-button--small{
      padding: 8px 12px;
    }
    .el-button-group{
      margin-bottom: 5px;
      margin-top: 5px;
    }
  }
  .app-container{
    padding: 10px;
  }
  .input-with-select{
    width: 320px;
    margin-top: 5px;
  }
  .upload-btn{
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    font-size: 12px;
    width: auto!important;
  }
}
</style>
