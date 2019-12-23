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
                <b-navbar-item tag="router-link" to="/profitloss">
                    Monthly Profit/loss
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <b-navbar-item tag="router-link" to="/input">
                        Category Input
                    </b-navbar-item>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>