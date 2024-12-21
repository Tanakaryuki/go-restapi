import http from 'k6/http';
import { check, sleep, randomSeed } from 'k6';

randomSeed(1234);

export const options = {
    scenarios: {
        user_creation: {
            executor: 'constant-vus',
            vus: 50,
            duration: '20s',
        },
    },
};

const BASE_URL = 'http://localhost:8080';

function createUser(randomValue) {
    const username = `user${randomValue}`;
    const payload = JSON.stringify({
        email: `${username}@example.com`,
        password: 'password123',
        username: username,
        display_name: `User ${randomValue}`,
        is_admin: false,
    });
    const headers = { 'Content-Type': 'application/json' };
    return http.post(`${BASE_URL}/signup`, payload, { headers });
}

function loginUser(username) {
    const payload = JSON.stringify({
        username: username,
        password: 'password123',
    });
    const headers = { 'Content-Type': 'application/json' };
    return http.post(`${BASE_URL}/login`, payload, { headers });
}

function createTask(token, randomValue) {
    const payload = JSON.stringify({
        title: `Task for user${randomValue}`,
        detail: `Details for user${randomValue}'s task.`,
        id: `task_${randomValue}`,
    });
    const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
    };
    return http.post(`${BASE_URL}/tasks`, payload, { headers });
}

function getTaskDetail(token, taskId) {
    const headers = { Authorization: `Bearer ${token}` };
    return http.get(`${BASE_URL}/tasks/${taskId}`, { headers });
}

export default function () {
    const randomValue = Math.floor(Math.random() * 1000000);

    // Create a new user
    const createUserRes = createUser(randomValue);
    check(createUserRes, { 'created user': (res) => res.status === 201 });

    // Login with the new user
    const loginRes = loginUser(`user${randomValue}`);
    check(loginRes, { 'logged in': (res) => res.status === 200 });

    const token = loginRes.json('token');

    // Create a new task
    const createTaskRes = createTask(token, randomValue);
    check(createTaskRes, { 'created task': (res) => res.status === 201 });

    // Get task details
    const taskId = `task_${randomValue}`;
    const getTaskDetailRes = getTaskDetail(token, taskId);
    check(getTaskDetailRes, { 'retrieved task details': (res) => res.status === 200 });

    sleep(1);
}