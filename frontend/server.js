const express = require('express');
const axios = require('axios');
const client = require('prom-client');

const app = express();
const PORT = 3000;

// Prometheus metrics
const register = new client.Registry();
client.collectDefaultMetrics({ register });

const httpRequestsTotal = new client.Counter({
  name: 'frontend_http_requests_total',
  help: 'Total HTTP requests',
  labelNames: ['method', 'endpoint', 'status'],
});
register.registerMetric(httpRequestsTotal);

app.get('/health', (req, res) => {
  res.json({ status: 'ok' });
});

app.get('/api/proxy', async (req, res) => {
  try {
    const start = Date.now();
    const response = await axios.get('http://backend:8080/api/data');
    const latency = Date.now() - start;
    
    httpRequestsTotal.labels('GET', '/api/proxy', response.status).inc();
    
    res.json({
      backend_data: response.data,
      proxy_latency_ms: latency,
    });
  } catch (error) {
    httpRequestsTotal.labels('GET', '/api/proxy', error.response?.status || 500).inc();
    res.status(500).json({ error: error.message });
  }
});

app.get('/metrics', async (req, res) => {
  res.set('Content-Type', register.contentType);
  res.end(await register.metrics());
});

app.listen(PORT, () => {
  console.log(`Frontend running on :${PORT}`);
});