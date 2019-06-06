import Vue from "vue";
import Router from "vue-router";
import Options from "@/components/options";
import Login from "@/components/login";
import Insert from "@/components/insert";

Vue.use(Router);

export default new Router({
  routes: [
    { name: "default", path: "", redirect: { name: "options" } }, // 默认
    { name: "options", path: "/options", component: Options },
    { name: "insert", path: "/insert", component: Insert },
    { name: "login", path: "/login", component: Login }

  ]
});
