<template>
  <nav>
    <router-link to="/">Transaction Table</router-link>
    <router-link to="/add">Add Transactions</router-link>
  </nav>
  <router-view />
</template>

<style lang="scss">
* {
  font-family: Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

nav {
  display: flex;
  justify-content: center;

  a {
    display: block;
    box-sizing: border-box;
    padding: 1rem;
    color: black;
    background-color: white;
    text-decoration: none;

    &.router-link-exact-active {
      color: white;
      background-color: black;
    }
  }
}
</style>
