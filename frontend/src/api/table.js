import { instance as axios } from "../plugins/axios";

export const table = {
  async dropTable(table) {
    try {
      return await axios.delete("/table/drop/" + table);
    } catch (err) {
      return err;
    }
  },
  async createTable(table) {
    try {
      return await axios.post("/table/create/" + table);
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
};
