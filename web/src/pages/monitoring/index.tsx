import React from 'react';
import { DiskMonitor } from './disk';
import { MemoryMonitor } from './memory';

type Props = {};

const Monitoring = (props: Props) => {
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        height: '100%',
        justifyContent: 'space-around',
      }}
    >
      <DiskMonitor />
      <MemoryMonitor />
    </div>
  );
};

export default Monitoring;
