<template>
    <div class="container-fluid">
        <div class="row">
          <div class="col-1 nav">
            <p class="display-6">
              <small>AÇÕES</small>
            </p>
  
            <a :href="'/posts/' + post.ID" class="btn btn-secondary" title="Voltar">
              <i class="fa fa-arrow-circle-left"></i>
            </a>
          </div>
          <div class="col-11">
            <div class="card">
              <div class="card-body">
                <h5 class="title"><i class="fa fa-file-text-o"></i> {{ post.title }} <small>Remoção de artigo do blog</small></h5>
  
                <div class="alert alert-danger">
                  Tem certeza que quer remover este artigo, essa ação não poderá ser desfeita!
                </div>
  
              </div>
              <div class="card-footer">
                <a :href="'/posts/' + post.ID" class="btn btn-secondary">Não remover</a>
                <a href="" class="btn btn-primary" @click.prevent="remove()">Apagar definitivamente</a>
              </div>
            </div>
          </div>
        </div>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      remove() {
        this.$store.dispatch('removePost', this.post.ID)
          .then(() => {
            this.$router.push({path: '/posts'});
          })
      }
    },
    computed: {
      post() {
        return this.$store.state.Posts.onePost
      }
    },
    mounted () {
      this.$store.dispatch('getPost', this.$route.params.id)
    }
  }
  </script>