import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        takingsByStylist: {},
        takingsByMonth: {},
        costsByCat: {},
        costsByMonth: {}
    },

    getters: {

    },

    mutations: {
        LOAD_TAKINGS_BY_STYLIST(state, payload) {
            state.takingsByStylist = payload
        },

        LOAD_TAKINGS_BY_MONTH(state, payload) {
            state.takingsByMonth = payload
        },

        LOAD_COSTS_BY_CAT(state, payload) {
            state.costsByCat = payload
        },

        LOAD_COSTS_BY_MONTH(state, payload) {
            state.costsByMonth = payload
        }
    },

    actions: {
        loadTakingsByStylist({ commit }) {
            axios
                .get('/api/takings-by-stylist/' + 'all/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_STYLIST', data)
                })
        },

        loadTakingsByMonth({ commit }) {
            axios
                .get('/api/takings-by-month/' + 'all/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_MONTH', data)
                })
        },

        loadCostsByCat({ commit }) {
            axios
                .get('/api/costs-by-cat/' + '06517160/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_CAT', data)
                    console.log(data)
                })
        },

        loadCostsByMonth({ commit }) {
            axios
                .get('/api/costs-by-month/' + 'all/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_MONTH', data)
                })
        }
    }
})