import Vue from 'vue'
import App from './App.vue'
import TotalTakings from './components/TotalTakings'
import JakTakings from './components/JakTakings'
import PkTakings from './components/PkTakings'
import BaseTakings from './components/BaseTakings'

import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(require('vue-moment'))

Vue.component('TotalTakings', TotalTakings)
Vue.component('JakTakings', JakTakings)
Vue.component('PkTakings', PkTakings)
Vue.component('BaseTakings', BaseTakings)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})