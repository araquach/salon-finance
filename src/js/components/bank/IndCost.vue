<template>
    <tr>
        <td>{{c.trans_date | moment("MMM")}}</td>
        <td>{{c.account_number}}</td>
        <td>{{c.trans_description}}</td>
        <td>{{c.debit_amount.toFixed(2)}}</td>
        <td>
            <input v-model="category" type="text" class="input">
        </td>
        <td>
            <button @click="addCategory" class="button is-small is-primary">Add Catagory</button>
        </td>
        <td>{{c.category}}</td>
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
                axios.put('/api/bankdata/' + this.c.id, {
                    category: this.category
                })
                    .then(response => {
                        console.log(this.category)
                    })
                    .catch(error => {
                        console.log(err)
                    })
            }
        }
    }
</script>