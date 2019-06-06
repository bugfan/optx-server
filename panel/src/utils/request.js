import axios from 'axios';
import router from "../router";
import { MessageBox } from 'element-ui'
import { Message } from 'element-ui'



// 创建axios实例
const service = axios.create({
  baseURL: '/api',  // process.env.BASE_API, // api的base_url
  timeout: 15000 // 请求超时时间
});

// respone拦截器
service.interceptors.response.use(
  response => {
    console.log(1)
    return response
  },
  error => {
    // console.log(2,error)
    const res = error.response
    const req = error.request
    const url = new URL(req.responseURL)

    if (res.status == 401){
      // jump to login page
      MessageBox.confirm('你已被登出，可以取消继续留在该页面，或者重新登录', '确定登出', {
        confirmButtonText: '重新登录',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        router.push({path: '/login'});           
      })
    }   
    if (res.status == 400){
      Message({
            type: 'warning',
            message: '参数错误!'
        });
    }
    return Promise.reject(error)
  }
)

export default service;

export function curd(path) {
  return {
    all (params) {
      return service({
        url: path,
        method: 'get',
        params
      });
    },
    get (id, params) {
      return service({
        url: path + '/' + id,
        method: 'get',
        params
      });
    },
    update (id, data, params) {
      return service({
        url: path + '/' + id,
        method: 'put',
        data,
        params
      });
    },
    create (data, params) {
      return service({
        url: path,
        method: 'post',
        data,
        params
      });
    },
    del (id, params) {
      return service({
        url: path + '/' + id,
        method: 'delete',
        params
      });
    }
  };
}
