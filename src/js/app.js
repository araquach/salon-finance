import Vue from 'vue'
import App from './App.vue'
import TakingsComponent from './components/TakingsComponent'

import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(require('vue-moment'))

Vue.component('TakingsComponent', TakingsComponent)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})