import { useEffect, useState } from 'react';

import { IStation, TStations } from '../../types/stationTypes';

import styles from './styles/Stations.module.scss';

const Stations = () => {
  const [stations, setStations] = useState<TStations>([]);

  useEffect(() => {
    setStations(
      Array(8)
        .fill(null)
        .map((_, index) => ({
          station: index + 1,
          mixer: null,
        }))
    );
  }, []);

  return (
    <aside className={styles.stationsContainer}>
      <h3>Stations</h3>

      <section>
        {stations.map((station: IStation) => (
          <div key={station?.station} className={styles.station}>
            <div className={styles.stationName}>{station?.station}</div>
            <div className={styles.mixer}>{station?.mixer || 'Empty'}</div>
          </div>
        ))}
      </section>
    </aside>
  );
};

export default Stations;
