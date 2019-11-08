<template>
    <tr>
        <td>{{c.trans_date | moment("MMM")}}</td>
        <td>{{c.account_number}}</td>
        <td>{{c.trans_description}}</td>
        <td>{{c.debit_amount}}</td>
        <td v-if="!catagory"><button class="button is-small is-primary">Select</button></td>
        <td v-else >{{category}}</td>
    </tr>
</template>

<script>
    export default {
        props: ['c'],

        data() {
            return {
                catagory: ''
            }
        }
    }
</script>