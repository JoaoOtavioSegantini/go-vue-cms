<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-1 nav">
        <p class="display-6">
          <small>AÇÕES</small>
        </p>

        <a href="/posts/add" class="btn btn-secondary" title="Novo artigo">
          <i class="fa fa-plus"></i>
        </a>
      </div>
      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="title">
              <i class="fa fa-file-text-o"></i> Artigos
              <small>Gerenciamento de artigos do blog</small>
            </h5>

            <table
              class="table"
              aria-label="uma tabela contendo uma lista de posts"
            >
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">title</th>
                  <th scope="col">body</th>
                  <th scope="col"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="post in posts" v-bind:key="post">
                  <th scope="row">{{ post.ID }}</th>
                  <td>{{ post.title }}</td>
                  <td>
                    <div class="cut-text">{{ post.body }}</div>
                  </td>
                  <td class="text-right">
                    <a
                      :href="'/posts/' + post.ID"
                      class="btn btn-primary btn-sm"
                      >ver</a
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    posts() {
      return this.$store.state.Posts.all;
    }
  },
  mounted() {
    this.$store.dispatch("listPosts");
  },
};
</script>

<style>
.cut-text {
  max-width: 300px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
</style>
