import { useEffect, useState } from 'react';
import Stations from './features/Stations/Stations';
import SelectedStation from './features/SelectedStation/SelectedStation';
import Cocktails from './features/Cocktails/Cocktails';
import CocktailRecipe from './features/CocktailRecipe/CocktailRecipe';

import logo from './assets/bartender-buddy-logo.png';

import styles from './App.module.scss';

const App = () => {
  const [isStationSelected, setIsStationSelected] = useState<boolean>(false);
  const [showSelectedStation, setShowSelectedStation] = useState(false);

  const [isCocktailSelected, setIsCocktailSelected] = useState<boolean>(false);
  const [showCocktailRecipe, setShowCocktailRecipe] = useState(false);

  useEffect(() => {
    if (isStationSelected) {
      setShowSelectedStation(true);
    } else {
      const timeout = setTimeout(() => setShowSelectedStation(false), 600);
      return () => clearTimeout(timeout);
    }
  }, [isStationSelected]);

  useEffect(() => {
    if (isCocktailSelected) {
      setShowCocktailRecipe(true);
    } else {
      setShowCocktailRecipe(false);
    }
  }, [isCocktailSelected]);

  return (
    <div className={styles.appContainer}>
      <div className={styles.appHeader}>
        <img src={logo} alt="Bartender Buddy Logo" className={styles.logo} />
      </div>

      <div
        className={`${styles.appBodyContainer} ${showCocktailRecipe ? styles.shiftLeft : styles.shiftRight}`}
      >
        <div className={styles.stationsContainer}>
          <Stations setIsStationSelected={setIsStationSelected} />
          <div className={styles.verticalLine}></div>
        </div>
        <div className={styles.cocktailsContainer}>
          <div className={styles.cocktailSlideContainer}>
            <Cocktails setIsCocktailSelected={setIsCocktailSelected} />
            {showSelectedStation && (
              <div
                className={`${styles.selectedStationSlide} ${
                  !isStationSelected ? styles.slideOut : ''
                }`}
              >
                <SelectedStation />
              </div>
            )}
          </div>
        </div>
      </div>
      <div
        className={`${styles.cocktailRecipeSlide} ${showCocktailRecipe ? styles.slideIn : styles.slideOut}`}
      >
        <CocktailRecipe setIsCocktailSelected={setIsCocktailSelected} />
      </div>
    </div>
  );
};

export default App;
