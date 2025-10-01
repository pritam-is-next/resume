import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 200 },
    { duration: '30s', target: 500 },
    { duration: '30s', target: 0 },
  ],
  thresholds: {
    'http_req_duration{expected_response:true}': ['p(95)<500'], // fail if 95% responses > 500ms
    'http_req_failed': ['rate<0.01'], // allow <1% failure
  }

};

export default function () {
  // let res = http.get('http://192.168.1.105:1080');
  //http://localhost:1080/
  let res = http.get('http://localhost:1088');
  // let res = http.get('http://pritam.dutta.vrianta.in');

  check(res, {
    'status is 200': (r) => r.status === 200,
    'response time < 500ms': (r) => r.timings.duration < 500,
  });

  sleep(Math.random() * 2);  // Simulate more realistic user pacing
}
