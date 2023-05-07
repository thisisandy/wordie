import { useState } from "react";
import { Request } from "../network/request";
export const useRequest = () => {
  const [request]  = useState(() => new Request());
  return [request]
}
