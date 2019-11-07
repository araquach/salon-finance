import Vue from 'vue'
import App from './App.vue'
import Takings from './components/Takings'
import SalonTakings from './components/SalonTakings'

import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(require('vue-moment'))

Vue.component('Takings', Takings)
Vue.component('SalonTakings', SalonTakings)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})