import React, { Suspense, useEffect, useState } from 'react';

type Props = {};

interface ISwapMemoryStat {
  total: number;
  used: number;
  free: number;
  usedPercent: number;
  sin: number;
  sout: number;
  pgIn: number;
  pgOut: number;
  pgFault: number;
  pgMajFault: number;
}

interface IVirtualMemoryStat {
  total: number;
  available: number;
  used: number;
  usedPercent: number;
  free: number;
  active: number;
  inactive: number;
  wired: number;
  laundry: number;
  buffers: number;
  cached: number;
  writeBack: number;
  dirty: number;
  writeBackTmp: number;
  shared: number;
  slab: number;
  sreclaimable: number;
  sunreclaim: number;
  pageTables: number;
  swapCached: number;
  commitLimit: number;
  committedAS: number;
  highTotal: number;
  highFree: number;
  lowTotal: number;
  lowFree: number;
  swapTotal: number;
  swapFree: number;
  mapped: number;
  vmallocTotal: number;
  vmallocUsed: number;
  vmallocChunk: number;
  hugePagesTotal: number;
  hugePagesFree: number;
  hugePageSize: number;
}

interface ISwapDevice {
  name: string;
  usedBytes: number;
  freeBytes: number;
}

type MemoryUsageResponse = {
  vmem: IVirtualMemoryStat;
  smem: ISwapMemoryStat;
  sdev: ISwapDevice[];
  error: boolean;
  msg: string;
};

export const MemoryMonitor = (props: Props) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);
  const [resp, setResp] = useState({} as MemoryUsageResponse);

  const fetchData = async () => {
    setLoading(true);
    setError(false);

    fetch('http://localhost:8080/api/v1/memory')
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        if (json.error) {
          setError(true);
          console.log(error);
        }
        console.log(json);
        setResp(json);
      })
      .catch((error) => {
        setError(true);
        console.log(error);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <Suspense>
      <div style={{ width: '100%', height: '100%' }}>
        {loading && <div>Loading...</div>}
        {error && <div>Error!</div>}
        <h1>Memory</h1>
        <div>
          <div>{RenderVmem(resp.vmem)}</div>
        </div>

        {/* {resp && resp.smem && (
          <div>
            <h2>Swap Memory</h2>
            <div>
              <span>Total:</span>
              <p>{resp.smem.total}</p>
            </div>
            <div>
              <span>Used:</span>
              <p>{resp.smem.used}</p>
            </div>
            <div>
              <span>Free:</span>
              <p>{resp.smem.free}</p>
            </div>
            <div>
              <span>Used Percent:</span>
              <p>{resp.smem.usedPercent}</p>
            </div>
          </div>
        )}

        {resp && resp.sdev && (
          <div>
            <h2>Swap Devices</h2>
            <div>
              {resp.sdev.map((device, index) => (
                <div key={index}>
                  <span>Name:</span>
                  <p>{device.name}</p>
                  <span>Used:</span>
                  <p>{device.usedBytes}</p>
                  <span>Free:</span>
                  <p>{device.freeBytes}</p>
                </div>
              ))}
            </div>
          </div>
        )} */}
        <button onClick={fetchData}>Refresh</button>
      </div>
    </Suspense>
  );
};

const RenderVmem = (vmem: IVirtualMemoryStat | undefined) => {
  if (!vmem) return null;
  return (
    <div>
      <h2>Virtual Memory</h2>
      <div>
        <span>Total:</span>
        <p>{vmem.total}</p>
      </div>
      <div>
        <span>Available:</span>
        <p>{vmem.available}</p>
      </div>
      <div>
        <span>Used:</span>
        <p>{vmem.used}</p>
      </div>
      <div>
        <span>Used Percent:</span>
        <p>{vmem.usedPercent}</p>
      </div>
      <div>
        <span>High:</span>
        <p>{vmem.highTotal}</p>
      </div>
      <div>
        <span>High Free:</span>
        <p>{vmem.highFree}</p>
      </div>
      <div>
        <span>Low:</span>
        <p>{vmem.lowTotal}</p>
      </div>
      <div>
        <span>Low Free:</span>
        <p>{vmem.lowFree}</p>
      </div>
    </div>
  );
};
