<template>
  <div>
    <h1>Papers</h1>
    <div>
      <div v-for="paper in papers" :key="paper.ID">
        <a :href="'/paper' + paper.ID">{{ paper.Title }}</a>
        <span>{{ paper.Authors }}</span>
        <span>{{ paper.CreatedAt }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PaperList',
  props: {
    msg: String
  },
  data: () => ({
    papers: []
  }),
  created: function() {
    this.fetchData()
  },
  methods: {
    fetchData: function(){
      fetch('/api/paper')
        .then(x => x.json())
        .then(r => this.papers = r)
        .catch(console.error)
    },
  },
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
