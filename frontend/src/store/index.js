import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    darkmode: "darkmode",
    lang: "locale",
    snackbar: false,
    snackbarMessage: "",
    newTask: {},
    overlay: false
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  }
})
