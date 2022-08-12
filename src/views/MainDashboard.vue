<template>
  <div>
    <router-view/>
  </div>
</template>
<script>
export default {
  created() {
    this.$store.dispatch('loadTakingsByStylist')
    this.$store.dispatch('loadTakingsByMonth')
    this.$store.dispatch('loadCostsByCat')
    this.$store.dispatch('loadCostsByMonth')
  }
}
</script>