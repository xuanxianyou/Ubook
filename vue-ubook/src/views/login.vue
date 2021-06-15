<template>
  <div class="loginContainer">
    <div class="loginInner">
      <div class="login_header">
        <div class="login_logo">Login</div>
        <div class="login_header_title">
          <a
            href="javascript:;"
            :class="{ on: loginWay }"
            @click="loginWay = true"
            >短信登录</a
          >
          <a
            href="javascript:;"
            :class="{ on: !loginWay }"
            @click="loginWay = false"
            >密码登录</a
          >
        </div>
      </div>
      <!-- 内容部分 -->
      <div class="login_content">
        <form @submit.prevent="login">
          <!-- 短信登录 -->
          <div :class="{ on: loginWay }">
            <section class="login_message">
              <input
                type="tel"
                minlength="11"
                maxlength="11"
                placeholder="Phone"
                v-model="phone"
              />
              <button
                :disabled="!rightPhone"
                class="get_verification"
                :class="{ right_phone: rightPhone }"
                @click.prevent="getCode"
              >
                {{ computeTime > 0 ? `(${computeTime}s)已发送` : "获取验证码" }}
              </button>
            </section>
            <section class="login_verification">
              <input
                type="tel"
                minlength="4"
                maxlength="4"
                placeholder="验证码"
                v-model="code"
              />
            </section>
            <section class="login_hint">
              温馨提示：未注册本系统的手机号，登录时将自动注册，且代表已同意
              <a href="javascript:;">《用户服务协议》</a>
            </section>
          </div>
          <!-- 密码登录 -->
          <div :class="{ on: !loginWay }">
            <section class="login_message">
              <input
                type="tel"
                maxlength="11"
                placeholder="Username / Phone"
                v-model="username"
              />
            </section>
            <section class="login_verification">
              <input
                type="text"
                minlength="8"
                maxlength="16"
                placeholder="Password"
                v-if="showPwd"
                v-model="password"
              />
              <input
                type="password"
                minlength="8"
                maxlength="16"
                placeholder="Password"
                v-else
                v-model="password"
              />
              <section
                class="switch_button"
                :class="showPwd ? 'on' : 'off'"
                @click="showPwd = !showPwd"
              >
                <section
                  class="switch_circle"
                  :class="{ right: showPwd }"
                ></section>
                <section class="switch_text">
                  {{ showPwd ? "" : "" }}
                </section>
              </section>
            </section>
          </div>
          <button class="login_submit primary">Login</button>
        </form>
        <button class="registerTip" @click="register">还未注册？</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      loginWay: false, //true代表短信登陆, false代表密码
      phone: "", //手机号,
      computeTime: 0,
      code: "", //验证码
      timer: null,
      showPwd: false,
      username: "",
      password: "",
    };
  },
  computed: {
    rightPhone() {
      //利用正则对手机号匹配
      return /^1[3456789]\d{9}$/.test(this.phone);
    },
  },
  methods: {
    getCode() {
      // 向后端提交请求
      this.form = new FormData()
      this.form.append('phone', this.phone)
      console.log(this.phone)
      const { data } = this.$post('/user/captcha',this.form)
            .then((data) => {
              // 获取验证码
              if (data.code === 200 && data.success === true) {
                this.$message({
                  type: 'success',
                  message: data.message
                })
              } else {
                this.$message({
                  type: 'error',
                  message: data.message
                })
                return false
              }
            }).catch(() => {
              
            })
          console.log(data)
      // 计时    
      if (!this.computeTime) {
        this.computeTime = 60;
        this.timer = setInterval(() => {
          this.computeTime--;
          if (this.computeTime <= 0) {
            clearInterval(this.timer);
          }
        }, 1000);
      }
    },
    login() {
      //短信验证
      if (this.loginWay) {
        console.log(this.rightPhone);
        if (!this.rightPhone) {
          this.$message({
            type: 'error',
            message: "请正确填写手机号"
          })
        } else if (!/^\d{4}$/.test(this.code)) {
          this.$message({
            type: 'error',
            message: "验证码必须是4位"
          })
        }else{
          this.form = new FormData()
          this.form.append('mode', 0)
          this.form.append('identity', this.phone)
          this.form.append('voucher', this.code)
          const { data } = this.$post('/user/login',this.form)
            .then((data) => {
              // 登录成功
              if (data.code === 200 && data.success === true) {
                this.$message({
                  type: 'success',
                  message: data.message
                })
                this.$router.push({ path: "/home" });
              } else {
                this.$message({
                  type: 'error',
                  message: data.message
                })
                return false
              }
            }).catch(() => {
              
            })
          console.log(data)
        }
      } else {
        //密码验证
        if (!/^[\w]{2,11}$/.test(this.username)) {
          this.$message({
            type: 'error',
            message: "用户名长度必须2-11位"
          })
        } else if (!/^\w{8,16}$/.test(this.pwd)) {
          this.$message({
            type: 'error',
            message: "密码长度必须8-16位"
          })
        }else{
          this.form = new FormData()
          this.form.append('mode', 1)
          this.form.append('identity', this.username)
          this.form.append('voucher', this.password)
          const { data } = this.$post('/user/login',this.form)
            .then((data) => {
              // 登录成功
              if (data.code === 200 && data.success === true) {
                this.$message({
                  type: 'success',
                  message: data.message
                })
                this.$router.push({ path: "/home" });
              } else {
                this.$message({
                  type: 'error',
                  message: data.message
                })
                return false
              }
            }).catch(() => {
              
            })
          console.log(data)
        }
      }
    },
    register() {
       this.$router.push({ path: "/register" });
    }
  },
};
</script>

