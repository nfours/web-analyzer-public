import axios from "axios";

export const analyzePage = async (url) => {
  const response = await axios.post("/api/analyze", { url });

 return response.data;
};
