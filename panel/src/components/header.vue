<template>
    <div class="app-head">
      <div class="app-head-inner">
        <!--<router-link :to="{name: 'index'}" class="head-logo">-->
          <img class="head-logo" src="@/assets/img/4.png">
        <!--</router-link>-->
        <div class="head-nav">
          <ul class="nav-list">
            <!--<li @click="showDialog('isShowLogin')">登录</li>
            <li class="nav-pile">|</li>
            <li @click="showDialog('isShowReg')">注册</li>
            <li class="nav-pile">|</li>
            <li @click="showDialog('isShowAbout')">关于</li> -->
            <li>
            <el-button @click="logout()">退出</el-button>
            </li>
          </ul>
        </div>  
      </div>
    </div>
</template>

<script type="text/javascript">
import api from '@/apis/login'

  export default {
    data() {
      return {
        
      };
    },
    methods: {
      logout() {
        api.logout().then(res => {
          this.$router.push(0)
          this.$router.push({path: '/login'});  
        })
      },
    },
    created() {
    }
  };
</script>

<style>
.app-head {
  background: #363636;
  color: #b2b2b2;
  height: 60px;
  line-height: 60px;
  width: 100%;
}
.app-head-inner {
  width: 1200px;
  margin: 0 auto;
}
.head-logo {
  float: left;
  margin-top: 8px;
  width: 50px;
}
.head-nav {
  float: right;
}
.head-nav ul {
  overflow: hidden;
}
.head-nav li {
  cursor: pointer;
  float: left;
}
</style>