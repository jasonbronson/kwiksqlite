<template>
  <hero-bar>
    <div class="createdb">
      <div class="title">Create Table</div>
      <input type="text" placeholder="tablename" v-model="createtablename" />
      <button @click="createdatabase()" class="button blue medium">
        CREATE
      </button>
    </div>

    <template #right> </template>
  </hero-bar>
  <main-section>
    <card-component title="Tables" :icon="mdiTable" has-table>
      <db-table :tabledata="tabledata" @reload="loadData" />
    </card-component>
  </main-section>
</template>

<script>
import { ref } from "vue";

export default {
  name: "Tables",
  components: {},
  data() {
    return {
      tabledata: [],
      createtablename: "",
    };
  },
  methods: {
    createdatabase() {
      this.axios
        .post("/table/create?table=" + this.createtablename)
        .then((response) => {
          console.log(response);
          if (response.data === "Success") {
            this.createtablename = "";
            this.loadData();
            //throw success msg
          } else {
            //throw error msg
          }
        });
    },
  },
  mounted() {},
};
</script>
<style scoped>
input {
  border: 1px solid black;
  font-size: 75%;
  width: 35%;
}
.createdb {
  display: flex;
  align-items: center;
}
.createdb .title {
  padding-right: 10px;
  font-size: 0.8em;
}
.button {
  border-radius: 0;
  font-size: 13px;
  padding: 4px;
  padding-left: 10px;
  padding-right: 10px;
}
</style>
