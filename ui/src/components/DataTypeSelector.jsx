import { Dropdown } from "primereact/dropdown";
import { useState } from "react";


export default function DataTypeSelector({request, setRequest}) {
    const [datatype, setDatatype] = useState(null)

    const types = [
        { name: "B树" , code: "btree" },
        { name: "B+树" , code: "bplustree" },
    ]

    const set = (e) => {
        setDatatype(e.value)
        setRequest({...request, type:e.value.code})
    }

    return (
        <Dropdown value={datatype} onChange={set} options={types} optionLabel="name" placeholder="请选择数据结构类型" />
    )
}