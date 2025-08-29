import { useQuery } from '@tanstack/react-query';
import { getRequestString } from './api';

export const useFetchCatchphraseString = () => {
  const { data, error, isError, isLoading } = useQuery<string>({
    queryKey: ['catchphraseString'],
    queryFn: () => getRequestString('catchphrase/string'),
  });

  return { catchphraseString: data || '', error, isError, isLoading };
};
