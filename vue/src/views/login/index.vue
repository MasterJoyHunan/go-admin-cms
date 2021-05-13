<template>
    <div class="login-container">
        <el-form ref="loginForm"
                 :model="loginForm"
                 :rules="loginRules"
                 class="login-form"
                 auto-complete="on"
                 label-position="left">

            <div class="title-container">
                <h3 class="title">Go Admin CMS 系统 V 1.0</h3>
            </div>

            <el-form-item prop="username">
                <span class="svg-container">
                    <svg-icon icon-class="user" />
                </span>
                <el-input ref="username"
                          v-model="loginForm.username"
                          type="text" />
            </el-form-item>

            <el-form-item prop="password">
                <span class="svg-container">
                    <svg-icon icon-class="password" />
                </span>
                <el-input :key="passwordType"
                          ref="password"
                          v-model="loginForm.password"
                          :type="passwordType"
                          @keyup.enter.native="handleLogin()" />
                <span class="show-pwd"
                      @click="showPwd">
                    <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
                </span>
            </el-form-item>

            <el-form-item prop="verify_code">
                <span class="svg-container">
                    <svg-icon icon-class="verify" />
                </span>
                <el-input ref="verify_code"
                          v-model="loginForm.verify_code"
                          type="text"
                          @keyup.enter.native="handleLogin()" />
                <span class="show-virify"
                      @click="changeVerify()">
                    <img :src="verify_code">
                </span>
            </el-form-item>

            <el-button :loading="loading"
                       type="primary"
                       style="width:100%;margin-bottom:30px;"
                       @click.native.prevent="handleLogin()">登录</el-button>
        </el-form>
    </div>
</template>

<script>
export default {
    name: 'Login',
    data() {
        return {
            loginForm: {
                username: 'admin',
                password: 'admin',
                verify_code: '9527'
            },
            loginRules: {
                user_name: [{ required: true, trigger: 'blur', message: '请输入用户名' }],
                password: [{ required: true, trigger: 'blur', min: 4, message: '长度至少为4位' }],
                verify_code: [{ required: true, trigger: 'blur', len: 4, message: '请输入合法的验证码' }]
            },
            loading: false,
            passwordType: 'password',
            redirect: undefined,
            verify_code: process.env.VUE_APP_BASE_API + '/captcha?' + Math.random()
        }
    },
    watch: {
        $route: {
            handler(route) {
                this.redirect = route.query && route.query.redirect
            },
            immediate: true
        }
    },
    methods: {
        // 显示密码
        showPwd() {
            this.passwordType = this.passwordType === 'password' ? '' : 'password'
            this.$nextTick(() => {
                this.$refs.password.focus()
            })
        },
        // 刷新验证码
        changeVerify() {
            this.verify_code = process.env.VUE_APP_BASE_API + '/captcha?' + Math.random()
        },
        // 登录
        async handleLogin() {
            await this.$refs.loginForm.validate()
            this.loading = true
            try {
                await this.$store.dispatch('user/login', this.loginForm)
                this.$router.push({ path: this.redirect || '/' })
            } finally {
                this.changeVerify()
                this.loading = false
            }
        }
    }
}
</script>

<style lang="scss">
$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .login-container .el-input input {
        color: $cursor;
    }
}

.login-container {
    .el-input {
        display: inline-block;
        height: 47px;
        width: 85%;

        input {
            background: transparent;
            border: 0px;
            -webkit-appearance: none;
            border-radius: 0px;
            padding: 12px 5px 12px 15px;
            color: $light_gray;
            height: 47px;
            caret-color: $cursor;

            &:-webkit-autofill {
                box-shadow: 0 0 0px 1000px $bg inset !important;
                -webkit-text-fill-color: $cursor !important;
            }
        }
    }

    .el-form-item {
        border: 1px solid rgba(255, 255, 255, 0.1);
        background: rgba(0, 0, 0, 0.1);
        border-radius: 5px;
        color: #454545;
    }
}
</style>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
    min-height: 100%;
    width: 100%;
    background-color: $bg;
    overflow: hidden;

    .login-form {
        position: relative;
        width: 520px;
        max-width: 100%;
        padding: 160px 35px 0;
        margin: 0 auto;
        overflow: hidden;
    }

    .tips {
        font-size: 14px;
        color: #fff;
        margin-bottom: 10px;

        span {
            &:first-of-type {
                margin-right: 16px;
            }
        }
    }

    .svg-container {
        padding: 6px 5px 6px 15px;
        color: $dark_gray;
        vertical-align: middle;
        width: 30px;
        display: inline-block;
    }

    .title-container {
        position: relative;

        .title {
            font-size: 26px;
            color: $light_gray;
            margin: 0px auto 40px auto;
            text-align: center;
            font-weight: bold;
        }
    }

    .show-pwd {
        position: absolute;
        right: 10px;
        top: 7px;
        font-size: 16px;
        color: $dark_gray;
        cursor: pointer;
        user-select: none;
    }
    .show-virify {
        position: absolute;
        right: 2px;
        top: 2px;
        bottom: 2px;
        font-size: 16px;
        color: $dark_gray;
        cursor: pointer;
        user-select: none;
        img {
            border-radius: 2px;
            width: 100%;
            height: 100%;
        }
    }
}
</style>
