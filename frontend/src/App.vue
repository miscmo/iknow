<template>
<!--  <img alt="Vue logo" src="./assets/logo.png">-->
<!--  <ShowNodes msg="Iknow"/>-->
  <IKnowMain></IKnowMain>
<!--  <Iknow msg="Welcome to Your Vue.js App"/>-->
</template>

<script>
// import ShowNodes from './components/ShowNodes.vue'
import IKnowMain from "@/components/IKnowMain";
// import Iknow from './components/Iknow.vue'

export default {
  name: 'App',
  components: {
    IKnowMain,
     // Iknow,
    // ShowNodes,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
