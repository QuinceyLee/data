<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-box-left">
        <div class="login-title">RULEX</div>
      </div>
      <div class="login-box-right">
        <canvas id="myCanvas" width="200" height="200" />
        <el-form
          ref="loginFormRef"
          :model="loginForm"
          :rules="loginFormRules"
          class="login-form"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginForm.username"
              type="text"
              maxlength="10"
              placeholder="Username"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              type="password"
              maxlength="16"
              placeholder="Password"
              @keyup.enter.native="handleLogin"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              :loading="loading"
              class="login-btn"
              @click.native.prevent="handleLogin"
            >
              登录
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <ul class="w3lsg-bubbles">
      <li />
      <li />
      <li />
      <li />
      <li />
      <li />
      <li />
      <li />
      <li />
      <li />
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginFormRules: {
        username: [
          {
            required: true,
            trigger: 'blur',
            message: 'Please enter user name'
          },
          {
            min: 5,
            max: 10,
            trigger: 'blur',
            message: 'Please enter the correct user name'
          }
        ],
        password: [
          {
            required: true,
            trigger: 'blur',
            message: 'Please enter password'
          },
          {
            min: 5,
            max: 16,
            trigger: 'blur',
            message: 'The password can not be less than 6 digits'
          }
        ]
      },
      loading: false,
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

  mounted() {
    this.initLogo()
  },

  methods: {
    initLogo() {
      var c = document.getElementById('myCanvas')
      var ctx = c.getContext('2d')
      ctx.lineWidth = 0
      ctx.beginPath()
      ctx.arc(100, 100, 50, 0, 2 * Math.PI)
      ctx.strokeStyle = '#00BFFF'
      ctx.stroke()
      ctx.beginPath()
      ctx.arc(100, 100, 20, 0, 2 * Math.PI)
      ctx.fillStyle = '#00BFFF'
      ctx.fill()
      ctx.strokeStyle = '#00BFFF'
      ctx.stroke()
      ctx.beginPath()
      ctx.arc(70, 60, 10, 0, 2 * Math.PI)
      ctx.fillStyle = '#00BFFF'
      ctx.fill()
      ctx.strokeStyle = '#00BFFF'
      ctx.stroke()
    },

    handleLogin() {
      this.$refs.loginFormRef.validate((valid) => {
        if (valid) {
          this.loading = true
          this.$store
            .dispatch('user/login', this.loginForm)
            .then(() => {
              this.$router.push({ path: this.redirect || '/' })
              this.loading = false
            })
            .catch(() => {
              this.loading = false
            })
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login-container {
  height: 100%;
  position: relative;
  overflow: hidden;
  .login-box {
    width: 900px;
    height: 600px;
    display: flex;
    position: absolute;
    top: 50%;
    left: 50%;
    z-index: 7;
    transform: translate(-50%, -50%);
    background: #fff;
    box-shadow: 0px 0px 5px 5px rgba(0, 0, 0, 0.2);
    .login-box-left {
      width: 55%;
      display: flex;
      align-items: center;
      justify-content: center;
      background-image: url(./../../assets/login_bg.jpg);
      background-size: 100% 100%;
      .login-title {
        padding: 15px 40px;
        color: #00a1e7;
        font-size: 60px;
        line-height: 1;
        font-weight: bold;
        background: #fff;
      }
    }
    .login-box-right {
      width: 45%;
      display: flex;
      flex-direction: column;
      align-items: center;
      padding: 40px;
      .login-form {
        width: 100%;
        /deep/ .el-input__inner {
          border: none;
          border-radius: 0;
          border-bottom: 3px solid #409eff;
        }
        .login-btn {
          width: 100%;
          margin-top: 20px;
        }
      }
    }
  }
  .w3lsg-bubbles {
    width: 100%;
    height: 100%;
    position: absolute;
    z-index: -1;
    li {
      width: 40px;
      height: 40px;
      position: absolute;
      bottom: -160px;
      list-style: none;
      text-decoration: none;
      font-size: 0;
      border-radius: 5px;
      background-color: rgba(64, 158, 255, 0.5);
      animation: square 13s infinite;
      transition-timing-function: linear;
    }
  }
  .w3lsg-bubbles li:nth-child(1) {
    left: 5%;
  }
  .w3lsg-bubbles li:nth-child(2) {
    left: 20%;
    width: 80px;
    height: 80px;
    animation-delay: 6s;
    background-color: rgba(64, 158, 255, 0.2);
  }
  .w3lsg-bubbles li:nth-child(3) {
    left: 10%;
    animation-delay: 10s;
    background-color: rgba(64, 158, 255, 0.3);
  }
  .w3lsg-bubbles li:nth-child(4) {
    left: 90%;
    width: 60px;
    height: 60px;
    animation-delay: 8s;
    background-color: rgba(64, 158, 255, 0.4);
  }
  .w3lsg-bubbles li:nth-child(5) {
    left: 85%;
    animation-delay: 6s;
  }
  .w3lsg-bubbles li:nth-child(6) {
    width: 120px;
    height: 120px;
    left: 82%;
    animation-delay: 5s;
    background-color: rgba(64, 158, 255, 0.6);
  }
  .w3lsg-bubbles li:nth-child(7) {
    left: 10%;
    width: 160px;
    height: 160px;
    animation-delay: 2s;
  }
  .w3lsg-bubbles li:nth-child(8) {
    left: 75%;
    width: 20px;
    height: 20px;
    animation-delay: 2s;
  }
  .w3lsg-bubbles li:nth-child(9) {
    left: 25%;
    width: 10px;
    height: 10px;
    animation-delay: 3s;
    background-color: rgba(64, 158, 255, 0.6);
  }
  .w3lsg-bubbles li:nth-child(10) {
    left: 90%;
    width: 160px;
    height: 160px;
    background-color: rgba(64, 158, 255, 0.4);
  }
  @keyframes square {
    0% {
      transform: translateY(0);
    }
    100% {
      transform: translateY(-1200px) rotate(600deg);
    }
  }
}
</style>
