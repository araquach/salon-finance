<template>
    <div>
        <h1 class="title is-1">
            PK Turnover
        </h1>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
            </tr>
            <tr v-for="(pk, index) in pkFigures">
                <td>{{pk.month_year}}</td>
                <td>{{pk.products}}</td>
                <td>{{pk.services}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                pkFigures: []
            }
        },

        created() {
            axios.get('/api/pk').then(response => this.pkFigures = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>