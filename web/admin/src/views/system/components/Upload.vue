<template>
  <div class="app-container">
    <el-form label="存储管理" label-width="120px">
      <el-form-item label="存储类型">
        <el-select v-model="type" placeholder="存储类型" style="width: 590px;">
          <el-option label="本地存储" value="local" />
          <el-option label="对象存储" value="object" />
        </el-select>
      </el-form-item>
      <el-form-item v-if="type === 'object'" label="存储名称">
        <el-select v-model="objectName" placeholder="存储名称" style="width: 590px;" @change="handlerSelect">
          <el-option v-for="(item, index) in objectList" :key="index" :label="objectNames[index]" :value="item" />
        </el-select>
      </el-form-item>
      <div v-if="objectName==='oss' && type === 'object'">
        <el-form-item label="AccessKey">
          <el-input v-model.trim="objectConfig[objectName].accessKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="SecretKey">
          <el-input v-model.trim="objectConfig[objectName].secretKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model.trim="objectConfig[objectName].bucket" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Endpoint">
          <el-input v-model.trim="objectConfig[objectName].endpoint" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Domain">
          <el-input v-model.trim="objectConfig[objectName].domain" style="width: 590px;" />
        </el-form-item>
      </div>
      <div v-if="objectName ==='qiniu' && type === 'object'">
        <el-form-item label="AccessKey">
          <el-input v-model.trim="objectConfig[objectName].accessKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="SecretKey">
          <el-input v-model.trim="objectConfig[objectName].secretKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model.trim="objectConfig[objectName].bucket" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Domain">
          <el-input v-model.trim="objectConfig[objectName].domain" style="width: 590px;" />
        </el-form-item>
      </div>
      <div v-if="objectName ==='cos' && type === 'object'">
        <el-form-item label="AccessKey">
          <el-input v-model.trim="objectConfig[objectName].accessKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="SecretKey">
          <el-input v-model.trim="objectConfig[objectName].secretKey" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model.trim="objectConfig[objectName].bucket" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Region">
          <el-input v-model.trim="objectConfig[objectName].region" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Domain">
          <el-input v-model.trim="objectConfig[objectName].domain" style="width: 590px;" />
        </el-form-item>
      </div>
      <div v-if="objectName ==='upyun' && type === 'object'">
        <el-form-item label="操作员">
          <el-input v-model.trim="objectConfig[objectName].operator" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model.trim="objectConfig[objectName].password" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model.trim="objectConfig[objectName].bucket" style="width: 590px;" />
        </el-form-item>
        <el-form-item label="Domain">
          <el-input v-model.trim="objectConfig[objectName].domain" style="width: 590px;" />
        </el-form-item>
      </div>
      <el-form-item label="文件类型">
        <el-input v-model.trim="ext" style="width: 590px;" />
      </el-form-item>
      <el-form-item label="文件大小">
        <el-input v-model.trim="size" style="width: 590px;" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handlerSubmit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getConfigs, updateConfigs } from '@/api/config'

export default {
  data() {
    return {
      type: 'local',
      objectName: 'oss',
      objectList: ['oss', 'qiniu', 'cos', 'upyun'],
      objectNames: ['阿里云 OSS', '七牛云 未实现', '腾讯云 COS 未实现', '又拍云 未实现'],
      type_list: ['local', 'object'],
      objectConfig: {
        oss: {
          accessKey: '',
          secretKey: '',
          bucket: '',
          endpoint: '',
          domain: ''
        },
        qiniu: {
          accessKey: '',
          secretKey: '',
          bucket: '',
          domain: ''
        },
        cos: {
          accessKey: '',
          secretKey: '',
          bucket: '',
          region: '',
          domain: ''
        },
        upyun: {
          operator: '',
          password: '',
          bucket: '',
          domain: ''
        }
      },
      ext: 'png,jpg,jpeg,png,gif,webp',
      size: 51200
    }
  },
  mounted() {
    getConfigs({ key: 'upload' }).then((res) => {
      this.type = res.data.type
      this.objectName = res.data.name
      this.objectConfig[this.objectName] = res.data.config
      this.ext = res.data.ext
      this.size = res.data.size
    })
  },
  methods: {
    handlerSelect(value) {
      this.objectName = value
    },
    handlerSubmit() {
      updateConfigs({
        key: 'upload',
        config: {
          type: this.type,
          name: this.objectName,
          config: this.objectConfig[this.objectName],
          ext: this.ext,
          size: this.size
        }
      }).then(res => {
        this.$message({
          message: '更新上传',
          type: 'success'
        })
      })
    }
  }
}
</script>
