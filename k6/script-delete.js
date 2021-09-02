import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    http.get('http://localhost:8080/delete?id=7404802');
    sleep(1);
}