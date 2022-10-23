const state = {
  all: [],
  onePage: {},
};

const mutations = {
  updatePagesList(state, res) {
    state.all = res.data;
  },
  updatePage(state, res) {
    state.onePage = res.data;
  },
  addToPagesList(state, res) {
    state.all.push(res.data);
  },
};

const actions = {
  listPages(context) {
    window.axios.get("/api/v1/res-data").then((res) => {
      context.commit("updatePagesList", res);
    });
  },
  getPage(context, id) {
    return window.axios.get("/api/v1/res-data/" + id).then((res) => {
      context.commit("updatePage", res);
    });
  },
  createPage(context, data) {
    // let qs = require("qs");

    let config = {
      swalTitle: "Página salva com sucesso",
      swalMessage: "Sua nova página já está disponível",
    };

    return window.axios.post("/api/v1/res-data", data, config).then((res) => {
      context.commit("addToPagesList", res);
      return res;
    });
  },
  updatePage(context, data) {
    let id = data.ID;

    let config = {
      swalTitle: "Página salva com sucesso",
      swalMessage: "Sua nova página já está disponível",
    };

    return window.axios
      .put("/api/v1/res-data/" + id, data, config)
      .then((res) => {
        return res;
      });
  },
  removePage(context, id) {
    return window.axios.delete("/api/v1/res-data/" + id).then((res) => {
      return res;
    });
  },
};

export default {
  state,
  mutations,
  actions,
};
