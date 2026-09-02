import requests
from datetime import datetime

class SentinelClient:
    def __init__(self, endpoint: str, api_key: str, tenant_id: str, service_name: str = 'python-service'):
        self.endpoint = endpoint.rstrip('/')
        self.api_key = api_key
        self.tenant_id = tenant_id
        self.service_name = service_name

    def report_incident(self, title: str, division: str, severity: str = 'HIGH', details: str = ''):
        payload = {
            'title': title,
            'division': division,
            'severity': severity,
            'details': details,
            'service': self.service_name,
            'timestamp': datetime.utcnow().isoformat()
        }
        headers = {
            'Content-Type': 'application/json',
            'Authorization': f'Bearer {self.api_key}',
            'X-Sentinel-Tenant': self.tenant_id
        }
        try:
            r = requests.post(f'{self.endpoint}/api/v1/incidents', json=payload, headers=headers, timeout=5)
            return r.json()
        except Exception as e:
            return {'success': False, 'error': str(e)}
