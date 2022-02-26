<template>
  <div class="section">
    <h1 class="title is-4">Categorised Costs By Date Range</h1>
    <table class="table is-size-5">
      <tr>
        <th>Category</th>
        <th>Amount</th>
        <th>Percent</th>
        <th>Average</th>
      </tr>
      <tr v-for="cost in costsByCat" :key="costsByCat.index">
        <td>{{ cost.category | upperCaseFirst }}</td>
        <td class="has-text-warning">{{ cost.total | toCurrency }}</td>
        <td>Percent</td>
        <td>Average</td>
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
      costsByCat: state => state.costsByCat
    })
  }
}
</script>