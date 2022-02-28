<template>
  <div class="section">
    <h1 class="title is-4">Categorised Costs By Date Range</h1>
    <p>Number of months: {{ months }}</p>
    <table class="table is-size-5">
      <tr>
        <th>Category</th>
        <th>Amount</th>
        <th>Percent</th>
        <th>Monthly Average</th>
      </tr>
      <tr v-for="cost in figures">
        <td>{{ cost.category | upperCaseFirst }}</td>
        <td class="has-text-warning">{{ cost.total | toCurrency }}</td>
        <td>{{ cost.percent | toRounded(1)}}</td>
        <td>{{ cost.average | toCurrency }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import {mapState} from "vuex"

export default {
  filters: {
    upperCaseFirst: value => value[0].toUpperCase() + value.slice(1)
  },

  computed: {
    ...mapState({
      figures: state => state.costsByCat.figures,
      total: state => state.costsByCat.total,
      months: state => state.costsByCat.months
    })
  }
}
</script>