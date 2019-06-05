import Vue from "vue";
import Router from "vue-router";
import Create from "@/components/create";
import List from "@/components/list";
// import Charts from "@/components/charts";

Vue.use(Router);

export default new Router({
  routes: [
    { name: "default", path: "", redirect: { name: "create" } }, // 默认
    { name: "create", path: "/create", component: Create }, //
    { name: "list", path: "/list", component: List }

  ]
});
