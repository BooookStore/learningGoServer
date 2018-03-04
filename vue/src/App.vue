<template>
  <div id="app">
    <el-header>
      <el-menu :default-active="'1'" mode="horizontal">
        <el-menu-item index="1"><router-link to="/user">Users</router-link></el-menu-item>
        <el-menu-item index="2">Animal</el-menu-item>
        <el-menu-item index="3">Area</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>
      <router-view/>
    </el-main>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
