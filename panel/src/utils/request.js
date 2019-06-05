import axios from 'axios';

// 创建axios实例
const service = axios.create({
  baseURL: '/api',  // process.env.BASE_API, // api的base_url
  timeout: 15000 // 请求超时时间
});

// // respone拦截器
// service.interceptors.response.use(
//   response => {
//     console.log(1)
//     return response
//   },
//   error => {
//     console.log(2,error)
//     return Promise.reject(error)
//   }
// )

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
