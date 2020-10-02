<template>
    <div class="section">
        <h2 class="title is-4">
            Total Salon Turnover
        </h2>
        <table class="table is-size-5">
            <tr>
                <th>Services</th>
                <th>Products</th>
                <th><strong>Grand Total</strong></th>
                <th>Monthly Average</th>
            </tr>
            <tr>
                <td>{{ totalTakings.services | toCurrency}}</td>
                <td>{{ totalTakings.products | toCurrency }}</td>
                <td><strong>{{ totalTakings.total | toCurrency }}</strong></td>
                <td>{{ totalTakings.average | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>
<script>
    import { mapState } from "vuex"

    export default {
      computed: {
        ...mapState(["totalTakings"])
      },

      created() {
        this.$store.dispatch('loadTotalTakings')
      }
    }
</script>