import { RefreshCcwIcon } from "lucide-react";

import { Button } from "./button";

export function ButtonWithIcon({ children }: { children: React.ReactNode }) {
  return (
    <Button>
      <RefreshCcwIcon /> {children}
    </Button>
  );
}
