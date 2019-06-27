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
            <i>{{username}}</i><br/>
            </li>
            <li class="nav-pile">&nbsp;| &nbsp;</li>
             <li>
            <el-button @click="logout()">{{action}}</el-button>
            </li>
             <!-- <li>
            <el-button @click="look()">看状态</el-button>
            </li> -->
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
          username:this.getUsername,
          action:'登录'
      };
    },
    mounted() {
      console.log("username:",this.$store.state.username)
      this.username=this.getUName()
    },
    computed: {
      changeUsername() {
        return this.$store.state.username;
      }
    },
    watch: {
      changeUsername(val) {
        this.username=this.getUName(val)
      },
      username(val){
        if (val=="未登录"){
          this.action="登录"
        }else{
          this.action="退出"
        }
      }
    },
    methods: {
      look(){
        console.log("look:",this.$store.state.username)
      },
      getUName() {
        return this.$store.state.username=='' ? '未登录' : this.$store.state.username
      },
      logout() {
        api.logout().then(res => {
          this.$store.commit("setUsername","")
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