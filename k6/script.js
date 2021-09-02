import http from 'k6/http';
import {sleep} from 'k6';

export let options = {
    vus: 10,
    duration: '30s',
};
export default function () {
    http.get('http://localhost:8080/get?from=2019-02-01&to=2021-02-04&spaceId=Gomel_BrL2_turnstile');
    http.get(`http://localhost:8080/getCount?field=id`);
    http.put(`http://localhost:8080/update?field=id&value=1234&toId=10&fromId=10000`)
    // http.delete(` http://localhost:8080/delete?id=121321`)
    sleep(1);
}