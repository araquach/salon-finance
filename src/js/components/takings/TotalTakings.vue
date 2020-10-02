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
                <td>{{ takings.services | toCurrency}}</td>
                <td>{{ takings.products | toCurrency }}</td>
                <td><strong>{{ takings.total | toCurrency }}</strong></td>
                <td>{{ takings.average | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>
<script>
    import { mapState } from "vuex"

    export default {
      computed: {
        ...mapState(["takings"])
      },

      created() {
        this.$store.dispatch('loadTakings')
      }
    }
</script>