export default function Svg({data}) {
  return (
    <div dangerouslySetInnerHTML={{
      __html: data
    }}/>
  )
}
