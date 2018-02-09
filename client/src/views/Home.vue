<template>
  <div class="home">
    <h1>所有論文 <span class="count">513筆</span></h1>
    <Uploader />
    <PaperList />
  </div>
</template>

<script>
import PaperList from '../components/PaperList.vue'
import Uploader from '../components/Uploader.vue'

export default {
  name: 'Home',
  components: {
    PaperList,
    Uploader
  }
}
</script>

<style>
.home {
  padding: 40px 0 40px;
}
.count {
  font-size: 20px;
}
</style>
