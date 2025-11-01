<template>
  <div class="app-container">
    <el-dialog
      title="新增兑换码"
      :visible.sync="visible"
      width="700px"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="账户昵称">
          <el-input v-model="redeem.remark" autocomplete="off" style="width: 560px;" />
          <span style="color: red; margin-left: 5px;">*</span>
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生成数量">
              <el-input v-model="redeem.quantity" autocomplete="off" style="width: 220px;" />
              <span style="color: red; margin-left: 5px;">*</span>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="到期时间(天)">
              <el-input v-model="redeem.valid_days" type="text" placeholder="0代表不过期" autocomplete="off" style="width: 220px;" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="所属代理">
          <el-select v-model="redeem.school_id" style="width: 560px;" @change="handleSchoolChange">
            <el-option
              v-for="option in options"
              :key="option.id"
              :label="option.name"
              :value="option.id"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleSubmit">保存</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import * as ApiSchool from '@/api/school'
import * as ApiRedeem from '@/api/school_redeem'
export default {
  name: 'RedeemCreate',
  props: {
    visible: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      redeem: {
        remark: '',
        valid_days: 0,
        quantity: 1,
        school_id: 0
      },
      options: []
    }
  },
  mounted() {
    ApiSchool.SchoolList({ page: 1, limit: 100 }).then(response => {
      this.options = response.data.items.map((item) => {
        return {
          id: item.id,
          name: item.nickname
        }
      })
      this.redeem.school_id = this.options.length > 0 ? this.options[0].id : 0
    })
  },
  methods: {
    handleSchoolChange(value) {
      this.redeem.school_id = value
    },
    handleClose() {
      this.$emit('update:visible', false)
    },
    handleSubmit() {
      this.loading = true
      if (!this.redeem.remark) {
        this.$message.error('请填写账户昵称')
        this.loading = false
        return
      }
      if (!this.redeem.quantity || isNaN(this.redeem.quantity) || this.redeem.quantity <= 0) {
        this.$message.error('生成数量必须为大于0的数字')
        this.loading = false
        return
      }
      if (!this.redeem.school_id || isNaN(this.redeem.school_id) || this.redeem.school_id <= 0) {
        this.$message.error('请选择所属代理')
        this.loading = false
        return
      }
      this.redeem.school_id = parseInt(this.redeem.school_id)
      this.redeem.quantity = parseInt(this.redeem.quantity)
      this.redeem.valid_days = parseInt(this.redeem.valid_days) || 0
      ApiRedeem.add(this.redeem).then(response => {
        this.$message.success('兑换码生成成功')
        this.$emit('update:visible', false)
        this.$emit('refresh')
        this.loading = false
      }).finally(() => {
        this.loading = false
      })
    }
  }
}
</script>

<style scoped lang="scss">

</style>
