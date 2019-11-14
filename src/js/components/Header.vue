<template>
    <b-navbar>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '' }">
                <h1 class="title is-3 has-text-primary">Salon Finance</h1>
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" to="/">
                Home
            </b-navbar-item>
            <b-navbar-item tag="router-link" to="/costs">
                Costs
            </b-navbar-item>
            <b-navbar-dropdown label="Info">
                <b-navbar-item tag="router-link" to="/input">
                    Input
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/total">
                    Total
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/jakata">
                    Jakata
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/pk">
                    PK
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/base">
                    Base
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>

        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <a class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>