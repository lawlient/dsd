import { Dropdown } from "primereact/dropdown"
import { useState } from "react"


export default function OutTypeSelector({request, setRequest}) {
    const [out, setOut] = useState(null)

    const outs = [
        { name: "SVG", code: "svg", },
        { name: "DOT", code: "dot" },
    ]

    const set = (e) => {
        setOut(e.value)
        setRequest({...request, out:e.value.code})
    }

    return (
        <Dropdown value={out} onChange={set} options={outs} optionLabel="name" placeholder="请选择生成方式" />
    )
}