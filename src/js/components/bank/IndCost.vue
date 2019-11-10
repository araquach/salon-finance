<template>
    <tr>
        <td>{{c.trans_date | moment("MMM")}}</td>
        <td>{{c.account_number}}</td>
        <td>{{c.trans_description}}</td>
        <td>{{c.debit_amount.toFixed(2)}}</td>
        <td :catagory="category" v-if="!category">
            <button @click="addCategory" class="button is-small is-primary">Select</button>
        </td>
        <td :catagory="category" v-else >{{category}}</td>
    </tr>
</template>

<script>
    export default {
        props: ['c'],

        data() {
            return {
                category: ''
            }
        },

        methods: {
            addCategory() {
                axios.put('/api/bankdata/' + this.category.id, {
                    category: "Bank Charge"
                })
                    .then(response => {
                        console.log(response)
                    })
                    .catch(error => {
                        console.log(err)
                    })
            }
        },

        created() {

        }

    }
</script>