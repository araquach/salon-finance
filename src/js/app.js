import Vue from 'vue'
import App from './App.vue'
import BankComponent from './components/BankComponent'

import Buefy from 'buefy'

Vue.use(Buefy)

Vue.component('BankComponent', BankComponent)


window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})