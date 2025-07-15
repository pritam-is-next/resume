import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 100 },
    { duration: '20s', target: 500 },
    { duration: '10s', target: 0 },
  ],
};

export default function () {
  const url = 'http://localhost:1080/login'; // Replace with your actual login endpoint

  const payload = 'loginEmail=pritam&loginPassword=pass'; // Adjust field names to match your form

  const params = {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
  };

  const res = http.post(url, payload, params);

  check(res, {
    'status is 200': (r) => r.status === 200,
    'login success text found': (r) =>
      r.body.includes('Dashboard'),
    'response time < 500ms': (r) => r.timings.duration < 500,
  });

  sleep(Math.random() * 2); // simulate real user pacing
}
