import { sleep } from 'k6';
import http from 'k6/http';

export default function () {
    const url = 'http://localhost:8080/handle';

    let tmp = 25; // 温度
    let moi = 65; // 湿度
    let air = 1019; // 気圧

    const data = {
        "surroundings": [
            {
                "number": 1,
                "timestamp": (new Date()).toISOString(),
                "tempreture": tmp,
                "moisuture": moi,
                "airPressure": air,
            },
        ]
    };

    http.post(url, JSON.stringify(data), { headers: { "Content-Type": "application/json" } });
}
