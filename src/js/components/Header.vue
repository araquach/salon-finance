<template>
    <b-navbar spaced>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ name: 'dashboard' }">
                <h1 class="title is-3">Salon Finance</h1>
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" :to="{ name: 'dashboard' }">
                Dashboard
            </b-navbar-item>
            <b-navbar-dropdown label="Menu">
                <b-navbar-item tag="router-link" :to="{name: 'total-takings'}">
                    Total Turnover Date Range
                </b-navbar-item>
                <b-navbar-item tag="router-link" :to="{name: 'cost-totals'}">
                    Categorised Costs
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">

        </template>
    </b-navbar>
</template>