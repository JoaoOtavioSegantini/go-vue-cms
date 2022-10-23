import axios from "axios";
import swal from "sweetalert2";
import router from "./router";

if (window.localStorage.getItem("token")) {
  axios.defaults.headers.common["Authorization"] =
    window.localStorage.getItem("token");
}

axios.defaults.baseURL = process.env.VUE_APP_BACKEND_URL
axios.defaults.headers.post['Content-Type'] = "application/json";

axios.interceptors.request.use(
  function (config) {
    let defaultTitle = "Salvo com sucesso";
    let defaultMessage = "O registro foi salvo com sucesso";

    if (config.method === "delete") {
      defaultTitle = "Removido com sucesso";
      defaultMessage = "O registro foi removido com sucesso";
    }

    if (config.method !== "get") {
      new swal({
        title: config.swalTitle || defaultTitle,
        text: config.swalMessage || defaultMessage,
        icon: "success",
        confirmButtonText: "Ok!",
      });
    }

    return config;
  },
  function (err) {
    return Promise.reject(err);
  }
);

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  function (err) {
    let title = "Algo deu errado";
    let text =
      "Uma situação inesperada ocorreu no servidor, por favor, entre em contato com o administrador";
    let btnLabel = "Eu entendo";
    let closeCallback = (result) => result;

    if (err.response && err.response.status == 401) {
      title = "Autenticação!";
      text = "Você precisa estar logado para acessar este recurso";
      btnLabel = "Autenticar";
      // eslint-disable-next-line no-unused-vars
      closeCallback = (_result) => router.push({ path: "/auth" });
    }

    if (err.response && err.response.status == 422) {
      title = "Verifique os dados";
      text = "O servidor recusou os dados informado, confira as informações";
    }

    new swal({
      title: title,
      text: text,
      icon: "error",
      confirmButtonText: btnLabel,
    }).then(closeCallback);

    return Promise.reject(err);
  }
);

window.axios = axios;
