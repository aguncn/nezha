<template>
  <div id="app">
    <router-view />
    <div class="foot">
      哪吒NeZha@2020
    </div>
  </div>
</template>
<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.foot {
  height: 50px;
  padding-top: 15px;
  background: #ddd;
  margin: 0px auto 30px auto;
  text-align: center;
  color: #505458;
}
</style>

<script>
export default {
  name: 'App'
}
</script>
