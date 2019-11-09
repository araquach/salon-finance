<template>
    <tr>
        <td>{{c.trans_date | moment("MMM")}}</td>
        <td>{{c.account_number}}</td>
        <td>{{c.trans_description}}</td>
        <td>{{c.debit_amount.toFixed(2)}}</td>
        <td :catagory="catagory" v-if="!catagory">
            <button @click="addCatagory" class="button is-small is-primary">Select</button>
        </td>
        <td :catagory="catagory" v-else >{{catagory}}</td>
    </tr>
</template>

<script>
    export default {
        props: ['c'],

        data() {
            return {
                catagory: ''
            }
        },

        methods: {
            addCatagory() {
                this.catagory = "Bank Charge"

                axios.put('/api/bankdata', {
                    category: "stock"
                })
                    .then(response => {
                        console.log(response);
                    })
                    .catch(error => {
                        console.log(err);
                    })
            }
        },

        created() {

        }

    }
</script>