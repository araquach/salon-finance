import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        takingsByStylist: {},
        takingsByDateRange: {},
        totalsByDateRange: {},
        costsByCat: {},
        costsByDateRange: {}
    },

    getters: {

    },

    mutations: {
        LOAD_TAKINGS_BY_STYLIST(state, payload) {
            state.takingsByStylist = payload
        },

        LOAD_TAKINGS_BY_DATE_RANGE(state, payload) {
            state.takingsByDateRange = payload
        },

        LOAD_TOTALS_BY_DATE_RANGE(state, payload) {
            state.totalsByDateRange = payload
        },

        LOAD_COSTS_BY_CAT(state, payload) {
            state.costsByCat = payload
        },

        LOAD_COSTS_BY_DATE_RANGE(state, payload) {
            state.costsByDateRange = payload
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

        loadTakingsByDateRange({ commit }) {
            axios
                .get('/api/takings-by-date-range/' + 'all/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_DATE_RANGE', data)
                })
        },

        loadTotalsByDateRange({commit}) {
            axios
                .get('/api/totals-by-date-range/' + '2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TOTALS_BY_DATE_RANGE', data)
                })
        },

        loadCostsByCat({ commit }) {
            axios
                .get('/api/costs-by-cat/' + 'all/2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_CAT', data)
                })

        },

        loadCostsByDateRange({ commit }) {
            axios
                .get('/api/costs-by-date-range/' + '2021-01-01/2021-12-31')
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_DATE_RANGE', data)
                })
        }
    }
})