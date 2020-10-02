<template>
    <b-navbar spaced>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '' }">
                <h1 class="title is-3">Salon Finance</h1>
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" to="/">
                Dashboard
            </b-navbar-item>
            <b-navbar-dropdown label="Menu">
                <b-navbar-item tag="router-link" to="/totals">
                    Totals
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/profit-loss">
                    Monthly Profit/loss
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">

        </template>
    </b-navbar>
</template>