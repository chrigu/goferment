import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    steps: []
  },
  getters: {
    steps: (state) => state.steps
  },
  mutations: {
    addStep (state, step) {
      state.steps = [...state.steps, step]
    }
  },
  actions: {
    addStep ({ commit }, step) {
      console.log('some foo', step)
      commit('addStep', step)
    }
  },
  modules: {
  }
})
