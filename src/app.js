import Vue from 'vue'
import Vuelidate from 'vuelidate'
import App from './App.vue'
import router from "./router"
import { store } from './store'
import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(Vuelidate)

window.axios = require('axios')

Vue.filter('toUpperCase', s => s[0].toUpperCase() + s.slice(1))

Vue.filter("toCurrency",  amount => Number(Math.round(amount)).toLocaleString('en-GB',
    {
        style: 'currency',
        currency: 'GBP'
    }))

Vue.filter("toRounded",   (n, amount) => Number(n.toFixed(amount)))

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
})