<template>
  <div class="register-container">
    <el-form ref="registerForm" :model="registerForm" :rules="registerRules" class="register-form" autocomplete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">Register</h3>
      </div>

      <el-form-item prop="name">
        <span class="iconfont el-icon-thirduser svg-container"></span>
        <el-input
          ref="name"
          v-model="registerForm.name"
          placeholder="Username"
          name="username"
          type="text"
          tabindex="1"
          autocomplete="on"
        />
      </el-form-item>

      <el-form-item prop="phone">
        <span class="iconfont el-icon-thirdphone svg-container"></span>
        <el-input
          ref="phone"
          v-model="registerForm.phone"
          placeholder="Phone"
          name="phone"
          type="text"
          tabindex="2"
          autocomplete="on"
        />
      </el-form-item>

      <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
        <el-form-item prop="password">
          <span class="iconfont el-icon-thirdpassword svg-container"></span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="registerForm.password"
            :type="passwordType"
            placeholder="Password"
            name="password"
            tabindex="3"
            autocomplete="on"
            @keyup.native="checkCapslock"
            @blur="capsTooltip = false"
            @keyup.enter.native="submitForm"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
      </el-tooltip>

      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="submitForm">Register</el-button>
      <button class="loginTip" @click="login">直接登录</button>
    </el-form>
  </div>
</template>

<script>

export default {
  name: 'Register',
  data() {
    // 验证账户
    const validateName = (rule, value, callback) => {
      if (value==="") {
        callback(new Error('Please enter the correct user name'))
      } else {
        callback()
      }
    }
    // 验证手机
    const validatePhone = (rule, value, callback) => {
      if (value==="") {
        callback(new Error('Please enter your phone'))
      } else if(value.length!=11){
        callback(new Error('The format of phone is wrong'))
      }else{
        callback()
      }
    }
    // 验证密码
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else if(value.length>16) {
        callback(new Error('The password can not be more than 16 digits'))
      }else{
        callback()
      }
    }
    return {
      registerForm: {
        name: '',
        phone: '',
        password: ''
      },
      // 登录验证规则
      registerRules: {
        name: [{ required: true, trigger: 'blur', validator: validateName }],
        phone: [{ required: true, trigger: 'blur', validator: validatePhone }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      passwordType: 'password',
      capsTooltip: false,
      loading: false,
      showDialog: false,
      redirect: undefined,
      otherQuery: {}
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        const query = route.query
        if (query) {
          this.redirect = query.redirect
          this.otherQuery = this.getOtherQuery(query)
        }
      },
      immediate: true
    }
  },
  created() {
    // window.addEventListener('storage', this.afterQRScan)
  },
  mounted() {
    if (this.registerForm.name === '') {
      this.$refs.name.focus()
    } else if (this.registerForm.password === '') {
      this.$refs.password.focus()
    }
  },
  destroyed() {
    // window.removeEventListener('storage', this.afterQRScan)
  },
  methods: {
    checkCapslock(e) {
      const { key } = e
      this.capsTooltip = key && key.length === 1 && (key >= 'A' && key <= 'Z')
    },
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
    login(){
      this.$router.push({ path: "/login" });
    },
    // 处理登录逻辑
    submitForm() {
      this.$refs.registerForm.validate(valid => {
        if (valid) {
          // 验证成功
          this.loading = true
          this.form = new FormData()
          this.form.append('name', this.registerForm.name)
          this.form.append('password', this.registerForm.password)
          // 请求后端认证用户信息
          const { data } = this.$post('/user/register', this.form)
            .then((data) => {
              // 用户验证成功
              if (data.code === 200 && data.success === true) {
                this.$message({
                  type: 'success',
                  message: data.message
                })
                // 浏览器的缓存空间有两种： localStorage和sessionStorage
                sessionStorage.setItem('token', 'InnerMongoliaUniversity' + ' ' + data.token)
                // 将token保存到请求头
                this.$setToken()
                // 请求后端获取用户角色
                this.$router.push({ path: "/home" });
              } else {
                this.$message({
                  type: 'error',
                  message: data.message
                })
                return false
              }
            }).catch(() => {
              this.loading = false
            })
          console.log(data)
        } else {
          this.$message({
            type: 'error',
            message: '用户格式错误'
          })
          return false
        }
      })
    },
    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {})
    }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg:#283443;
$light_gray:#fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .register-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.register-container {
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
$bg:#2d3a4b;
$dark_gray:#889aa4;
$light_gray:#eee;

.register-container {
  position: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  background: url("../assets/img/background.jpg") center no-repeat #efeff4;
  overflow: hidden;

  .register-form {
    background: rgba($color: #000000, $alpha: 0.5);
    position: relative;
    height: 400px;
    width: 520px;
    max-width: 100%;
    padding: 100px 35px 0;
    margin: 0 auto;
    overflow: hidden;
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

  .thirdparty-button {
    position: absolute;
    right: 0;
    bottom: 6px;
  }

  @media only screen and (max-width: 470px) {
    .thirdparty-button {
      display: none;
    }
  }
}
.loginTip{
  background: rgba(0, 0, 0, 0);
  border: rgba(0, 0, 0, 0);
  font-size: 14px;
  color: #409EFF;
  float: right;
}
</style>
