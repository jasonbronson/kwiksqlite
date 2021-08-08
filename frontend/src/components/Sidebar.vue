<template>
  <el-aside class="sidebar">
    <p>KwikSQLite</p>
    <el-menu
      default-active="1"
      class="el-menu"
      @open="handleOpen"
      @close="handleClose"
    >
      <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span>Database</span>
      </el-menu-item>
      <el-menu-item index="2" :disabled="activateMenu">
        <i class="el-icon-menu"></i>
        <span>Query</span>
      </el-menu-item>
      <el-menu-item index="3" :disabled="activateMenu">
        <i class="el-icon-menu"></i>
        <span>Create</span>
      </el-menu-item>
      <el-submenu index="4" :disabled="activateMenu">
        <template #title><i class="el-icon-menu"></i>Tables</template>
        <el-menu-item :index="index" v-for="(table, index) in tables">
          <i class="el-icon-caret-right"></i>
          <span @click="showTable(table.Name)">{{ table.Name }}</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: "HomeView",
  props: {
    menu: {
      required: true,
    },
  },
  data() {
    return {};
  },
  computed: {
    activateMenu() {
      console.log(this.$store.state.dbName, "***");
      if (
        this.$store.state.dbName &&
        this.$store.state.dbName != null &&
        this.$store.state.dbName !== "" &&
        this.$store.state.dbName.length > 0
      ) {
        return false;
      } else {
        return true;
      }
    },
    tables() {
      return this.$store.state.tables;
    },
  },
  methods: {
    showTable(tableName) {
      console.log(tableName);
    },
  },
};
</script>
<style scoped>
.sidebar {
}
</style>
