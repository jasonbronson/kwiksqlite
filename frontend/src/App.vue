<template>
  <el-container style="border: 1px solid #eee">
    <el-aside>
      <sidebar :menu="menu" />
    </el-aside>
    <el-container>
      <el-header>
        <navbar />
      </el-header>
      <el-main>
        <router-view />
      </el-main>
      <div><textarea v-model="logger" name="" id="" cols="120" rows="10">Nothing</textarea></div>
      <footerbar />
    </el-container>
  </el-container>
</template>

<script>
// @ is an alias to /src
import { ref } from "vue";
import { useStore } from "vuex";
import Navbar from "@/components/Navbar";
import Sidebar from "@/components/Sidebar";
import Footerbar from "@/components/Footerbar";

export default {
  name: "Home",
  components: {
    Footerbar,
    Sidebar,
    Navbar,
  },
  data() {
    return {};
  },
  computed: {
    logger(){
      return this.$store.state.logger
    }
  },
  mounted() {
    //If we are missing the db then push back to connect db
    this.polling = setInterval(() => {
      if (this.$store.state.dbName == null) {
        this.$router.push("/");
      }
    }, 3000);
  },
};
</script>
<style>
body {
  margin: 0;
  padding: 0;
  border: 0;
}
.el-aside {
  color: #333;
  width: 250px;
  background-color: #fff;
  padding: 10px 20px;
  overflow: hidden !important;
}
.el-header,
.el-footer {
  background-color: #c6ddfc;
  color: #333;
  text-align: center;
  line-height: 60px;
}
.el-main {
  color: #333;
  text-align: center;
  padding: 10px;
}
</style>
