<template>
    <div class="container">
        <Header/>
        <router-view/>
    </div>
</template>

<script>
    import Header from './components/Header'
    export default {
        components: {Header},

      created() {
        this.$store.dispatch('loadTakingsByStylist')
        this.$store.dispatch('loadTakingsByDateRange')
        this.$store.dispatch('loadTotalsByDateRange')
        this.$store.dispatch('loadCostsByCat')
        this.$store.dispatch('loadCostsByDateRange')
      }
    }
</script>