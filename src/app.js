import Vue from 'vue'
import Vuelidate from 'vuelidate'
import App from './App.vue'
import router from "./router"
import { store } from './store'
import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(Vuelidate)

window.axios = require('axios')

Vue.filter('textLimit', function (text, length) {
    return text.substring(0, length)
})

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