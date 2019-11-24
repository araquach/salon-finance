import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import { routes } from "./routes"

import Buefy from 'buefy'
Vue.use(Buefy)
Vue.use(require('vue-moment'))
Vue.use(VueRouter)

const router = new VueRouter({
    routes
})

window.axios = require('axios')

Vue.filter("toCurrency",  amount => "£" + Number(amount).toLocaleString())

new Vue({
    el: '#app',
    router,
    render: h => h(App)
})