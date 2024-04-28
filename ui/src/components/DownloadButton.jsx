import { Button } from "primereact/button";

export default function DownloadButton({response}) {
    const download = () => {
        const url = `data:image/svg+xml;base64,${window.btoa(response.data)}`

        // 创建一个a标签用于下载
        const downloadLink = document.createElement('a');
        downloadLink.href = url;
        downloadLink.download = "dsd.svg";

        // 触发下载
        document.body.appendChild(downloadLink);
        downloadLink.click();
        document.body.removeChild(downloadLink);
    }
    return <Button color={"success"} onClick={download} >Download</Button>
}