<template>
  <div class="app-container">
    <el-form label-width="80px">
      <el-form-item label="登录账号">
        <el-input v-model="username" />
      </el-form-item>
      <el-form-item label="登录密码">
        <el-input v-model="password" />
      </el-form-item>
      <el-form-item label="用户编码">
        <el-input v-model="nickname" />
      </el-form-item>
      <el-form-item label="平台">
        <el-select v-model="platform_type">
          <el-option
            v-for="platform in platform_list"
            :key="platform.id"
            :label="platform.name"
            :value="platform.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="用户机构">
        <el-select v-model="org_id">
          <el-option
            v-for="org in org_list"
            :key="org.id"
            :label="org.name"
            :value="org.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="是否选择机构">
        <el-checkbox v-model="is_select_org" label="1" />
      </el-form-item>
      <el-form-item label="不选机构" :label-width="formLabelWidth">
        <el-switch v-model="is_select_org" />
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button type="primary">添加账号</el-button>
    </span>
  </div>
</template>
<script>
export default {
  name: 'CreateUser',
  data() {
    return {
      user: {
        nickname: '',
        username: '',
        password: 'Dp111111',
        org_id: 0,
        code: '',
        sign: '',
        image: '',
        is_select_org: true
      }
    }
  },
  methods: {
    createUser() {
      this.$store.dispatch('users/createUser', this.user)
    }
  }
}
</script>
<style scoped lang="scss">

</style>
