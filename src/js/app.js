import Vue from 'vue'
import App from './App.vue'
import TotalTakings from './components/takings/TotalTakings'
import JakTakings from './components/takings/JakTakings'
import PkTakings from './components/takings/PkTakings'
import BaseTakings from './components/takings/BaseTakings'
import Costs from './components/bank/Costs'
import IndCost from './components/bank/IndCost'

import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(require('vue-moment'))

Vue.component('TotalTakings', TotalTakings)
Vue.component('JakTakings', JakTakings)
Vue.component('PkTakings', PkTakings)
Vue.component('BaseTakings', BaseTakings)
Vue.component('Costs', Costs)
Vue.component('IndCost', IndCost)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})