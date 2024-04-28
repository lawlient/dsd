import { Button } from "primereact/button";
import { getdotpng } from "../api/api";


export default function ConvertButton({request, setResponse}) {
    const convert = () => {
        getdotpng(request).then(res => {
            console.log(res.data)
            if (request.out === "dot") {
                setResponse({data: res.data.data, out:"dot"})
            } else {
                setResponse({data:res.data, out:"svg"})
            }
        })
    }

    return <Button color={"primary"} onClick={convert} >Convert</Button>
}