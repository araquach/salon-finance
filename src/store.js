import Vue from 'vue'
import Vuex from 'vuex'
import axios from "axios"
import {format, parseISO} from "date-fns"

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        salon: 'base',
        stylist: 'beth',
        dateRange: {
            startDate: '2021-12-01',
            endDate: '2022-02-28',
        },
        totalTurnover: 895304.00,
        totalCosts: null,
        takingsByStylist: {},
        takingsByDateRange: {},
        stylistTakingsMonthByMonth: {},
        totalsByDateRange: {},
        costsByCat: {},
        costsByDateRange: {}
    },

    getters: {
        getStylistTakingsMonthByMonth(state) {
            let stmbm = state.stylistTakingsMonthByMonth

            let res = {
                datasets: [
                    {
                        borderColor: '#3e95cd',
                        data: [],
                        fill: false,
                        label: (state.stylist.charAt(0).toUpperCase() + state.stylist.slice(1)),
                    }
                ],
                labels: []
            }

            Array.from(stmbm).forEach(el => res.labels.push(format(parseISO(el.month), 'LLL yy')))
            Array.from(stmbm).forEach(el => res.datasets[0].data.push(el.total))

            return res
        }
    },

    mutations: {
        LOAD_TAKINGS_BY_STYLIST(state, payload) {
            state.takingsByStylist = payload
        },

        LOAD_TAKINGS_BY_DATE_RANGE(state, payload) {
            state.takingsByDateRange = payload
        },

        LOAD_STYLIST_TAKINGS_MONTH_BY_MONTH(state, payload) {
            state.stylistTakingsMonthByMonth = payload
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
        loadStylistTakingsMonthByMonth({commit}) {
            axios.get(`/api/stylist-takings-month-by-month/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}/${store.state.stylist}`).then((response) => {
                commit('LOAD_STYLIST_TAKINGS_MONTH_BY_MONTH', response.data)
            })
        },

        loadTakingsByStylist({commit}) {
            axios
                .get(`/api/takings-by-stylist/${store.state.salon}/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_TAKINGS_BY_STYLIST', data)
                })
        },

        loadTakingsByDateRange({commit}) {
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

        loadCostsByCat({commit}) {
            axios
                .get(`/api/costs-by-cat/${store.state.salon}/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_CAT', data)
                })

        },

        loadCostsByDateRange({commit}) {
            axios
                .get(`/api/costs-by-date-range/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_DATE_RANGE', data)
                })
        },

        loadCostAndTakingsTotals({commit}) {
            axios
                .get(`/api/costs-and-takings-totals/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}`)
                .then(r => r.data)
                .then(data => {
                    commit('LOAD_COSTS_BY_DATE_RANGE', data)
                })
        }
    }
})