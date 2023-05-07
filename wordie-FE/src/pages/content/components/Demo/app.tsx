
import useMount from "react-use/lib/useMount";
import { Hr } from "../../hr/Hr";
export default function App() {
  useMount(async () => {
    new Hr(document.body)
  })

  return <></>;
}
