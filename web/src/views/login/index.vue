<template>
  <el-card class="login-card" :body-style="{padding: '38px 30px'}">
    <el-form ref="form" :model="form" size="large">
      <el-form-item prop="account">
        <el-input v-model="form.account"/>
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="form.password"
                  type="password"
                  show-password
        />
      </el-form-item>
    </el-form>
    <el-button type="primary" round class="login-button" @click="confirmLogin">登录</el-button>
  </el-card>
</template>

<script lang="ts">
import {login} from '@/api/system/base'
import {defineComponent} from 'vue';

export default defineComponent({
  name: 'Login',
  props: {
    msg: String,
  },
  data() {
    return {
      form: {
        account: 'admin',
        password: '123456'
      }
    }
  },
  methods: {
    confirmLogin() {
      login(this.form).then((res: any) => {
        console.log(res)
      })
    }
  }
});
</script>

<style scoped lang="scss">
.login-card {
  margin: 50px auto;
  width: 380px;

  .login-button {
    width: 100%;
  }
}
</style>