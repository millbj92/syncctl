import axios from 'axios';
import React, { Suspense, useEffect, useState } from 'react';

type Props = {};

interface IMemoryStat {
  total: number;
  used: number;
  usedPercent: number;
}

interface ISwapMemoryStat extends IMemoryStat {
  free: number;
}

interface IVirtualMemoryStat extends IMemoryStat {
  available: number;
}

interface AllMemoryStat {
  VirtualMemory: IVirtualMemoryStat;
  SwapMemory: ISwapMemoryStat;
  SwapDevices: ISwapDevice[];
}

interface ISwapDevice {
  name: string;
  usedBytes: number;
  freeBytes: number;
}

type NetResponse<T> = {
  error: boolean;
  msg: string;
  data: T;
};

export const MemoryMonitor = (props: Props) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);
  const [resp, setResp] = useState({} as NetResponse<AllMemoryStat>);

  const fetchData = async () => {
    setLoading(true);
    setError(false);

    try {
      const response = await axios.get('http://localhost:8101/api/v1/memory');
      console.log(response);
      setResp(response.data);

      if (response.data.error) {
        setError(true);
      }
    } catch (error) {
      setError(true);
    }
    setLoading(false);
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
        <div>{RenderVmem(resp.data)}</div>
        <div>{RenderSwapMemory(resp.data)}</div>
        <div>{RenderSwapDevices(resp.data)}</div>
      </div>
    </Suspense>
  );
};

const RenderVmem = (mem: AllMemoryStat | undefined) => {
  if (!mem) return null;

  const vmem = mem.VirtualMemory;
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
    </div>
  );
};

const RenderSwapMemory = (mem: AllMemoryStat | undefined) => {
  if (!mem) return null;

  const smem = mem.SwapMemory;
  return (
    <div>
      <h2>Swap Memory</h2>
      <div>
        <span>Total:</span>
        <p>{smem.total}</p>
      </div>
      <div>
        <span>Used:</span>
        <p>{smem.used}</p>
      </div>
      <div>
        <span>Free:</span>
        <p>{smem.free}</p>
      </div>
      <div>
        <span>Used Percent:</span>
        <p>{smem.usedPercent}</p>
      </div>
    </div>
  );
};

const RenderSwapDevices = (mem: AllMemoryStat | undefined) => {
  if (!mem) return null;

  const sdev = mem.SwapDevices;
  return (
    <div>
      <h2>Swap Devices</h2>
      <div>
        {sdev.map((device, index) => (
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
  );
};
