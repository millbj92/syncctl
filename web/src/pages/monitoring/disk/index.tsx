import { v4 } from 'uuid';
import React, { Suspense, useEffect, useState } from 'react';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  ArcElement,
  TooltipItem,
} from 'chart.js';

import {
  StyledWidgetContainer,
  StyledWidgetTitle,
} from '../../../components/widgets';

import { Bar, Line, Scatter, Bubble, Pie, Doughnut } from 'react-chartjs-2';

import styled from 'styled-components';

const StyledDashboard = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: flex-start;
  width: 100%;
  height: 100%;
  padding: 4rem;
`;

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
  Title,
  Tooltip,
  Legend,
  ArcElement
);

type Props = {};
interface IDiskStat {
  path: string;
  fstype: string;
  total: number;
  free: number;
  used: number;
  usedPercent: number;
  inodesTotal: number;
  inodesUsed: number;
  inodesFree: number;
  inodesUsedPercent: number;
}

type DiskStatResponse = {
  data: IDiskStat;
  error: boolean;
  msg: string;
};

export const DiskMonitor = (props: Props) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);
  const [data, setData] = useState({} as IDiskStat);

  const fetchData = async () => {
    setLoading(true);
    setError(false);
    try {
      const response = await fetch(
        `http://${process.env.REACT_APP_API_HOST}:${process.env.REACT_APP_API_PORT}/api/v1/disk/usage/`
      );
      console.log(response);
      const json = (await response.json()) as DiskStatResponse;
      if (json.error) {
        console.log(json);
        //setError(true);
        //return;
      }
      console.log(json);
      setData(json.data);
      setLoading(false);
    } catch (error) {
      console.log(error);
      setError(true);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);
  return (
    <Suspense>
      <StyledDashboard>
        <h1>PC 01</h1>
        {loading && <p>Loading...</p>}
        {error && <p>Error!</p>}
        {data && (
          <StyledWidgetContainer>
            <StyledWidgetTitle>Disk Usage</StyledWidgetTitle>
            <hr style={{ width: '100%', color: 'black' }} />
            <Pie
              data={{
                labels: ['Free', 'Used'],
                datasets: [
                  {
                    label: 'Disk Usage',
                    data: [
                      data.usedPercent?.toFixed(2) ?? 0.0,
                      (100 - data.usedPercent)?.toFixed(2) ?? 0.0,
                    ],
                    backgroundColor: ['rgb(2, 255, 88)', 'rgb(255,2,88)'],
                    hoverOffset: 4,
                  },
                ],
              }}
              options={{
                elements: {
                  arc: {
                    offset: -10,
                    borderAlign: 'center',
                    borderWidth: 3,
                    hoverBorderWidth: 5,
                    borderColor: 'rgb(255,255,255)',
                    hoverBorderColor: '#ffffff',
                    borderRadius: 0,
                    hoverOffset: 15,
                  },
                },
                plugins: {
                  tooltip: {
                    callbacks: {
                      label: function (tooltipItem: TooltipItem<'pie'>) {
                        return `${tooltipItem.label}: ${tooltipItem.formattedValue}%`;
                      },
                    },
                  },
                },
              }}
              width={4}
              height={4}
            />
          </StyledWidgetContainer>
        )}
      </StyledDashboard>
    </Suspense>
  );
};
