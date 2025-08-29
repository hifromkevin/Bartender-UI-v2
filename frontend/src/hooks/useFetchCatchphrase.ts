import { useQuery } from '@tanstack/react-query';
import { getRequest } from './api';

export const useFetchCatchphrase = () => {
  const { data, error, isError, isLoading } = useQuery<
    {
      catchphrase: string;
      byline: string;
    }[]
  >({
    queryKey: ['catchphrase'],
    queryFn: () => getRequest('catchphrase'),
  });

  return {
    catchphrase: data || [],
    error,
    isError,
    isLoading,
  };
};
