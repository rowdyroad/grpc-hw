import {useEffect, useState} from "react";
import moment from "moment"
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar, Line } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    BarElement,
    Title,
    Tooltip,
    Legend
);

export default () => {
    const [from, setFrom] = useState(null)
    const [to, setTo] = useState(null)
    const [data, setData] = useState(null)
    const load = () => {
        let q = ""

        if (from) {
            q += "&from="+from + ":00Z"
        }
        if (to) {
            q += "&to="+to+ ":00Z"
        }

        fetch("/api/stats/daily?" + q).then(response=> response.json()).
        then(items=>{
            setData(items)
        })
    }

    useEffect(()=>{
        load()
    }, [])

    return <div>
        <h1>Stats</h1>
        <div className={"row"}>
            <div className={"offset-md-2 col-md-8"}>
                <div className={"row"}>
                    <div className="col">
                        <input type="datetime-local" className="form-control" onChange={e=>setFrom(e.target.value)}/>
                    </div>
                    <div className="col">
                        <input type="datetime-local" className="form-control" onChange={e=>setTo(e.target.value)}/>
                    </div>
                </div>
            </div>
            <div className="col"><button className={"btn btn-primary"} onClick={e=>load()}>Search</button></div>
        </div>
        <div className={"row"}>
            <div className={"col-md-6"}>
                {data && <Bar
                    options={{responsive: true}}
                    data={{
                        labels:Object.keys(data).map(e=>moment(e).format("DD.MM.YYYY")),
                        datasets: [
                            {
                                label: 'Count',
                                data: Object.keys(data).map(t => data[t].count),
                                backgroundColor: 'rgba(53, 162, 235, 0.5)',
                            }
                        ]
                    }}/>}
            </div>
            <div className={"col-md-6"}>
                {data && <Line
                    options={{responsive: true}}
                    data={{
                        labels:Object.keys(data).map(e=>moment(e).format("DD.MM.YYYY")),
                        datasets: [
                            {
                                label: 'Average',
                                data: Object.keys(data).map(t => data[t].average),
                                backgroundColor: 'rgba(153, 162, 235, 0.5)',
                            }
                        ]
                    }}/>}
            </div>
            <div className={"col-md-6"}>
                {data && <Line
                    options={{responsive: true}}
                    data={{
                        labels:Object.keys(data).map(e=>moment(e).format("DD.MM.YYYY")),
                        datasets: [
                            {
                                label: 'Min',
                                data: Object.keys(data).map(t => data[t].min),
                                backgroundColor: 'rgba(153, 162, 135, 0.5)',
                            }
                        ]
                    }}/>}
            </div>
            <div className={"col-md-6"}>
                {data && <Line
                    options={{responsive: true}}
                    data={{
                        labels:Object.keys(data).map(e=>moment(e).format("DD.MM.YYYY")),
                        datasets: [
                            {
                                label: 'Max',
                                data: Object.keys(data).map(t => data[t].max),
                                backgroundColor: 'rgba(123, 132, 125, 0.5)',
                            }
                        ]
                    }}/>}
            </div>
        </div>




    </div>
}