<template>
  <div class="container">
    <Header/>
    <router-view/>
  </div>
</template>

<script>
import Header from './components/Header'
import Selector from "./components/other/Selector"


export default {
  components: {Header, Selector},

  // created() {
  //   this.$store.dispatch('loadTakingsByStylist')
  //   this.$store.dispatch('loadTakingsByDateRange')
  //   this.$store.dispatch('loadTotalsByDateRange')
  //   this.$store.dispatch('loadCostsByCat')
  //   this.$store.dispatch('loadCostsByDateRange')
  // }
}
</script>