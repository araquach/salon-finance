import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import { routes } from "./routes"

import Buefy from 'buefy'
Vue.use(Buefy)
Vue.use(require('vue-moment'))
Vue.use(VueRouter)

// Vue.prototype.numMonths = 14

const router = new VueRouter({
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
    router,
    render: h => h(App)
})