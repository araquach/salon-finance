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
        LOAD_TOTAL_TAKINGS(state, payload) {
            state.totalTakings = payload
        },
    },

    actions: {
        loadTotalTakings({ commit }) {
            axios.get('/api/total-takings').then((response) => {
                commit('LOAD_TOTAL_TAKINGS', response.data
                )
            })
        }
    }
})