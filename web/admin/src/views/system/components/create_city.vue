<template>
  <div class="app-container">
    <el-dialog
      title="增加地区"
      :visible.sync="visible"
      max-width="550px"
      :before-close="handleClose"
    >
      <el-form ref="createForm" :model="form">
        <el-form-item label="地区名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="form.name" placeholder="深圳" autocomplete="off" />
        </el-form-item>
        <el-form-item label="地区英文" :label-width="formLabelWidth" prop="city_name">
          <el-input v-model="form.city_name" placeholder="shenzhen" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户端ID" :label-width="formLabelWidth" prop="client_id">
          <el-input v-model="form.client_id" placeholder="client_id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="地区描述" :label-width="formLabelWidth" prop="description">
          <el-input v-model="form.description" type="textarea" rows="4" placeholder="描述" autocomplete="off" />
        </el-form-item>
        <el-form-item label="平台排序" :label-width="formLabelWidth" prop="sort_num">
          <el-input v-model="form.sort_num" placeholder="99" autocomplete="off" />
        </el-form-item>
        <el-form-item label="地区状态" :label-width="formLabelWidth" prop="code">
          <el-radio-group v-model="form.status" @change="handleStatusChange">
            <el-radio label="正常">正常</el-radio>
            <el-radio label="维护中">维护中</el-radio>
            <el-radio label="下架">下架</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as ApiArea from '@/api/area'

export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      formLabelWidth: '100px',
      form: {
        name: '',
        client_id: '',
        city_name: '',
        description: '',
        sort_num: 99,
        status: '正常'
      }
    }
  },
  methods: {
    cancel() { this.$emit('update:visible', false) },
    submit() {
      ApiArea.Create(this.form).then(response => {
        this.$message({
          message: '操作成功',
          type: 'success'
        })
      }).finally(() => {
        this.$emit('update:visible', false)
      })
    },
    handleStatusChange(status) {
      this.form.status = status
    },
    handleClose() {
      this.$emit('update:visible', false)
    }
  }
}
</script>

<style scoped lang="scss">
::v-deep {
  .el-dialog {
    width: 600px;
  }
}
</style>
