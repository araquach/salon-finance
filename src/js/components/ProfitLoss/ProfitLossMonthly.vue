<template>
    <div>
        <table class="table is-narrow">
            <tr>
                <th>Months</th>
                <th v-for="(month in months">{{month}}</th>
            </tr>
            <tr>
                <th>Costs</th>
                <td v-for="(data, index) in monthCostTotal">{{data | toCurrency}}</td>
            </tr>
            <tr>
                <th>Takings</th>
                <td v-for="(data, index) in monthTakingsTotal">{{data | toCurrency}}</td>
            </tr>
            <tr>
                <th>Profit/Loss</th>
                <td v-for="(data, index) in monthProfitLoss">{{data | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>
<script>
    export default {

    }
</script>