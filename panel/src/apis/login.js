import request from '@/utils/request'

export default{
      login(data, params) {
        return request({
          url: '/admin/login',
          method: 'post',
          data,
          params
        })
      },
      logout(data, params) {
        return request({
          url: '/admin/logout',
          method: 'post',
          data,
          params
        })
      }
}
