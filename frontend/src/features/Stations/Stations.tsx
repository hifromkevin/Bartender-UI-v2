import { useEffect, useState } from 'react';

import { IStation, TStations } from '../../types/stationTypes';

import styles from './styles/Stations.module.scss';

interface IStationsProps {
  setIsStationSelected: Function;
}

const Stations = ({ setIsStationSelected }: IStationsProps) => {
  const [stations, setStations] = useState<TStations>([]);

  useEffect(() => {
    setStations(
      Array(8)
        .fill(null)
        .map((_, index) => ({
          station: index + 1,
          mixer: index === 2 ? 'Spiced Pear Liqueur' : null,
        }))
    );
  }, []);

  return (
    <aside className={styles.stations}>
      <h3>Stations</h3>

      <section>
        {stations.map((station: IStation) => (
          <div
            key={station?.station}
            className={styles.station}
            onClick={() =>
              setIsStationSelected((prevState: boolean): boolean => !prevState)
            }
          >
            <div
              className={`${styles.stationName} ${station?.station !== 3 ? styles.stationIsEmpty : ''}`}
            >
              {station?.station}
            </div>
            <div
              className={`${styles.mixer} ${station?.station !== 3 ? styles.stationIsEmpty : ''}`}
            >
              {station?.mixer || 'Empty'}
            </div>
          </div>
        ))}
      </section>
    </aside>
  );
};

export default Stations;
