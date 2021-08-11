<template>
  <el-tabs type="card" @tab-click="handleClick">
    <el-tab-pane label="Structure"><structure :sql="sql"/></el-tab-pane>
    <el-tab-pane label="Content"><content /></el-tab-pane>
    <el-tab-pane label="Import"><import /></el-tab-pane>
    <el-tab-pane label="Drop"><drop /></el-tab-pane>
  </el-tabs>
</template>

<script>
import { ref } from "vue";
import Structure from "@/components/Structure";
import Content from "@/components/Content";
import Drop from "@/components/Drop";
import Import from "@/components/Import";

export default {
  name: "Table",
  components: {
    Structure,
    Content,
    Import,
    Drop,
  },
  data() {
    return {
      tabledata: [],
      sql: "",
    };
  },
  computed: {
    getSelectedTable() {
      return this.$store.state.selectedTable;
    },
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
    changeTable() {
      let tables = this.$store.state.tables;
      tables.forEach((item) => {
        if (item.Name == this.$store.state.selectedTable) {
          this.sql = item.SQL;
        }
      });
    },
    // createdatabase() {
    //   this.axios
    //     .post("/table/create?table=" + this.createtablename)
    //     .then((response) => {
    //       console.log(response);
    //       if (response.data === "Success") {
    //         this.createtablename = "";
    //         this.loadData();
    //         //throw success msg
    //       } else {
    //         //throw error msg
    //       }
    //     });
    // },
  },
  watch: {
    getSelectedTable(newValue, oldValue) {
      console.log(oldValue, newValue);
      this.changeTable();
    },
  },
  mounted() {
    this.changeTable();
  },
};
</script>
<style scoped></style>
