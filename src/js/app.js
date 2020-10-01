import Vue from "vue"
import VueRouter from "vue-router"
import App from "./App.vue"
import { store } from "./store"
import { routes } from "./routes"

import Buefy from "buefy"
Vue.use(Buefy)
Vue.use(require('vue-moment'))
Vue.use(VueRouter)

const router = new VueRouter({
    mode: 'history',
    routes
})

window.axios = require('axios')

Vue.filter("toCurrency",  amount => Number(amount).toLocaleString('en-GB',
    {
    style: 'currency',
    currency: 'GBP'
}))

new Vue({
    el: '#app',
    store,
    router,
    render: h => h(App)
})