import { useQuery } from '@tanstack/react-query';
import { getRequest } from './api';

export const useFetchCatchphrase = () => {
  const { data, error, isError, isLoading } = useQuery<
    { catchphrase: string },
    {
      message: string;
      status: number;
    }
  >({
    queryKey: ['catchphrase'],
    queryFn: () => getRequest<{ catchphrase: string }>('catchphrase'),
  });

  return { catchphrase: data?.catchphrase || '', error, isError, isLoading };
};
