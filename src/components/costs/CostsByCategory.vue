<template>
  <div class="section">
    <h1 class="title is-5">Categorised Costs By Date Range</h1>
    <p class="is-size-4">Salon: {{ salon | toUpperCase }}</p>
    <p>Number of months: {{ months }}</p>
    <br>
    <table class="table is-size-5">
      <tr>
        <th>Category</th>
        <th>Amount</th>
        <th>Percent</th>
        <th>Monthly Average</th>
      </tr>
      <tr v-for="cost in figures">
        <td>{{ cost.category | toUpperCase }}</td>
        <td class="has-text-warning">{{ cost.total | toCurrency }}</td>
        <td>{{ cost.percent | toRounded(1)}}</td>
        <td>{{ cost.average | toCurrency }}</td>
      </tr>
      <tr class="is-size-3">
        <td>Grand Total</td>
        <td class="has-text-warning">{{ total | toCurrency }}</td>
      </tr>
      <tr v-if="months !== 12" class="is-size-3">
        <td>Yearly</td>
        <td class="has-text-warning">{{ byYear | toCurrency }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import {mapState} from "vuex"

export default {
  computed: {
    ...mapState({
      salon: state => state.costsByCat.salon,
      figures: state => state.costsByCat.figures,
      total: state => state.costsByCat.grand_total,
      byYear: state=> state.costsByCat.by_year,
      months: state => state.costsByCat.months
    })
  }
}
</script>