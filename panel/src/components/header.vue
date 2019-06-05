<template>
    <div class="app-head">
      <div class="app-head-inner">
        <!--<router-link :to="{name: 'index'}" class="head-logo">-->
          <img class="head-logo" src="@/assets/img/logo.png">

        <!--</router-link>-->
        <div class="head-nav">
          <ul class="nav-list">
            <!--<li @click="showDialog('isShowLogin')">登录</li>
            <li class="nav-pile">|</li>
            <li @click="showDialog('isShowReg')">注册</li>
            <li class="nav-pile">|</li>
            <li @click="showDialog('isShowAbout')">关于</li>-->
          </ul>
        </div>  
      </div>
    </div>
</template>

<script type="text/javascript">
/*import star from 'components/star/star';

  export default {
    props: {
      seller: {
        type: Object
      }
    },
    data() {
      return {
        detailShow: false
      };
    },
    methods: {
      showDetail() {
        this.detailShow = true;
      },
      hideDetail() {
        this.detailShow = false;
      }
    },
    created() {
      this.classMap = ['decrease', 'discount', 'special', 'invoice', 'guarantee'];
    },
    components: {
      star
    }
  };*/
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