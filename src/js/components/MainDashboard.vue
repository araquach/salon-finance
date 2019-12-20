<template>
    <div>
        <TotalTakings/>
        <CostTotals/>
        <CostMain/>
    </div>
</template>
<script>
    import CostMain from "./bank/CostMain"
    import CostTotals from "./bank/CostTotals1"
    import TotalTakings from "./takings/TotalTakings";
    export default {
        components: {CostMain, CostTotals, TotalTakings}
    }
</script>