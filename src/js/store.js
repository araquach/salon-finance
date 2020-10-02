import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        takings: {},
        costs: []
    },

    getters: {

    },

    mutations: {
        LOAD_TAKINGS(state, payload) {
            state.takings = payload
        },

        LOAD_COSTS(state, payload) {
            state.costs = payload
        }
    },

    actions: {
        loadTakings({ commit }) {
            axios.get('/api/takings').then((response) => {
                commit('LOAD_TAKINGS', response.data
                )
            })
        },

        loadCosts({ commit }) {
            axios.get('/api/costs').then((response) => {
                commit('LOAD_COSTS', response.data
                )
            })
        }
    }
})