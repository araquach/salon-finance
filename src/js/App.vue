<template>
    <div class="container">
        <BaseComponent :bf="bf"/>
        <JakComponent :jf="jf"/>
        <PkComponent :pf="pf"/>
        <totalsComponant/>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                bf: [],
                jf: [],
                pf: []
            }
        },

        created() {
            axios.get('/api/base').then(response => this.bf = response.data)
                .catch(error => {
                    console.log(error)
                }),

            axios.get('/api/jakata').then(response => this.jf = response.data)
                .catch(error => {
                    console.log(error)
                }),

            axios.get('/api/jakata').then(response => this.pf = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>
