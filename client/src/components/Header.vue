<template>
  <header>
    <div class="left">
      <router-link to="/">論文發射器</router-link>
      <router-link to="/about">上傳論文</router-link>
    </div>
    <div class="right">
      <router-link to="/login">登入</router-link>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style lang='less' scoped>

header{
  display: flex;
  justify-content: space-between;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}
#nav a.router-link-exact-active {
  color: #42b983;
}

</style>
