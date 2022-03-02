import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        salon: "all",
        dateRange: {
            startDate: '2021-07-01',
            endDate: '2022-02-28',
        },
        totalTurnover: 895304.00,
        totalCosts: null,
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
        },

        UPDATE_SALON(state, payload) {
            state.salon = payload
        },

        UPDATE_DATE_RANGE(state, payload) {
            state.dateRange = payload
        },

        SET_TOTAL_TURNOVER(state, payload) {
            state.totalTurnover = payload
        },

        ADJUST_TOTAL_TURNOVER(state, payload) {
            state.totalTurnover = state.totalTurnover - payload
        }
    },

    actions: {
        loadTakingsByStylist({ commit }) {
            axios
                .get(`/api/takings-by-stylist/${store.state.salon}/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_STYLIST', data)
                })
        },

        loadTakingsByDateRange({ commit }) {
            axios
                .get(`/api/takings-by-date-range/${store.state.salon}/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_DATE_RANGE', data)
                })
        },

        loadTotalsByDateRange({commit}) {
            axios
                .get(`/api/totals-by-date-range/${store.state.dateRasnge.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TOTALS_BY_DATE_RANGE', data)
                })
        },

        loadCostsByCat({ commit }) {
            axios
                .get(`/api/costs-by-cat/${store.state.salon}/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_CAT', data)
                })

        },

        loadCostsByDateRange({ commit }) {
            axios
                .get(`/api/costs-by-date-range/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_DATE_RANGE', data)
                })
        },

        loadCostAndTakingsTotals({ commit}) {
            axios
                .get(`/api/costs-andtakings-totals/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_DATE_RANGE', data)
                })
        }
    }
})