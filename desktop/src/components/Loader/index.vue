<template>
  <transition
    enter-active-class="animate__animated animate__fadeIn"
    leave-active-class="animate__animated animate__fadeOut"
  >
    <div id="preloader" v-show="counter > 0">
      <div class="preloader-box">
        <div class="item item-1"></div>
        <div class="item item-2"></div>
        <div class="item item-3"></div>
        <div class="item item-4"></div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data() {
    return {
      counter: 0,
    };
  },
  mounted() {
    let counter = this.counter;
    console.log(counter);

    window.axios.interceptors.request.use(
      function (config) {
        // Do something before request is sent
        counter++;

        return config;
      },
      function (error) {
        // Do something with request error
        return Promise.reject(error);
      }
    );

    // Add a response interceptor
    window.axios.interceptors.response.use(
      function (response) {
        // Any status code that lie within the range of 2xx cause this function to trigger
        // Do something with response data

        counter--;

        return response;
      },
      function (error) {
        counter--;

        // Any status codes that falls outside the range of 2xx cause this function to trigger
        // Do something with response error
        return Promise.reject(error);
      }
    );
  },
};
</script>

<style>
#preloader {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.7);
  z-index: 999;
}
#preloader .preloader-box {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 100px;
  height: 100px;
  transform: translate(-50%, -50%);
  padding: 15px 30px;
  overflow: hidden;
}
#preloader .preloader-box .item {
  width: 50px;
  height: 50px;
  position: absolute;
}
#preloader .preloader-box .item-1 {
  background-color: #fa5667;
  top: 0;
  left: 0;
  z-index: 1;
  animation: item-1_move 1.8s cubic-bezier(0.6, 0.01, 0.4, 1) infinite;
}
#preloader .preloader-box .item-2 {
  background-color: #7a45e5;
  top: 0;
  right: 0;
  animation: item-2_move 1.8s cubic-bezier(0.6, 0.01, 0.4, 1) infinite;
}
#preloader .preloader-box .item-3 {
  background-color: #1b91f7;
  bottom: 0;
  right: 0;
  animation: item-3_move 1.8s cubic-bezier(0.6, 0.01, 0.4, 1) infinite;
}
#preloader .preloader-box .item-4 {
  background-color: #fac24c;
  bottom: 0;
  left: 0;
  animation: item-4_move 1.8s cubic-bezier(0.6, 0.01, 0.4, 1) infinite;
}
@keyframes item-1_move {
  0%,
  100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(0, 50px);
  }
  50% {
    transform: translate(50px, 50px);
  }
  75% {
    transform: translate(50px, 0);
  }
}
@keyframes item-2_move {
  0%,
  100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(-50px, 0);
  }
  50% {
    transform: translate(-50px, 50px);
  }
  75% {
    transform: translate(0, 50px);
  }
}
@keyframes item-3_move {
  0%,
  100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(0, -50px);
  }
  50% {
    transform: translate(-50px, -50px);
  }
  75% {
    transform: translate(-50px, 0);
  }
}
@keyframes item-4_move {
  0%,
  100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(50px, 0);
  }
  50% {
    transform: translate(50px, -50px);
  }
  75% {
    transform: translate(0, -50px);
  }
}
</style>
