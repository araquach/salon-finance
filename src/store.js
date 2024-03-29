import Vue from 'vue'
import Vuex from 'vuex'
import axios from "axios"
import {format, parseISO} from "date-fns"

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        salon: null,
        stylists: [],
        stylist: {},
        dateRange: {
            startDate: '2018-07-01', endDate: '2022-07-28',
        },
        loaded: false,
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
                datasets: [{
                    borderColor: '#3e95cd', data: [], fill: false, label: state.stylist.first_name,
                }], labels: []
            }

            Array.from(stmbm).forEach(el => res.labels.push(format(parseISO(el.month), 'LLL yy')))
            Array.from(stmbm).forEach(el => res.datasets[0].data.push(el.total))

            return res
        },

        getJakataStylists(state) {
            return state.stylists.filter(s => s.salon === 1)
        },

        getPKStylists(state) {
            return state.stylists.filter(s => s.salon === 2)
        },

        getBaseStylists(state) {
            return state.stylists.filter(s => s.salon === 3)
        },

        getChosenStylist(state) {
            return state.stylist
        },

        stylistLink(state) {
            return state.stylist.first_name + ' ' + state.stylist.last_name
        }
    },

    mutations: {
        LOAD_STYLISTS(state, payload) {
            state.stylists = payload
        },

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

        UPDATE_STYLIST(state, payload) {
            state.stylist = payload
            state.loaded = true
        },

        UPDATE_DATE_RANGE(state, payload) {
            state.dateRange = payload
        },

        UPDATE_DATE_RANGE_TYPE(state, payload) {
            if (payload === 'quarter') {
                this.state.dateRange = {
                    startDate: '2022-05-01', endDate: '2022-07-31'
                }
            }
            if (payload === 'year') {
                this.state.dateRange = {
                    startDate: '2021-07-31', endDate: '2022-07-31'
                }
            }
            if (payload === 'all') {
                this.state.dateRange = {
                    startDate: '2016-07-01', endDate: '2022-07-31'
                }
            }
        },

        SET_TOTAL_TURNOVER(state, payload) {
            state.totalTurnover = payload
        },

        ADJUST_TOTAL_TURNOVER(state, payload) {
            state.totalTurnover = state.totalTurnover - payload
        }
    },

    actions: {
        loadStylists({commit}) {
            axios.get(`/api/team-members`).then((response) => {
                commit('LOAD_STYLISTS', response.data)
            })
        },

        loadStylistTakingsMonthByMonth({commit}) {
            axios.get(`/api/stylist-takings-month-by-month/${store.state.dateRange.startDate}/${store.state.dateRange.endDate}/${store.getters.stylistLink}`)
                .then((response) => {
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