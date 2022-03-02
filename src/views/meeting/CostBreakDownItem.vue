<template>
  <tr>
    <td><button @click="revealAndAdapt" class="button">Show</button></td>
    <td v-if="showItem">{{ cost.category | toUpperCase }}</td>
    <td v-if="showItem" class="has-text-warning">{{ (cost.total / 8) * 12 | toCurrency }}</td>
    <td v-if="showItem">{{ cost.percent | toRounded(1)}}</td>
    <td v-if="showItem">{{ cost.average | toCurrency }}</td>
  </tr>
</template>
<script>
export default {
  props: ['cost'],

  data() {
    return {
      showItem: false
    }
  },

  methods: {
    revealAndAdapt() {
      this.$store.commit('ADJUST_TOTAL_TURNOVER', (this.cost.total / 8) * 12)
      this.showItem = true
    }
  }
}
</script>