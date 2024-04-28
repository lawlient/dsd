import "./App.css"
import { useState } from 'react'
import { Splitter, SplitterPanel } from 'primereact/splitter';
import { InputTextarea } from 'primereact/inputtextarea';
import { Checkbox } from "primereact/checkbox";

import Svg from './components/Svg';
import DownloadButton from './components/DownloadButton';
import ConvertButton from './components/ConvertButton';
import DataTypeSelector from './components/DataTypeSelector';
import OutTypeSelector from './components/OutTypeSelector';


function App() {
    const [req, setReq] = useState({
        type: "bplustree",
        out: "svg",
        data: "",
    })
    const [res, setRes] = useState({
        data:"",
        out: "",
    })

    const [vert, setVert] = useState(false)

    return (
        <>
            <div className="head">
                <DataTypeSelector request={req} setRequest={setReq} />
                <OutTypeSelector request={req} setRequest={setReq} />
                <ConvertButton request={req} setResponse={setRes} />
                <DownloadButton response={res} />
                <Checkbox onChange={e => setVert(e.checked)} checked={vert}></Checkbox>
                <label>{"垂直分屏"}</label>
            </div>
            <Splitter className="content" layout={vert ? 'vertical' : 'horizontal'}>
                <SplitterPanel >
                    <InputTextarea className="input" value={req.data} onChange={(e) => setReq({...req, data:e.target.value})} />
                </SplitterPanel>
                <SplitterPanel >
                    <div className="output">
                        { res.out === "svg" ?  <Svg data={res.data} /> : res.data }
                    </div>
                </SplitterPanel>
            </Splitter>
        </>
    )
}

export default App
