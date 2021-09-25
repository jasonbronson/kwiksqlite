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
  async getTableContent(table, pageOffset, pageLimit) {
    try {
      return await axios.get(
        "/table/content/" +
          table +
          `?pageOffset=${pageOffset}&pageLimit=${pageLimit}`
      );
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
};
