<template>
    <div class="section">
        <h2 class="title is-3">
            Base Turnover
        </h2>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
                <th>Total</th>
            </tr>
            <tr v-for="(t, index) in t">
                <td>{{t.month_year | moment("MMMM")}}</td>
                <td>{{t.services | toCurrency }}</td>
                <td>{{t.products | toCurrency}}</td>
                <td><strong>{{t.total | toCurrency}}</strong></td>
            </tr>
        </table>
        <p class="is-size-3">Total: {{total | toCurrency}}</p>
    </div>
</template>
<script>
    export default {

    }
</script>