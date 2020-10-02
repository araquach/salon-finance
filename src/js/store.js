import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        takings: {},
    },

    getters: {

    },

    mutations: {
        LOAD_TAKINGS(state, payload) {
            state.takings = payload
        },
    },

    actions: {
        loadTakings({ commit }) {
            axios.get('/api/takings').then((response) => {
                commit('LOAD_TAKINGS', response.data
                )
            })
        }
    }
})