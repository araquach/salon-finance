<template>
    <div>
        <TotalTakings/>
        <CostTotals/>
        <CostMain/>
    </div>
</template>
<script>
    import CostMain from "../components/bank/CostMain"
    import CostTotals from "../components/bank/CostTotals"
    import TotalTakings from "../components/takings/TotalTakings"
    export default {
        components: {CostMain, CostTotals, TotalTakings}
    }
</script>