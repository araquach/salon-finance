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
        this.$store.dispatch('loadTakingsByMonth')
        this.$store.dispatch('loadCostsByCat')
        this.$store.dispatch('loadCostsByMonth')
      }
    }
</script>