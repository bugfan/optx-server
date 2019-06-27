<template>
    <div class="login-container">
    <el-form class="login-form" autoComplete="on" :model="loginForm" :rules="loginRules" ref="loginForm" label-position="left">
        <h3 class="title">后台入口</h3>
        <el-form-item prop="username">
        <span class="svg-container svg-container_login">
            <svg-icon icon-class="user" />
        </span>
        <el-input name="username" type="text" v-model="loginForm.username" autoComplete="on" placeholder="用户名" />
        </el-form-item>
        <el-form-item prop="password">
        <span class="svg-container">
            <svg-icon icon-class="password"></svg-icon>
        </span>
        <el-input name="password" :type="pwdType" @keyup.enter.native="handleLogin" v-model="loginForm.password" autoComplete="on"
                    placeholder="口令"></el-input>
        <span class="show-pwd" @click="showPwd"><svg-icon icon-class="eye" /></span>
        </el-form-item>
        <el-form-item>
        <el-button type="primary" style="width:100%;" :loading="loading" @click.native.prevent="handleLogin">
            登录
        </el-button>
        </el-form-item>
    </el-form>
    </div>
</template>

<script>
import api from '@/apis/login'

export default {
  name: 'login',
  data() {
    const validateUsername = (rule, value, callback) => {
        if (!this.isvalidUsername(value)) {
            callback(new Error('请输入正确的用户名'))
        } else {
            callback()
        }
    }
    const validatePass = (rule, value, callback) => {
        if (value.length < 5) {
            callback(new Error('密码不能小于5位'))
        } else {
            callback()
        }
    }
    return {
        loginForm: {
          username: '',
          password: ''
        },
        loginRules: {
            username: [{ required: true, trigger: 'blur', validator: validateUsername }],
            password: [{ required: true, trigger: 'blur', validator: validatePass }]
        },
        loading: false,
        pwdType: 'password'
    }
  },
  created: function() {

  },
  methods: {
    isvalidUsername(str) {
        return true
    },
    reset() {
        this.loginForm={}
    },
    showPwd() {
        if (this.pwdType === 'password') {
            this.pwdType = ''
        } else {
            this.pwdType = 'password'
        }
    },
    handleLogin() {
        console.log(this.loginForm.username,this.loginForm.password)
        return  api.login({
            Name:this.loginForm.username,
            Password:this.loginForm.password,
        }).then(res => {
            if (res.status == 200){
                this.$store.commit("setUsername",'rrr')
                this.$router.push({path: '/options'})
                this.$router.push(0)
                this.reset()
            }
        })
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
  .title {
    font-size: 26px;
    font-weight: 400;
    color: $light_gray;
    margin: 0px auto 40px auto;
    text-align: center;
    font-weight: bold;
  }
.login-container {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;
    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      &:-webkit-autofill {
        -webkit-box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: #fff !important;
      }
    }
  }
  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>