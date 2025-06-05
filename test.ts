import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '15s', target: 200 },
    { duration: '30s', target: 400 },
    { duration: '15s', target: 400 },
    { duration: '30s', target: 0 },
  ],
};

export default function () {
  // let res = http.get('http://192.168.1.105:1080');
  //http://localhost:1080/
  let res = http.get('http://localhost:1080');
  // let res = http.get('http://pritam.dutta.vrianta.in');
  
  check(res, {
    'status is 200': (r) => r.status === 200,
    'response time < 500ms': (r) => r.timings.duration < 500,
  });

  sleep(Math.random() * 2);  // Simulate more realistic user pacing
}
