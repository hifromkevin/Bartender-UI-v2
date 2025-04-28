import { useState } from 'react';
import Stations from './features/Stations/Stations';
import SelectedStation from './features/SelectedStation/SelectedStation';
import Cocktails from './features/Cocktails/Cocktails';

import logo from './assets/bartender-buddy-logo.png';

import styles from './App.module.scss';

const App = () => {
  const [isStationSelected, setIsStationSelected] = useState<boolean>(false);

  return (
    <div className={styles.appContainer}>
      <div className={styles.appHeader}>
        <img src={logo} alt="Bartender Buddy Logo" className={styles.logo} />
      </div>

      <div className={styles.appBodyContainer}>
        <div className={styles.stationsContainer}>
          <Stations setIsStationSelected={setIsStationSelected} />
          <div className={isStationSelected ? '' : styles.verticalLine}> </div>
        </div>
        <div className={styles.cocktailsContainer}>
          {isStationSelected ? <SelectedStation /> : <Cocktails />}
        </div>
      </div>
    </div>
  );
};

export default App;
