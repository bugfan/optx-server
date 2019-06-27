// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import ElementUI from 'element-ui'
import store from './store'

import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)

import VueResource from "vue-resource";
//import mock from "./mock/mock";
// 引入echarts
import echarts from "echarts";
Vue.prototype.$echarts = echarts;

Vue.config.productionTip = false;
Vue.use(VueResource);
/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  store,
  components: { App },
  template: "<App/>"
});
