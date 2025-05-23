import { useEffect, useState } from 'react';

import { IStation, TStations } from '../../types/stationTypes';

import { arrow } from '../../assets/exportIcons';

import styles from './styles/Stations.module.scss';

interface IStationsProps {
  setIsStationSelected: Function;
}

const Stations = ({ setIsStationSelected }: IStationsProps) => {
  const [stations, setStations] = useState<TStations>([]);
  const [showStations, setShowStations] = useState<boolean>(true);

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
      <div className={styles.stationsTitle}>
        <h3>Stations</h3>
        <img
          src={arrow}
          alt="Toggle Open/Close"
          onClick={() => setShowStations((prevState) => !prevState)}
          className={`${styles.arrow} ${showStations ? '' : styles.toggleArrowRight}`}
        />
      </div>

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
              className={`${styles.stationName} ${showStations ? '' : styles.stationNameOnly} ${station?.station !== 3 ? styles.stationIsEmpty : ''}`}
            >
              {station?.station}
            </div>
            <div
              className={`${showStations ? styles.mixer : styles.hideMixerName} ${station?.station !== 3 ? styles.stationIsEmpty : ''}`}
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
