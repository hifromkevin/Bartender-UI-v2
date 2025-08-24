import axios from 'axios';

const getAxiosClient = () =>
  axios.create({
    baseURL: process.env.REACT_API_BASE_URL,
  });

export const getRequest = async <T>(URL: string): Promise<T> => {
  const { data } = await getAxiosClient().get<T>(URL);
  return data;
};

export const postRequest = async <T>(URL: string, queryData: any) => {
  const { data } = await getAxiosClient().post<T>(URL, queryData);
  return data;
};

export const updateRequest = async (
  URL: string,
  id: string,
  updatedData: any
) => {
  return await getAxiosClient().put(`${URL}/${id}`, updatedData);
};

export const deleteRequest = async (URL: string, id: string): Promise<void> => {
  await getAxiosClient().delete(`${URL}/${id}`);
};
