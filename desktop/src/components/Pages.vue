<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-1 nav">
        <p class="display-6">
          <small>AÇÕES</small>
        </p>

        <a href="/pages/add" class="btn btn-secondary" title="Nova página">
          <i class="fa fa-plus"></i>
        </a>
      </div>
      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="title">
              <i class="fa fa-file-text-o"></i> Páginas
              <small>Gerenciamento de páginas do site</small>
            </h5>

            <table
              class="table"
              aria-label="tabela contendo uma lista de páginas"
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
                <tr v-for="page in pages" v-bind:key="page">
                  <th scope="row">{{ page.ID }}</th>
                  <td>{{ page.title }}</td>
                  <td>
                    <div class="cut-text">{{ page.body }}</div>
                  </td>
                  <td class="text-right">
                    <a
                      :href="'/pages/' + page.ID"
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
    pages() {
      return this.$store.state.Pages.all;
    }
  },
  mounted() {
    this.$store.dispatch("listPages");
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
