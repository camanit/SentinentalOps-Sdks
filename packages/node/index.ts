export interface SentinelConfig {
  endpoint: string;
  apiKey: string;
  tenantId: string;
  serviceName?: string;
}

export interface IncidentReport {
  title: string;
  division: string;
  severity: 'CRITICAL' | 'HIGH' | 'MEDIUM' | 'LOW';
  details?: string;
  metadata?: Record<string, unknown>;
}

export class SentinelOps {
  private config: SentinelConfig;

  constructor(config: SentinelConfig) {
    this.config = config;
  }

  async reportIncident(incident: IncidentReport): Promise<{ success: boolean; id?: string }> {
    try {
      const res = await fetch(\\/api/v1/incidents\, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': \Bearer \\,
          'X-Sentinel-Tenant': this.config.tenantId,
        },
        body: JSON.stringify({
          ...incident,
          service: this.config.serviceName,
          timestamp: new Date().toISOString(),
        }),
      });
      return await res.json();
    } catch (err) {
      console.error('[SentinelOps SDK Error]', err);
      return { success: false };
    }
  }
}
