import { useRouter } from "next/dist/client/router";
import { useEffect, useState } from "react";

export const useError = (redirectPath = "/") => {
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();
  const { query } = router;

  useEffect(() => {
    if (query == null || typeof query.auth_error !== "string") return;

    console.error("Error:", query.auth_error);
    router.replace(redirectPath);
    setError(query.auth_error);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query]);

  return { error };
};
