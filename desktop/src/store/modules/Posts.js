const state = {
  all: [],
  onePost: {},
};

const mutations = {
  updatePostsList(state, res) {
    state.all = res.data;
  },
  updatePost(state, res) {
    state.onePost = res.data;
  },
  addToPostsList(state, res) {
    state.all.push(res.data);
  },
};

const actions = {
  listPosts(context) {
    window.axios.get("/api/v1/site-admin-posts").then((res) => {
      context.commit("updatePostsList", res);
    });
  },
  getPost(context, id) {
    return window.axios.get("/api/v1/site-admin-posts/" + id).then((res) => {
      context.commit("updatePost", res);
    });
  },
  createPost(context, data) {
    //  let qs = require('qs');
    //  data = qs.stringify(data);

    let config = {
      swalTitle: "Artigo salvo com sucesso",
      icon: "success",
      swalMessage: "Seu novo artigo já está disponível",
    };

    return window.axios
      .post("/api/v1/site-admin-posts", data, config)
      .then((res) => {
        context.commit("addToPostsList", res);
        return res;
      });
  },
  updatePost(context, data) {
    //  let qs = require('qs');
    let id = data.ID;
    //  data = qs.stringify(data);

    let config = {
      swalTitle: "Artigo salvo com sucesso",
      icon: "success",
      swalMessage: "Sua alteração foi publicada",
    };

    return window.axios
      .put("/api/v1/site-admin-posts/" + id, data, config)
      .then((res) => {
        return res;
      });
  },
  removePost(context, id) {
    return window.axios.delete("/api/v1/site-admin-posts/" + id).then((res) => {
      return res;
    });
  },
};

export default {
  state,
  mutations,
  actions,
};
