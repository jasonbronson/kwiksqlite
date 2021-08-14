<template>
  <el-tabs class="createPane" type="card" @tab-click="handleClick">
    <el-tab-pane label="Create">
      <el-form ref="form" label-width="120px">
        <el-form-item label="">
          <el-form-item label="Table name">
            <el-input v-model="tableName"></el-input>
          </el-form-item>
        </el-form-item>
        <el-form-item>
          <el-col :span="11">
            <el-button type="primary" @click="createTable"
              >Create Table</el-button
            >
          </el-col>
        </el-form-item>
      </el-form>
      <h4>Naming Conventions in SQLite</h4>
      <p>
        A naming convention is a set of rules for choosing the character
        sequence to be used for identifiers that denote the name of a database,
        table, column, index, trigger, or view in SQLite.
      </p>
      <p>
        Valid Characters An identifier name must begin with a letter or the
        underscore character, which is followed by any alphanumeric character or
        underscore.
      </p>
      <p>
        Other characters are not valid. The following are examples of valid
        identifiers. 
        <ul class="conventionRules">
          <li>tablemaster </li>
          <li>table_master </li>
          <li>table1 </li>
          <li>_Table1 </li>
        </ul>
        <p>The following are the example of some invalid identifiers. </p>
        <ul class="conventionRules">
          <li>table master</li>
          <li>table-master</li>
          <li>1table</li>
        </ul>
        
      </p>
    </el-tab-pane>
    <el-tab-pane label="Import">Not Completed Yet</el-tab-pane>
  </el-tabs>
</template>
<script>
import Vue from "vue";
import api from "@/api";

export default {
  name: "Drop",
  data() {
    return {
      tableName: "",
    };
  },
  methods: {
    handleClick() {},
    createTable() {
      api.table.createTable(this.tableName).then((response) => {
        let success = false;
        if (response.status === 200) {
          this.$store.commit("setTables", response.data.tables);
          this.$store.commit("setTableCount", response.data.table_count);
          this.$store.commit("setTable", this.tableName);
          success = true;
          this.$notify({
            title: "Success",
            message: "Create table succeeded",
          });
          this.$router.push("/table/" + this.tableName);
        }
        if (!success) {
          console.log(response);
          this.$notify.error({
            title: "Error",
            message: response,
          });
        }
      });
    },
  },
  computed: {},
};
</script>
<style scoped>
.createPane {
  text-align: left;
}
.conventionRules{
  list-style: none;
}
.conventionRules li{
  padding-top: 15px;
}

</style>
