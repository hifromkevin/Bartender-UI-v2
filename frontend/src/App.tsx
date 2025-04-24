import Stations from './features/Stations/Stations';
import Cocktails from './features/Cocktails/Cocktails';

import logo from './assets/bartender-buddy-logo.png';

import styles from './App.module.scss';

const App = () => (
  <div className={styles.appContainer}>
    <div className={styles.appHeader}>
      <img src={logo} alt="Bartender Buddy Logo" className={styles.logo} />
    </div>
    <div className={styles.appBodyContainer}>
      <div className={styles.stationsContainer}>
        <Stations />
      </div>
      <div className={styles.cocktailsContainer}>
        <Cocktails />
      </div>
    </div>
  </div>
);

export default App;
