import {
  chocolatini,
  ginAndTonic,
  moscowMule,
  oldFashioned,
  rumAndCoke,
  shirleyTemple,
  vodkaCran,
  vodkaSoda,
} from '../../assets/exportCocktailIcons';

// import { cocktail, ice, info, wedge } from '../../assets/exportIcons';

import styles from './styles/CocktailRecipe.module.scss';

const cocktailImages: Record<string, string> = {
  'Moscow Mule': moscowMule,
  Chocolatini: chocolatini,
  'Gin and Tonic': ginAndTonic,
  'Old Fashioned': oldFashioned,
  'Rum and Coke': rumAndCoke,
  'Shirley Temple': shirleyTemple,
  'Vodka Cran': vodkaCran,
  'Vodka Soda': vodkaSoda,
};

const getCocktailImage = (cocktailName: string): string =>
  cocktailImages[cocktailName];

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
    <section
      className={styles.cocktailRecipe}
      onClick={() => setIsCocktailSelected(false)}
    >
      <div className={styles.cocktailRecipeImage}>
        <img
          src={getCocktailImage(cocktailData.name)}
          alt={cocktailData.name}
          className={styles.cocktailImage}
        />
      </div>
      <div className={styles.cocktailRecipeHeader}>
        <h3>Rum and Coke</h3>
        <p>{cocktailData.description}</p>
      </div>
    </section>
  );
};

export default CocktailRecipe;
