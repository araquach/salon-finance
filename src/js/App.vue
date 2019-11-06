<template>
    <div class="container">
        <TakingsComponent :t="t"/>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                t: []
            }
        },

        created() {
            axios.get('/api/takings').then(response => this.t = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>
