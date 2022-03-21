import {useEffect, useState} from "react";

export default () => {
    const [time, setTime] = useState(null)
    const [data, setData] = useState(null)

    useEffect(()=> {
        fetch("/api/events/" + time + ":00Z").then(response=> response.json()).
        then(d=>{
            setData(d)
        })
    }, [time])
    return <div>

        <h1>Find</h1>
        <div className={"row"}>
            <div className={"offset-md-2 col-md-8"}>
                <div className={"row"}>
                    <div className="col">
                        <input type="datetime-local" className="form-control" onChange={e=>setTime(e.target.value)}/>
                    </div>
                </div>
            </div>
        </div>
        <br/>
        <div>Value</div>
        <h1>{data}</h1>

    </div>
}