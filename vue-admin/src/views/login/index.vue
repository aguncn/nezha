<template>
  <div class="login-container">
    <el-tabs v-model="activeName" @tab-click="handleTabClick">
      <el-tab-pane label="登录" :key="'login'" name="login" >
        <el-form  v-if="tabLogin" ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
          <el-form-item prop="username">
            <span class="svg-container">
              <svg-icon icon-class="user" />
            </span>
            <el-input
              ref="username"
              v-model="loginForm.username"
              placeholder="Username"
              name="username"
              type="text"
              tabindex="1"
              auto-complete="on"
            />
          </el-form-item>
          <el-form-item prop="password">
            <span class="svg-container">
              <svg-icon icon-class="password" />
            </span>
            <el-input
              :key="passwordType"
              ref="password"
              v-model="loginForm.password"
              :type="passwordType"
              placeholder="Password"
              name="password"
              tabindex="2"
              auto-complete="on"
              @keyup.enter.native="handleLogin"
            />
            <span class="show-pwd" @click="showPwd">
              <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
            </span>
          </el-form-item>
          <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleLogin">登陆</el-button>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="注册" :key="'register'" name="register" >
        <el-form  v-if="tabRegister" ref="registerForm" :model="registerForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
          <el-form-item prop="username">
            <span class="svg-container">
              <svg-icon icon-class="user" />
            </span>
            <el-input
              ref="username"
              v-model="registerForm.username"
              placeholder="Username"
              name="username"
              type="text"
              tabindex="1"
              auto-complete="on"
            />
          </el-form-item>
          <el-form-item prop="password">
            <span class="svg-container">
              <svg-icon icon-class="password" />
            </span>
            <el-input
              :key="passwordType"
              ref="password"
              v-model="registerForm.password"
              :type="passwordType"
              placeholder="Password"
              name="password"
              tabindex="2"
              auto-complete="on"
              @keyup.enter.native="handleRegister"
            />
            <span class="show-pwd" @click="showPwd">
              <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
            </span>
          </el-form-item>
          <el-form-item prop="password">
            <span class="svg-container">
              <svg-icon icon-class="password" />
            </span>
            <el-input
              :key="passwordType"
              ref="password1"
              v-model="registerForm.password1"
              :type="passwordType"
              placeholder="Password"
              name="password1"
              tabindex="3"
              auto-complete="on"
              @keyup.enter.native="handleRegister"
            />
            <span class="show-pwd" @click="showPwd">
              <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
            </span>
          </el-form-item>
          <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleRegister">注册</el-button>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>

export default {
  name: 'Login',
  data() {
    const validatePwd = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('密码长度少于6位！'))
      } else {
        callback()
      }
    }
    return {
      activeName: "login",
      tabLogin:true,
      tabRegister:false,
      loginForm: {
        username: 'admin',
        password: '111111'
      },
      registerForm: {
        username: 'aguncn',
        password: '111111',
        password1: '111111',
      },
      loginRules: {
        password: [{ required: true, trigger: 'blur', validator: validatePwd }]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleTabClick(tab) {
      if (tab.name ==='login') {
        this.tabLogin =true
        this.tabRegister =false
      }else if (tab.name ==='register') {
        this.tabLogin =false
        this.tabRegister =true
      }
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm).then(() => {
            this.$router.push({ path: this.redirect || '/' })
            this.loading = false
          }).catch(() => {
            this.$message.error('用户名或密码错误，请重试！')
            this.loading = false
          })
        } else {
          return false
        }
      })
    },
    handleRegister() {
      this.$refs.registerForm.validate(valid => {
        if (valid) {
          if (this.registerForm.password !== this.registerForm.password1) {
            this.$message.error('两次密码输入不一致。')
            return false
          } else {
            this.loading = true
            this.$store.dispatch('user/register', this.registerForm).then((data) => {
              if (data.code === 200) {
                this.$message.success('注册成功，请登陆。')
              } else if (data.code === 30001) {
                this.$message.error('亲，用户已存在，请选择其它用户名！')
              } else {
                this.$message.error('未知错误！')
              }
              this.loading = false
            }).catch((e) => {
              this.$message.error('未知错误！')
              return false
            })
          }
        } else {
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss">
$dark_gray:#889aa4;
$light_gray:#fff;
$cursor: #fff;
$bg:#283443;

.login-container {
		-webkit-border-radius: 5px;
		border-radius: 5px;
		-moz-border-radius: 5px;
		background-clip: padding-box;
		margin: 100px auto;
		width: 500px;
		padding: 35px 35px 15px 35px;
		background: #fff;
		border: 1px solid #eaeaea;
		box-shadow: 0 0 25px #cac6c6;
	.title {
		margin: 0px auto 30px auto;
		text-align: center;
		color: #505458;
	}
  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    cursor: pointer;
    user-select: none;
  }
  .svg-container {
    padding: 6px 5px 6px 15px;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }
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
      height: 47px;

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
  }
}
</style>
