import axios from 'axios';

const getAxiosClient = () =>
  axios.create({
    baseURL: process.env.REACT_API_BASE_URL,
  });

const validateNonHtmlResponse = (res: any) => {
  if (res.headers['content-type'].includes('text/html')) {
    throw new Error('Unexpected response. Expected JSON.');
  }

  return res.data;
};

export const getRequestString = async (URL: string): Promise<string> => {
  const res = await getAxiosClient().get<string>(URL, {
    responseType: 'text',
  });

  return validateNonHtmlResponse(res);
};

export const getRequest = async <T>(URL: string): Promise<T> => {
  const res = await getAxiosClient().get<T>(URL);
  return validateNonHtmlResponse(res);
};

export const postRequest = async <T>(URL: string, queryData: any) => {
  const res = await getAxiosClient().post<T>(URL, queryData);
  return validateNonHtmlResponse(res);
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
