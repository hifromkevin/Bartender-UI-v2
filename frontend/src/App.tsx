import { useEffect, useState } from 'react';
import Stations from '@features/Stations/Stations';
import SelectedStation from '@features/SelectedStation/SelectedStation';
import Cocktails from '@features/Cocktails/Cocktails';
import CocktailRecipe from '@features/CocktailRecipe/CocktailRecipe';

import { useFetchCatchphrase } from '@hooks/useFetchCatchphrase';
import { useFetchCatchphraseString } from '@hooks/useFetchCatchphraseString';

import logo from '@assets/bartender-buddy-logo.png';

import styles from './App.module.scss';

const App = () => {
  const [isStationSelected, setIsStationSelected] = useState<boolean>(false);
  const [showSelectedStation, setShowSelectedStation] =
    useState<boolean>(false);
  const [selectedStation, setSelectedStation] = useState<number | null>(null);

  const [isCocktailSelected, setIsCocktailSelected] = useState<boolean>(false);
  const [showCocktailRecipe, setShowCocktailRecipe] = useState<boolean>(false);
  // const [selectedCocktail, setSelectedCocktail] = useState<string | null>(null);

  const { catchphraseString } = useFetchCatchphraseString();
  const { catchphrase, catchphraseErrorisError } = useFetchCatchphrase();

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
        <p>
          {!catchphraseErrorisError &&
            catchphrase?.map(
              (
                item: { catchphrase: string; byline: string },
                index: number
              ) => (
                <span key={index}>
                  {item.catchphrase} {item.byline}
                </span>
              )
            )}
          {catchphraseString && catchphraseString}
        </p>
      </div>

      <div
        className={`${styles.appBodyContainer} ${showCocktailRecipe ? styles.shiftLeft : styles.shiftRight}`}
      >
        <div className={styles.stationsContainer}>
          <Stations
            setIsStationSelected={setIsStationSelected}
            setSelectedStation={setSelectedStation}
          />
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
                <SelectedStation
                  selectedStation={selectedStation}
                  setIsStationSelected={setIsStationSelected}
                />
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
