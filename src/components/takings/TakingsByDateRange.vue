<template>
  <div class="section">
    <h2 class="title is-4">Total Takings by Date Range</h2>
    <p class="is-size-5">Salon: {{ salon | toUpperCase }}</p>
    <p>Number of months: {{ months }}</p>
    <br>
    <table class="table is-size-5">
      <tr>
        <th>Month</th>
        <th>Services</th>
        <th>Products</th>
        <th class="has-text-weight-bold">Total</th>
      </tr>
      <tr v-for="total in takings">
        <td>{{ total.month | showMonth }}</td>
        <td>{{ total.services | toCurrency }}</td>
        <td>{{ total.products | toCurrency }}</td>
        <td class="has-text-weight-bold has-text-warning">{{ total.services + total.products | toCurrency}}</td>
      </tr>
      <tr class="is-size-4">
        <td>Total Services</td>
        <td class="has-text-warning">{{ totals.services | toCurrency}}</td>
      </tr>
      <tr class="is-size-4">
        <td>Total Products</td>
        <td class="has-text-warning">{{ totals.products | toCurrency }}</td>
      </tr>
      <tr class="is-size-3">
        <td>Grand Total</td>
        <td class="has-text-warning">{{ totals.grand_total | toCurrency}}</td>
      </tr>
      <tr v-if="months !== 12" class="is-size-3">
        <td>Yearly Total</td>
        <td class="has-text-warning">{{ totals.yearly | toCurrency}}</td>
      </tr>
    </table>
  </div>
</template>
<script>
import {mapState} from "vuex"
import {parseISO, format} from "date-fns"

export default {
  filters: {
    showMonth: value => format(parseISO(value), 'LLLL')
  },

  computed: {
    ...mapState({
      salon: state => state.takingsByDateRange.salon,
      takings: state => state.takingsByDateRange.figures,
      months: state => state.takingsByDateRange.months,
      totals: state => state.takingsByDateRange.grand_totals
    })
  }
}
</script>