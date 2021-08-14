import { instance as axios } from "../plugins/axios";

export const structure = {
  async addColumnTable(table, column) {
    try {
      let data = { column: column };
      return await axios.post(`/table/${table}/column/create/`, data);
    } catch (err) {
      return err;
    }
  },
  async dropColumnTable(table, column) {
    try {
      return await axios.delete(`/table/${table}/column/drop/${column}`);
    } catch (err) {
      return err;
    }
  },
  async getColumns(table) {
    try {
      return await axios.get(`/table/columns/${table}`);
    } catch (err) {
      return err;
    }
  },
};
