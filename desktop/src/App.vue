<template>
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css"
    integrity="sha512-9usAa10IRO0HhonpyAIVpjrylPvoDwiPUiKdWk5t3PyolY1cOd4DSE0Ga+ri4AuTroPR5aQvXU9xC6qOPnzFeg=="
    crossOrigin="anonymous"
    referrerPolicy="no-referrer"
  />
  <div id="app">
    <loader></loader>
    <header id="header">
      <nav class="navbar navbar-light bg-light">
        <span class="navbar-brand h1"
          ><i class="fa fa-space-shuttle"></i> CMS Electron</span
        >
        <div class="btn-group btn-group-sm">
          <button class="btn btn-secondary" @click="minimize()">
            <i class="fa fa-window-minimize"></i>
          </button>
          <button class="btn btn-secondary" @click="maximize()">
            <i class="fa fa-window-maximize"></i>
          </button>
          <button class="btn btn-secondary" @click="fullScreen()">
            <i class="fa fa-arrows-alt"></i>
          </button>
          <button class="btn btn-dark" @click="close()">
            <i class="fa fa-window-close"></i>
          </button>
        </div>
      </nav>
    </header>

    <div class="container-fluid" id="content">
      <div class="row">
        <div id="nav" class="col-1 bg-dark">
          <p class="display-6 text-light">
            <small>MENU</small>
          </p>

          <p>
            <router-link to="/pages" class="btn btn-secondary"
              ><i class="fa fa-file-text"></i
            ></router-link>
          </p>

          <p>
            <router-link to="/posts" class="btn btn-secondary"
              ><i class="fa fa-rss"></i
            ></router-link>
          </p>

          <p>
            <router-link to="/users" class="btn btn-secondary"
              ><i class="fa fa-user"></i
            ></router-link>
          </p>

          <p>
            <a href="" class="btn btn-secondary" @click.prevent="logout()"
              ><i class="fa fa-sign-out"></i
            ></a>
          </p>
        </div>
        <div class="col-11">
          <router-view v-slot="{ Component }">
            <transition
              enter-active-class="animated animate__animated animate__fadeIn"
              leave-active-class="animated animate__animated animate__zoomOut"
            >
              <component :is="Component" />
            </transition>
          </router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const electron = require("electron");

import loader from "./components/Loader";
export default {
  name: "App",
  components: {
    loader: loader,
  },
  data() {
    return {
      win: null,
    };
  },
  methods: {
    minimize() {
      this.win.minimize();
    },
    maximize() {
      if (this.win.isMaximized()) {
        this.win.unmaximize();
      } else {
        this.win.maximize();
      }
    },
    fullScreen() {
      this.win.setFullScreen(!this.win.isFullScreen());
    },
    close() {
      this.win.close();
    },
    logout() {
      window.axios.defaults.headers.common["Authorization"] = null;
      window.localStorage.removeItem("token");
      this.$router.push({ path: "/auth" });
    },
  },
  mounted() {
    this.win = electron.remote.BrowserWindow.getFocusedWindow();
  },
};
</script>

<style>
html,
body,
#app {
  height: 100%;
}
#app {
  display: flex;
  flex-direction: column;
}
#content {
  flex: 1;
}
#content > .row {
  height: 100%;
}
body {
  background-color: rgba(236, 240, 241, 1);
}
#header {
  margin-bottom: 10px;
  -webkit-app-region: drag;
}
#header .btn {
  -webkit-app-region: no-drag;
}
#content {
  position: relative;
}
#content .animated {
  position: absolute;
  top: 0;
}
#nav,
.nav {
  padding-top: 15px;
  margin-top: -10px;
}
.nav {
  display: block !important;
}
.router-link-active {
  background-color: #dc3545 !important;
  border-color: #dc3545 !important;
}
</style>
