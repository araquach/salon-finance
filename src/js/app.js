import Vue from 'vue'
import App from './App.vue'
import TestComponent from './components/TestComponent'
import SecondComponent from './components/SecondComponent'

import Buefy from 'buefy'

Vue.use(Buefy)

Vue.component('TestComponent', TestComponent)
Vue.component('SecondComponent', SecondComponent)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
});