<template>
  <title-bar :title-stack="titleStack" />
  <!-- <hero-bar>
    Database Connect
    <template #right>
      <router-link to="/" class="button light">
        Dashboard
      </router-link>
    </template>
  </hero-bar> -->
  <main-section>
    <card-component title="Connect to Sqlite Database" :icon="mdiConnection">
      <field label="Database file location">
        <control>
          <input
            class="input"
            type="text"
            v-model="dblocation"
            placeholder="/home/jason/database.db"
          />
        </control>
      </field>

      <divider />

      <table>
        <thead>
          <tr>
            <td>Connection Status</td>
            <td>{{ status }}</td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Table Count</td>
            <td>{{ tableCount }}</td>
          </tr>
        </tbody>
      </table>

      <divider />

      <field grouped>
        <control>
          <button @click="connectDB" class="button blue">
            Connect
          </button>
        </control>
      </field>
    </card-component>
  </main-section>
</template>

<script>
import { mdiConnection } from "@mdi/js";
import CardComponent from "@/components/CardComponent";
import divider from "@/components/Divider";
import Field from "@/components/Field";
import Control from "@/components/Control";

export default {
  name: "DBConnect",
  components: {
    CardComponent,
    mdiConnection,
    divider,
    Field,
    Control,
  },
  data() {
    return {
      titleStack: [],
      dblocation: "",
      status: "not connected",
      tableCount: 0,
    };
  },
  methods: {
    connectDB() {
      let dbname = this.dblocation;
      console.log("dbname: ", dbname);
      this.axios.post("/dbconnect?db=" + dbname).then((response) => {
        if (response.status === 200) {
          console.log("db connection  ", dbname);
        }
      });
    },
    checkDatabase() {
      this.axios.get("/databaseinfo").then((response) => {
        if (response.status === 200) {
          console.log("db connection info ", response.data);
          if (response.data.db_name) {
            this.dblocation = response.data.db_name;
            this.status = " Connected";
            this.tableCount = response.data.table_count;
          }
        }
      });
    },
  },
  mounted() {
    this.titleStack = ["Admin", "Database Connect"];
    this.checkDatabase();
  },
};
</script>
