// docker run --network go-ecommerce-backend-api_app-network --rm -i grafana/k6 run - <loadtest-script.js
// k6 run loadtest-script.js
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 30, // vus (virtual users)
  duration: '30s',
};

export default function () {
  http.get('http://app:8002/ping/200');
  sleep(0.5);
}