<template>
  <div class="section">
    <h2 class="title is-3">Total Takings 2021</h2>
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

import {parseISO, format} from "date-fns"

export default {
  data() {
    return {
      takings: [],
      totals: null
    }
  },

  filters: {
    showMonth: value => format(parseISO(value), 'LLLL')
  },

    created() {
      axios
          .get(`/api/takings-by-date-range/all/2021-07-01/2022-02-28`)
          .then(r => r.data)
          .then(data => {
            this.takings = data.figures
            this.totals = data.grand_totals
          })
  }
}
</script>