import { instance as axios } from "../plugins/axios";

export const query = {
  async customQuery(query) {
    try {
      let q = { query: query };
      return await axios.post("/query", q);
    } catch (err) {
      console.log(err.response.data);
      console.log(err.response.status);
      console.log(err.response.headers);
      return err.response.data;
    }
  },
};
