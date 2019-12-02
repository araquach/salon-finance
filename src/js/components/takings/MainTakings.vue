<template>
    <div class="section">
        <jakTakings/>
        <pkTakings/>
        <baseTakings/>
        <totalTakings/>
    </div>
</template>
<script>
    import JakTakings from "./JakTakings";
    import PkTakings from "./PkTakings";
    import BaseTakings from "./BaseTakings";
    import TotalTakings from "./TotalTakings";
    export default {
        components: {JakTakings, PkTakings, BaseTakings, TotalTakings}
    }
</script>