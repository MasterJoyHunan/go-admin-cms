<template>
    <div class="dashboard-container">
        <div class="dashboard-text">欢迎: {{ user.real_name }}</div>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: 'Dashboard',
    computed: {
        ...mapGetters([
            'user'
        ])
    }
}
</script>

<style lang="scss" scoped>
.dashboard {
    &-container {
        margin: 30px;
    }
    &-text {
        font-size: 30px;
        line-height: 46px;
    }
}
</style>
