<template lang="pug">
  #app
    #nav
      router-link.mr-1(to="/")
        | Crear Árbol
      router-link(to="/view")
        | Ver Árbol
    router-view
</template>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
.mr-1 {
  margin-right: 10px
}
</style>
