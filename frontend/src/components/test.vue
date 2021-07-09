<template>
  <div class="hello">
    <h1>test module</h1>
    <button @click="sayHello">Run API</button>
  </div>
</template>

<script>
import Vue from "vue";
export default {
  name: "test",
  data() {
    return {
      data: "",
    };
  },
  methods: {
    sayHello() {
      console.log("call api");
      Vue.axios.get("/").then((response) => (this.data = response.data));
    },
  },
  mount() {
    //this.sayHello();
  },
};
</script>

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
