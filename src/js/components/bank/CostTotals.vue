<template>
    <div class="section">
        <h1 class="title is-4">Costs</h1>
        <table class="table">
            <tr>
                <th>Category</th>
                <th>Amount</th>
                <th>Percent</th>
                <th>Average</th>
            </tr>
            <tr v-for="cost in costs.individual">
                <td>{{ cost.category }}</td>
                <td>{{ cost.amount | toCurrency }}</td>
                <td>{{ cost.percent.toFixed(1) }}</td>
                <td>{{ cost.average | toCurrency }}</td>

            <tr>
                <th>Grand Total</th>
                <td class="is-size-5"><strong>{{ costs.total | toCurrency }}</strong></td>
                <td></td>
                <td></td>
            </tr>
        </table>
    </div>
</template>

<script>
  import { mapState, mapGetters } from "vuex"

  export default {
    computed: {
      ...mapState(['costs']),
    },

    created() {
      this.$store.dispatch('loadCosts')
    }
  }
</script>