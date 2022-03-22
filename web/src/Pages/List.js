import {useEffect, useState} from "react";
import Moment from 'react-moment';
import "moment-timezone"

export default () => {
    const [data, setData] = useState([])
    const [offset, setOffset] = useState(0)
    const [from, setFrom] = useState(null)
    const [to, setTo] = useState(null)
    const [low, setLow] = useState(null)
    const [high, setHigh] = useState(null)

    const load = (clear) => {
        let q = "offset=" + (clear ? 0 : offset)

        if (from) {
            q += "&from="+from + ":00Z"
        }
        if (to) {
            q += "&to="+to+ ":00Z"
        }
        if (low) {
            q += "&low="+low
        }
        if (high) {
            q += "&high="+high
        }
        fetch("/api/records?" + q).then(response=> response.json()).
        then(items=>{
            setData(clear ? items : [...data, ...items])
        })
    }
    useEffect(()=>{
        setOffset(data.length)
    }, [data])

    useEffect(()=>{
        load()
    },[])

    return <div>
        <h1>List</h1>
        <div className={"row"}>
            <div className={"offset-md-2 col-md-8"}>
                <div className={"row"}>
                    <div className="col">
                        <input type="datetime-local" className="form-control" onChange={e=>setFrom(e.target.value)}/>
                    </div>
                    <div className="col">
                        <input type="datetime-local" className="form-control" onChange={e=>setTo(e.target.value)}/>
                    </div>
                    <div className="col">
                        <input type="number" className="form-control" placeholder="Value from" onChange={e=>setLow(e.target.value)}/>
                    </div>
                    <div className="col">
                        <input type="number" className="form-control" placeholder="Value to" onChange={e=>setHigh(e.target.value)}/>
                    </div>
                    <div className="col"><button className={"btn btn-primary"} onClick={e=>load(true)}>Search</button></div>
                </div>

            </div>
        </div>
        <br/>
        {from} {to} {low} {high}
        <table className={"table"}>
            <thead>
                <tr><th>Time</th><th>Value</th></tr>
            </thead>
            {data.map(e=><tr><td><Moment utc format={"DD.MM.YYYY HH:mm:ss"}>{e.time}</Moment></td><td>{e.value}</td></tr>)}
        </table>
        <button className={"btn btn-primary"} onClick={e=>load()}>Load more</button>
    </div>
}