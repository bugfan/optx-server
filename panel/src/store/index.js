import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
 
const store = new Vuex.Store({
  state:{
      username: sessionStorage.getItem("opt_username"),
  },
  getters:{
      getUsername:state => state.username,
  },
  mutations:{
      setUsername(state,data){
          state.username=data
          sessionStorage.setItem("opt_username",state.username)
      }
  },
  actions:{
  },
  Modules:{
  }
})

export default store