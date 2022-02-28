<template>
  <div class="section">
    <h2 class="title is-4">Stylist Takings by Date Range</h2>
    <table class="table is-size-5">
      <tr>
        <th>Month</th>
        <th>Services</th>
        <th>Products</th>
        <th class="has-text-weight-bold">Total</th>
      </tr>
      <tr v-for="total in totalsByMonth">
        <td>{{ total.month | showMonth }}</td>
        <td>{{ total.services | toCurrency }}</td>
        <td>{{ total.products | toCurrency }}</td>
        <td class="has-text-weight-bold has-text-warning">{{ total.services + total.products | toCurrency}}</td>
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
      totalsByStylist: state => state.takingsByStylist
    })
  }
}
</script>