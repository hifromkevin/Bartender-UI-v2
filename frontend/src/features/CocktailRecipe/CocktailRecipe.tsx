import CloseButton from '@components/CloseButton/CloseButton';
import Tabs from '@components/Tabs/Tabs';

import fakeCocktail from '@assets/fakeCocktail.jpg';

import styles from './styles/CocktailRecipe.module.scss';

const cocktailData = {
  id: 1,
  name: 'Moscow Mule',
  description: 'A classic cocktail made with whiskey, lemon juice, and sugar.',
  glass_type: 'Cocktail',
  ice_type: 'Cubed',
  requires_shaker: true,
  instructions_id: 1,
  difficulty: 'Medium',
};

interface ICocktailRecipeProps {
  setIsCocktailSelected: Function;
}

const CocktailRecipe = ({ setIsCocktailSelected }: ICocktailRecipeProps) => {
  return (
    <section className={styles.cocktailRecipe}>
      <CloseButton action={setIsCocktailSelected} />

      <div className={styles.info}>
        <div className={styles.image}>
          <img
            src={fakeCocktail}
            alt={cocktailData.name}
            className={styles.cocktailImage}
          />
        </div>
        <div className={styles.instructions}>
          <div className={styles.header}>
            <h4>Rum and Coke</h4>
            <p>{cocktailData.description}</p>
          </div>
          <Tabs />
        </div>
        <div className={styles.buttons}>
          <button
            className={styles.goBack}
            onClick={() => setIsCocktailSelected(false)}
          >
            Go Back
          </button>
          <button className={styles.startCocktail}>Start Cocktail</button>
        </div>
      </div>
    </section>
  );
};

export default CocktailRecipe;
