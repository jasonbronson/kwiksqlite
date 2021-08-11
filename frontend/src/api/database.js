import { instance as axios } from "../plugins/axios";

export const database = {
  async connectDatabase(dbName) {
    try {
      return await axios.post("/dbconnect?db=" + dbName);
    } catch (err) {
      return err;
    }
  },
  async dropTable(dbName) {
    try {
      return await axios.post("/dbconnect?db=" + dbName);
    } catch (err) {
      return err;
    }
  },
};
