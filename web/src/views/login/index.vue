<template>
  <el-card
      v-show="mode === 'login'"
      shadow="hover"
      class="form-card"
  >
    <h3>登录系统</h3>
    <el-form>
      <el-form-item>
        <el-input v-model="loginForm.account"/>
      </el-form-item>
      <el-form-item>
        <el-input v-model="loginForm.password"/>
      </el-form-item>
      <el-form-item>
        <el-button class="confirm-btn" type="primary" @click="onLogin">登录</el-button>
      </el-form-item>
      <el-form-item>
        <el-button class="confirm-btn" type="primary" @click="mode = 'signup'">去注册</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="text">忘记密码</el-button>
      </el-form-item>
    </el-form>
  </el-card>

  <el-card
      v-show="mode === 'signup'"
      shadow="hover"
      class="form-card"
  >
    <h3>注册账号</h3>
    <el-form>
      <el-form-item>
        <el-input v-model="loginForm.account"/>
      </el-form-item>
      <el-form-item>
        <el-input v-model="loginForm.password"/>
      </el-form-item>
      <el-form-item>
        <el-button class="confirm-btn" type="primary" @click="onLogin">注册</el-button>
      </el-form-item>
      <el-form-item>
        <el-button class="confirm-btn" type="primary" @click="mode = 'login'">去登录</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
export default {
  name: "login"
}
</script>

<script setup>
import * as baseApi from '@/api/system/base'
import {reactive, ref} from 'vue'

const mode = ref('login')
const loginForm = reactive({
  account: '',
  password: '',
})

const onLogin = () => {
  console.log(baseApi)
  baseApi.login(loginForm).then(res =>{
    console.log(res)
  })
}

const onSignup = () => {
  baseApi.signup(loginForm).then(res =>{
    console.log(res)
  })
}

</script>


<style lang="scss" scoped>
.form-card {
  margin: 50px auto;
  width: 350px;
  height: 400px;
}

.confirm-btn {
  width: 100%;
}
</style>