<style scoped>
.loginContainer {
  position: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background-color: #fff;
  background: url("../assets/img/background.jpg") center no-repeat #efeff4;
}
.loginInner {
  width: 80%;
  margin: 0 auto;
  padding-top: 100px;
  background: rgba(0, 0, 0, 0.5);
  height: 400px;
  width: 520px;
  padding-left: 35px;
  padding-right: 35px;
}
.loginInner .login_header .login_logo {
  color: #eee;
  font-weight: bolder;
  font-size: 26px;
  text-align: center;
}
.login_header .login_header_title {
  text-align: center;
  padding-top: 40px;
}
.login_header_title a {
  text-decoration: none;
  font-size: 14px;
  color: #FFFFFF;
  padding-bottom: 4px;
}
.login_header_title a:first-child {
  margin-right: 40px;
}
.login_header_title a.on {
  color: #eee;
  font-weight: bolder;
  border-bottom: 2px solid #eee;
}
.login_content form div {
  display: none;
}
.login_content form div.on {
  display: block;
}
.login_content form input {
  color: white;
  font-size: 16px;
  width: 100%;
  height: 100%;
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  outline: none;
  padding-left: 10px;
  box-sizing: border-box;
  background-color: rgba(0, 0, 0, 0.1);
}
.login_content form input:focus {
  border: 1px solid rgba(0, 0, 0, 0.1);
  color: white;
}
.login_message {
  position: relative;
  margin-top: 20px;
  height: 48px;
  font-size: 14px;
  background: rgba(0, 0, 0, 0.1);
  border-radius: 5px;
}
.login_message .get_verification {
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  border: none;
  color: #ccc;
  background: transparent;
  font-size: 14px;
}
.login_message .get_verification.right_phone {
  color: blue;
}
.login_hint {
  color: #999;
  margin-top: 12px;
  font-size: 14px;
  line-height: 20px;
}
.login_hint a {
  text-decoration: none;
  color: #02a774;
}

.login_verification {
  position: relative;
  margin-top: 20px;
  height: 48px;
  font-size: 14px;
  background: rgba(0, 0, 0, 0.1);
  border-radius: 5px;
}
.login_verification .switch_button {
  border: 1px solid #ddd;
  width: 20px;
  height: 16px;
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  border-radius: 8px;
  padding: 0 6px;
  line-height: 16px;
  font-size: 12px;
  transition: background-color 0.3s;
}
.login_verification .switch_button.on {
  background: #02a774;
}
.login_verification .switch_button.off {
  background: #fff;
}
.switch_button .switch_circle {
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  position: absolute;
  left: -1px;
  top: -1px;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.1);
}
.switch_button .switch_circle.right {
  transform: translateX(20px);
}
.switch_button .switch_text {
  color: #ddd;
  float: right;
}
.login_submit {
  display: block;
  width: 100%;
  height: 42px;
  margin-top: 30px;
  background: #409EFF;
  border-radius: 4px;
  font-size: 16px;
  line-height: 42px;
  color: #fff;
  text-align: center;
  border: none;
}
.registerTip{
  background: rgba(0, 0, 0, 0);
  border: rgba(0, 0, 0, 0);
  margin: 10px 0;
  font-size: 14px;
  color: #409EFF;
  float: right;
}
</style>