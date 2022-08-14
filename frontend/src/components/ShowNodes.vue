<template>
  <div id="show_nodes" class="show_nodes">
    <h2>{{ msg }}</h2>

<!--    <p>{{ num1 }}</p>-->
<!--    <p>{{ num2 }}</p>-->
<!--    <p>{{ add }}</p>-->
<!--    <p>{{ mul }}</p>-->
<!--    <p>{{ sub }}</p>-->
<!--    <p>{{ div }}</p>-->

<!--    <input name="num1" v-model="num1">-->
<!--    <input name="num2" v-model="num2">-->

<!--    <button v-on:click="postreq">计算</button>-->

<!--    <input name="id" v-model="id">-->
<!--    <button v-on:click="getNode">获取节点信息</button>-->
<!--    <p>{{ nodeInfo }}</p>-->

    <NodeTables></NodeTables>
  </div>
</template>

<script>


import axios from "axios";
import NodeTables from "@/components/NodeTables";

export default {
  name: 'ShowNodes',
  components: {NodeTables},
  props: {
    msg: String
  },
  data() {
    return {
      add: "", mul: "", sub: "", div: "",
      num1: "", num2: "", id: "", nodeInfo: "",
    }
  },
  methods: {
    postreq: function () {
      var data = {"num1": parseFloat(this.num1), "num2": parseFloat(this.num2)}
      /*eslint-disable*/
      console.log(data)
      /*eslint-enable*/
      axios({
        method: "POST",
      url: "/api/v1/calc",
        data: data,
        headers: {"content-type": "text/plain"}
      }).then(result => {
        // this.response = result.data;
        this.add = result.data['add']
        this.mul = result.data['mul']
        this.sub = result.data['sub']
        this.div = result.data['div']
        /*eslint-disable*/
        console.log(result.data)
        /*eslint-enable*/
      }).catch(error => {
        /*eslint-disable*/
        console.error(error);
        /*eslint-enable*/
      });
    },
    getNode: function () {
      var data = {"id": this.id}
      /*eslint-disable*/
      console.log(data)
      /*eslint-enable*/
      axios({
        method: "POST",
        url: "/api/v1/getNode",
        data: data,
        headers: {"content-type": "text/plain"}
      }).then(result => {
        // this.response = result.data;
        this.nodeInfo = result.data['nodeInfo']
        /*eslint-disable*/
        console.log(result.data)
        /*eslint-enable*/
      }).catch(error => {
        /*eslint-disable*/
        console.error(error);
        /*eslint-enable*/
      });
    }
  }
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
