import CloseButton from '@components/CloseButton/CloseButton';

import fakeMixer from '@assets/fakeMixer.jpg';

import styles from './styles/SelectedStation.module.scss';

// const mixerImages: Record<string, string> = {
//   'Absolut Vodka': absolut,
//   'Bacardi Rum': bacardi,
//   'Bombay Gin': bombayGin,
//   Bottle: bottle,
//   'Canada Dry Ginger Ale': canadaDryGingerAle,
//   'Motts Organic Cranberry Juice': cranberryJuice,
// };

// const getMixerImage = (mixerName: string): string => mixerImages[mixerName];

const fakeMixers = [
  {
    id: 1,
    name: 'Absolut Vodka',
    brand: 'Absolut',
    type: 'Vodka',
    flavor: null,
    is_favorite: false,
    times_chosen: new Date(),
    times_used: new Date(),
    last_selected: new Date(),
  },
  {
    id: 2,
    name: 'Bacardi Rum',
    brand: 'Bacardi',
    type: 'Rum',
    flavor: null,
    is_favorite: false,
    times_chosen: new Date(),
    times_used: new Date(),
    last_selected: new Date(),
  },
  {
    id: 3,
    name: 'Bombay Gin',
    brand: 'Bombay',
    type: 'Gin',
    flavor: null,
    is_favorite: false,
    times_chosen: new Date(),
    times_used: new Date(),
    last_selected: new Date(),
  },
  {
    id: 4,
    name: 'Canada Dry Ginger Ale',
    brand: 'Canada Dry',
    type: 'Ginger Ale',
    flavor: null,
    is_favorite: false,
    times_chosen: new Date(),
    times_used: new Date(),
    last_selected: new Date(),
  },
  {
    id: 5,
    name: 'Motts Organic Cranberry Juice',
    brand: 'Motts',
    type: 'Cranberry Juice',
    flavor: null,
    is_favorite: false,
    times_chosen: new Date(),
    times_used: new Date(),
    last_selected: new Date(),
  },
];

interface ISelectedStation {
  selectedStation: number | null;
  setIsStationSelected: Function;
}
const SelectedStation = ({
  selectedStation,
  setIsStationSelected,
}: ISelectedStation) => {
  return (
    <section className={styles.selectedStation}>
      <div className={styles.stationHeader}>
        <h3>Station {selectedStation}</h3>
        <CloseButton action={setIsStationSelected} />
      </div>

      <section className={styles.selector}>
        <section className={styles.stationSelection}>
          <div className={styles.stationData}>
            <img
              src={fakeMixer}
              alt="Absolut Vodka"
              className={styles.stationImage}
            />
            <h4>Vodka</h4>
            <p>Absolut Vodka</p>
          </div>
          <div className={styles.verticalLine}> </div>
        </section>
        <section className={styles.stationOptions}>
          {fakeMixers.map((mixer) => (
            <div key={mixer.id} className={styles.mixerItem}>
              <img
                src={fakeMixer}
                alt={mixer.name}
                className={styles.mixerImage}
              />
              <h4>{mixer.name}</h4>
              <p>{mixer.brand}</p>
              <p>{mixer.type}</p>
            </div>
          ))}
        </section>
      </section>
    </section>
  );
};

export default SelectedStation;
