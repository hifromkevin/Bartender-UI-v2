import {
  absolut,
  bacardi,
  bombayGin,
  bottle,
  canadaDryGingerAle,
  cranberryJuice,
} from '../../assets/exportMixers';

import styles from './styles/SelectedStation.module.scss';

const mixerImages: Record<string, string> = {
  'Absolut Vodka': absolut,
  'Bacardi Rum': bacardi,
  'Bombay Gin': bombayGin,
  Bottle: bottle,
  'Canada Dry Ginger Ale': canadaDryGingerAle,
  'Motts Organic Cranberry Juice': cranberryJuice,
};

const getMixerImage = (mixerName: string): string => mixerImages[mixerName];

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
    brand: 'Canda Dry',
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

const SelectedStation = () => {
  return (
    <section className={styles.selectedStation}>
      <div className={styles.stationHeader}>
        <h3>Station 3</h3>
      </div>

      <section className={styles.selector}>
        <section className={styles.stationSelection}>
          <div>
            <div className={styles.stationData}>
              <img
                src={absolut}
                alt="Absolut Vodka"
                className={styles.stationImage}
              />
              <h4>Vodka</h4>
              <p>Absolut Vodka</p>
            </div>
          </div>
          <div className={styles.verticalLine}> </div>
        </section>
        <section className={styles.stationOptions}>
          {fakeMixers.map((mixer) => (
            <div key={mixer.id} className={styles.mixerItem}>
              <img
                src={getMixerImage(mixer.name)}
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
