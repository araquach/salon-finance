<template>
  <div class="section">
    <h1 class="title is-4">Turnover/Outgoings Chart</h1>
    <line-chart :data="costs"></line-chart>
    {{ costs }}
  </div>
</template>
<script>
export default {
  data() {
    return {
      costs: [],
      takings: []
    }
  },

  computed: {

  },

  created() {
    axios
        .get(`/api/costs-and-takings/2017-06-01/2022-02-28`)
        .then(r => r.data)
        .then(data => {
          this.costs = data.costs
          this.takings = data.takings
        })
  }
}
</script>
