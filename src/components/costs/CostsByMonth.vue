<template>
    <div class="section">
        <table class="table is-size-4">

        </table>
    </div>
</template>

<script>
  import { mapState } from "vuex"

  export default {
    computed: {
      ...mapState({
        costs: state => state.costsByDateRange
      }),
    }
  }
</script>