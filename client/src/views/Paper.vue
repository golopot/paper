<template>
  <div>
    <div>{{ title }}</div>
    <div>{{ authors }}</div>
    <div>
      <a class="download" :href="&quot;/file/&quot; + id">下載論文</a>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Paper',
  data: () => ({
    id: '',
    title: '',
    authors: '',
    link: '',
  }),
  created() {
    this.fetch()
  },
  methods: {
    fetch() {
      fetch('/api/paper/' + this.$route.params.id)
        .then( r => r.json() )
        .then( r => {
          this.id = r.ID
          this.title = r.Title
          this.authors = r.Authors
        })
        .catch(console.error)
    }
  },

}

</script>

<style lang='less' scoped>
.download {
  text-decoration: none;
  color: white;
  display: inline-block;
  background: #4367B3;
  padding: 6px 12px;
}
</style>
