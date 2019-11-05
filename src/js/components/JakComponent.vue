<template>
    <div>
        <h1 class="title is-1">
            Jakata Turnover
        </h1>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
            </tr>
            <tr v-for="(jak, index) in jakFigures">
                <td>{{jak.month_year}}</td>
                <td>{{jak.products}}</td>
                <td>{{jak.services}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                jakFigures: []
            }
        },

        created() {
            axios.get('/api/jakata').then(response => this.jakFigures = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>