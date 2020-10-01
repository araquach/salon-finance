import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        totalTakings: {},

    },

    getters: {

    },

    mutations: {

    },

    actions: {
        loadTotalTakings({ commit }) {
            axios.get('/api/takings').then((response) => {
                commit('UPDATE_TOTAL_TAKINGS', response.data)
            })
        }
    }
})