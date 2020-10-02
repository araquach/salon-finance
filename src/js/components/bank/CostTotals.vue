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
            <tr v-for="category in categoryTotal">
                <td>{{category.category}}</td>
                <td>{{category.amount | toCurrency}}</td>
                <td>{{category.percent.toFixed(1)}}</td>
                <td>{{category.average | toCurrency}}</td>
            <tr>
                <th>Grand Total</th>
                <td><strong>{{grandTotal | toCurrency}}</strong></td>
                <td></td>
                <td>{{totalAverage | toCurrency}}</td>
            </tr>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {

    }
</script>