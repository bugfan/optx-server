import Vue from "vue";
import Router from "vue-router";
import Options from "@/components/options";
import Login from "@/components/login";

Vue.use(Router);

export default new Router({
  routes: [
    { name: "default", path: "", redirect: { name: "options" } }, // 默认
    { name: "options", path: "/options", component: Options },
    { name: "login", path: "/login", component: Login }

  ]
});
